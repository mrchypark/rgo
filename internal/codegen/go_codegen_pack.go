// Copyright ©2019 The rgonomic Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package codegen

import (
	"bytes"
	"fmt"
	"go/types"

	"github.com/rgonomic/rgo/internal/pkg"
)

// packSEXPFuncGo returns the source of functions to pack the given Go-typed
// parameters into R SEXP values.
func packSEXPFuncGo(typs []types.Type) string {
	var buf bytes.Buffer
	for _, typ := range typs {
		fmt.Fprintf(&buf, "func packSEXP%s(p %s) C.SEXP {\n", pkg.Mangle(typ), nameOf(typ))
		packSEXPFuncBodyGo(&buf, typ)
		buf.WriteString("}\n\n")
	}
	return buf.String()
}

// packSEXPFuncGo returns the body of a function to pack the given Go-typed
// parameters into R SEXP values.
func packSEXPFuncBodyGo(buf *bytes.Buffer, typ types.Type) {
	switch typ := typ.(type) {
	case *types.Named:
		if pkg.IsError(typ) {
			fmt.Fprintf(buf, `	if p == nil {
		return C.R_NilValue
	}
	return packSEXP%s(p.Error())
`, pkg.Mangle(types.Typ[types.String]))
		} else {
			switch typ := typ.Underlying().(type) {
			case *types.Pointer:
				fmt.Fprintf(buf, "\treturn packSEXP%s((%s)(p))\n", pkg.Mangle(typ), typ)
			default:
				fmt.Fprintf(buf, "\treturn packSEXP%s(%s(p))\n", pkg.Mangle(typ), typ)
			}
		}

	case *types.Array:
		fmt.Fprintf(buf, "\treturn packSEXP%s(p[:])\n", pkg.Mangle(types.NewSlice(typ.Elem())))

	case *types.Basic:
		switch typ.Kind() {
		case types.Bool:
			fmt.Fprintf(buf, `	b := C.int(0)
	if p {
		b = 1
	}
	return C.ScalarLogical(b)
`)
		case types.Int, types.Int8, types.Int16, types.Int32, types.Int64, types.Uint, types.Uint16, types.Uint32, types.Uint64:
			fmt.Fprintln(buf, "\treturn C.ScalarInteger(C.int(p))")
		case types.Uint8:
			fmt.Fprintln(buf, "\treturn C.ScalarRaw(C.Rbyte(p))")
		case types.Float64, types.Float32:
			fmt.Fprintln(buf, "\treturn C.ScalarReal(C.double(p))")
		case types.Complex128, types.Complex64:
			fmt.Fprintln(buf, "\treturn C.ScalarComplex(C.struct_Rcomplex{r: C.double(real(p)), i: C.double(imag(p))})")
		case types.String:
			fmt.Fprintln(buf, `	s := C.Rf_mkCharLenCE(C._GoStringPtr(p), C.int(len(p)), C.CE_UTF8)
	return C.ScalarString(s)`)
		default:
			panic(fmt.Sprintf("unhandled type: %s", typ))
		}

	case *types.Map:
		// TODO(kortschak): Handle named simple types properly.
		elem := typ.Elem()
		if basic, ok := elem.Underlying().(*types.Basic); ok {
			switch basic.Kind() {
			// TODO(kortschak): Make the fast path available
			// to []T where T is one of these kinds.
			case types.Int, types.Int8, types.Int16, types.Int32, types.Uint, types.Uint16, types.Uint32:
				// Maximum length array type for this element type.
				type a [1 << 47]int32
				fmt.Fprintf(buf, `	n := len(p)
	r := C.Rf_allocVector(C.%[1]s, C.R_xlen_t(n))
	C.Rf_protect(r)
	names := C.Rf_allocVector(C.STRSXP, C.R_xlen_t(n))
	C.Rf_protect(names)
	s := (*[%[2]d]int32)(unsafe.Pointer(C.INTEGER(r)))[:len(p):len(p)]
	var i C.R_xlen_t
	for k, v := range p {
		C.SET_STRING_ELT(names, i, C.Rf_mkCharLenCE(C._GoStringPtr(k), C.int(len(k)), C.CE_UTF8))
		s[i] = int32(v)
		i++
	}
	C.setAttrib(r, packSEXP_types_Basic_string("names"), names)
	C.Rf_unprotect(2)
	return r
`, rTypeLabelFor(elem), len(&a{}))
				return

			case types.Uint8:
				// Maximum length array type for this element type.
				type a [1 << 49]byte
				fmt.Fprintf(buf, `	n := len(p)
	r := C.Rf_allocVector(C.%[1]s, C.R_xlen_t(n))
	C.Rf_protect(r)
	names := C.Rf_allocVector(C.STRSXP, C.R_xlen_t(n))
	C.Rf_protect(names)
	s := (*[%[2]d]uint8)(unsafe.Pointer(C.RAW(r)))[:len(p):len(p)]
	var i C.R_xlen_t
	for k, v := range p {
		C.SET_STRING_ELT(names, i, C.Rf_mkCharLenCE(C._GoStringPtr(k), C.int(len(k)), C.CE_UTF8))
		i++
	}
	copy(s, p)
	C.setAttrib(r, packSEXP_types_Basic_string("names"), names)
	C.Rf_unprotect(2)
	return r
`, rTypeLabelFor(elem), len(&a{}), pkg.Mangle(elem))
				return

			case types.Float32, types.Float64:
				// Maximum length array type for this element type.
				type a [1 << 46]float64
				fmt.Fprintf(buf, `	n := len(p)
	r := C.Rf_allocVector(C.%[1]s, C.R_xlen_t(n))
	C.Rf_protect(r)
	names := C.Rf_allocVector(C.STRSXP, C.R_xlen_t(n))
	C.Rf_protect(names)
	s := (*[%[2]d]float64)(unsafe.Pointer(C.REAL(r)))[:len(p):len(p)]
	var i C.R_xlen_t
	for k, v := range p {
		C.SET_STRING_ELT(names, i, C.Rf_mkCharLenCE(C._GoStringPtr(k), C.int(len(k)), C.CE_UTF8))
		s[i] = float64(v)
		i++
	}
	C.setAttrib(r, packSEXP_types_Basic_string("names"), names)
	C.Rf_unprotect(2)
	return r
`, rTypeLabelFor(elem), len(&a{}))
				return

			case types.Complex64, types.Complex128:
				// Maximum length array type for this element type.
				type a [1 << 45]complex128
				fmt.Fprintf(buf, `	n := len(p)
	r := C.Rf_allocVector(C.%[1]s, C.R_xlen_t(n))
	C.Rf_protect(r)
	names := C.Rf_allocVector(C.STRSXP, C.R_xlen_t(n))
	C.Rf_protect(names)
	s := (*[%[2]d]complex128)(unsafe.Pointer(C.COMPLEX(r)))[:len(p):len(p)]
	var i C.R_xlen_t
	for k, v := range p {
		C.SET_STRING_ELT(names, i, C.Rf_mkCharLenCE(C._GoStringPtr(k), C.int(len(k)), C.CE_UTF8))
		s[i] = complex128(v)
		i++
	}
	C.setAttrib(r, packSEXP_types_Basic_string("names"), names)
	C.Rf_unprotect(2)
	return r
`, rTypeLabelFor(elem), len(&a{}))
				return

			case types.String:
				fmt.Fprintf(buf, `	n := len(p)
	r := C.Rf_allocVector(C.%[1]s, C.R_xlen_t(n))
	C.Rf_protect(r)
	names := C.Rf_allocVector(C.STRSXP, C.R_xlen_t(n))
	C.Rf_protect(names)
	var i C.R_xlen_t
	for k, v := range p {
		C.SET_STRING_ELT(names, i, C.Rf_mkCharLenCE(C._GoStringPtr(k), C.int(len(k)), C.CE_UTF8))
		C.SET_STRING_ELT(r, i, packSEXP%s(v))
		i++
	}
	C.setAttrib(r, packSEXP_types_Basic_string("names"), names)
	C.Rf_unprotect(2)
	return r
`, rTypeLabelFor(elem), pkg.Mangle(elem))
				return

			case types.Bool:
				// Maximum length array type for this element type.
				type a [1 << 47]int32
				// FIXME(kortschak): Does Rf_allocVector return a
				// zeroed vector? If it does, the loop below doesn't
				// need the else clause.
				// Alternatively, convert the []bool to a []byte:
				//  for i, v := range *(*[]byte)(unsafe.Pointer(&p)) {
				//      s[i] = int32(v)
				//  }
				fmt.Fprintf(buf, `	n := len(p)
	r := C.Rf_allocVector(C.LGLSXP, C.R_xlen_t(n))
	C.Rf_protect(r)
	names := C.Rf_allocVector(C.STRSXP, C.R_xlen_t(n))
	C.Rf_protect(names)
	s := (*[%d]int32)(unsafe.Pointer(C.LOGICAL(r)))[:len(p):len(p)]
	var i C.R_xlen_t
	for k, v := range p {
		C.SET_STRING_ELT(names, i, C.Rf_mkCharLenCE(C._GoStringPtr(k), C.int(len(k)), C.CE_UTF8))
		if v {
			s[i] = 1
		} else {
			s[i] = 0
		}
		i++
	}
	C.setAttrib(r, packSEXP_types_Basic_string("names"), names)
	C.Rf_unprotect(2)
	return r
`, len(&a{}))
				return
			}
		}

		switch {
		case elem.String() == "error":
			fmt.Fprintf(buf, `	n := len(p)
	r := C.Rf_allocVector(C.%[1]s, C.R_xlen_t(n))
	C.Rf_protect(r)
	names := C.Rf_allocVector(C.STRSXP, C.R_xlen_t(n))
	C.Rf_protect(names)
	var i C.R_xlen_t
	for k, v := range p {
		C.SET_STRING_ELT(names, i, C.Rf_mkCharLenCE(C._GoStringPtr(k), C.int(len(k)), C.CE_UTF8))
		s := C.R_NilValue
		if v != nil {
			C.SET_STRING_ELT(r, i, packSEXP%[2]s(v))
		}
		C.SET_STRING_ELT(r, i, s)
		i++
	}
	C.setAttrib(r, packSEXP_types_Basic_string("names"), names)
	C.Rf_unprotect(2)
	return r
`, rTypeLabelFor(elem), pkg.Mangle(elem))

		default:
			fmt.Fprintf(buf, `	n := len(p)
	r := C.Rf_allocVector(C.%s, C.R_xlen_t(n))
	C.Rf_protect(r)
	names := C.Rf_allocVector(C.STRSXP, C.R_xlen_t(n))
	C.Rf_protect(names)
	var i C.R_xlen_t
	for k, v := range p {
		C.SET_STRING_ELT(names, i, C.Rf_mkCharLenCE(C._GoStringPtr(k), C.int(len(k)), C.CE_UTF8))
		C.SET_VECTOR_ELT(r, i, packSEXP%s(v))
		i++
	}
	C.setAttrib(r, packSEXP_types_Basic_string("names"), names)
	C.Rf_unprotect(2)
	return r
`, rTypeLabelFor(elem), pkg.Mangle(elem))
		}

	case *types.Pointer:
		fmt.Fprintf(buf, `	if p == nil {
		return C.R_NilValue
	}
	return packSEXP%s(*p)
`, pkg.Mangle(typ.Elem()))

	case *types.Slice:
		// TODO(kortschak): Handle named simple types properly.
		elem := typ.Elem()
		if elem, ok := elem.(*types.Basic); ok {
			switch elem.Kind() {
			// TODO(kortschak): Make the fast path available
			// to []T where T is one of these kinds.
			case types.Int32:
				// Maximum length array type for this element type.
				type a [1 << 47]int32
				fmt.Fprintf(buf, `	r := C.Rf_allocVector(C.INTSXP, C.R_xlen_t(len(p)))
	C.Rf_protect(r)
	s := (*[%d]%s)(unsafe.Pointer(C.INTEGER(r)))[:len(p):len(p)]
	copy(s, p)
	C.Rf_unprotect(1)
	return r
`, len(&a{}), nameOf(elem))
				return
			case types.Uint8:
				// Maximum length array type for this element type.
				type a [1 << 49]byte
				fmt.Fprintf(buf, `	r := C.Rf_allocVector(C.RAWSXP, C.R_xlen_t(len(p)))
	C.Rf_protect(r)
	s := (*[%d]%s)(unsafe.Pointer(C.RAW(r)))[:len(p):len(p)]
	copy(s, p)
	C.Rf_unprotect(1)
	return r
`, len(&a{}), nameOf(elem))
				return
			case types.Float64:
				// Maximum length array type for this element type.
				type a [1 << 46]float64
				fmt.Fprintf(buf, `	r := C.Rf_allocVector(C.REALSXP, C.R_xlen_t(len(p)))
	C.Rf_protect(r)
	s := (*[%d]%s)(unsafe.Pointer(C.REAL(r)))[:len(p):len(p)]
	copy(s, p)
	C.Rf_unprotect(1)
	return r
`, len(&a{}), nameOf(elem))
				return
			case types.Complex128:
				// Maximum length array type for this element type.
				type a [1 << 45]complex128
				fmt.Fprintf(buf, `	r := C.Rf_allocVector(C.CPLXSXP, C.R_xlen_t(len(p)))
	C.Rf_protect(r)
	s := (*[%d]%s)(unsafe.Pointer(C.CPLXSXP(r)))[:len(p):len(p)]
	copy(s, p)
	C.Rf_unprotect(1)
	return r
`, len(&a{}), nameOf(elem))
				return
			case types.Bool:
				// Maximum length array type for this element type.
				type a [1 << 47]int32
				// FIXME(kortschak): Does Rf_allocVector return a
				// zeroed vector? If it does, the loop below doesn't
				// need the else clause.
				// Alternatively, convert the []bool to a []byte:
				//  for i, v := range *(*[]byte)(unsafe.Pointer(&p)) {
				//      s[i] = int32(v)
				//  }
				fmt.Fprintf(buf, `	r := C.Rf_allocVector(C.LGLSXP, C.R_xlen_t(len(p)))
	C.Rf_protect(r)
	s := (*[%d]%s)(unsafe.Pointer(C.LOGICAL(r)))[:len(p):len(p)]
	for i, v := range p {
		if v {
			s[i] = 1
		} else {
			s[i] = 0
		}
	}
	C.Rf_unprotect(1)
	return r
`, len(&a{}), nameOf(elem))
				return
			}
		}

		switch {
		case elem.String() == "string":
			fmt.Fprint(buf, `	r := C.Rf_allocVector(C.STRSXP, C.R_xlen_t(len(p)))
	C.Rf_protect(r)
	for i, v := range p {
		s := C.Rf_mkCharLenCE(C._GoStringPtr(string(v)), C.int(len(v)), C.CE_UTF8)
		C.SET_STRING_ELT(r, C.R_xlen_t(i), s)
	}
	C.Rf_unprotect(1)
	return r
`)
		case elem.String() == "error":
			fmt.Fprint(buf, `	r := C.Rf_allocVector(C.STRSXP, C.R_xlen_t(len(p)))
	C.Rf_protect(r)
	for i, v := range p {
		s := C.R_NilValue
		if v != nil {
			s = C.Rf_mkCharLenCE(C._GoStringPtr(v), C.int(len(v)), C.CE_UTF8)
		}
		C.SET_STRING_ELT(r, C.R_xlen_t(i), s)
	}
	C.Rf_unprotect(1)
	return r
`)
		default:
			fmt.Fprintf(buf, `	r := C.Rf_allocVector(C.%s, C.R_xlen_t(len(p)))
	C.Rf_protect(r)
	for i, v := range p {
		C.SET_VECTOR_ELT(r, C.R_xlen_t(i), packSEXP%s(v))
	}
	C.Rf_unprotect(1)
	return r
`, rTypeLabelFor(typ), pkg.Mangle(elem))
		}

	case *types.Struct:
		n := typ.NumFields()
		fmt.Fprintf(buf, "\tr := C.allocList(%d)\n\tC.Rf_protect(r)\n", n)
		fmt.Fprintf(buf, "\tnames := C.Rf_allocVector(C.STRSXP, %d)\n\tC.Rf_protect(names)\n", n)
		fmt.Fprintln(buf, "\targ := r")
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			rName := targetFieldName(typ, i)
			fmt.Fprintf(buf, "\tC.SET_STRING_ELT(names, %d, C.Rf_mkCharLenCE(C._GoStringPtr(`%s`), %d, C.CE_UTF8))\n", i, rName, len(rName))
			fmt.Fprintf(buf, "\tC.SETCAR(arg, packSEXP%s(p.%s))\n", pkg.Mangle(f.Type()), f.Name())
			if i < n-1 {
				fmt.Fprintln(buf, "\targ = C.CDR(arg)")
			}
		}
		fmt.Fprintln(buf, "\tC.setAttrib(r, packSEXP_types_Basic_string(`names`), names)\n\tC.Rf_unprotect(2)\n\treturn r")

	default:
		panic(fmt.Sprintf("unhandled type: %s", typ))
	}
}

var typeLabelTable = map[string]string{
	"logical":   "LGLSXP",
	"integer":   "INTSXP",
	"double":    "REALSXP",
	"complex":   "CPLXSXP",
	"character": "STRSXP",
	"raw":       "RAWSXP",
	"list":      "VECSXP",
}

// rTypeLabelFor returns the R type label for the R atomic type
// corresponding to typ.
func rTypeLabelFor(typ types.Type) string {
	name, _ := rTypeOf(typ)
	label, ok := typeLabelTable[name]
	if !ok {
		return fmt.Sprintf("<%s>", typ)
	}
	return label
}