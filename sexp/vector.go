// Copyright ©2020 The rgonomic Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sexp

import "unsafe"

// base returns the address of the first element of the vector.
func (v *vector_sexprec) base() unsafe.Pointer {
	return add(unsafe.Pointer(v), unsafe.Sizeof(vector_sexprec{}))
}

func add(addr unsafe.Pointer, offset uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(addr) + offset)
}

// Integer is an R integer vector.
type Integer struct {
	vector_sexprec
}

// NewInteger returns an integer vector with length n.
//
// The allocation is made by the R runtime. The returned value may need to
// call its Protect method.
func NewInteger(n int) *Integer {
	return (*Integer)(allocateVector(INTSXP, n))
}

// Protect protects the SEXP value and returns it.
func (v *Integer) Protect() *Integer {
	return (*Integer)(protect(unsafe.Pointer(v)))
}

// Unprotect unprotects the SEXP. It is equivalent to UNPROTECT(1).
func (v *Integer) Unprotect() {
	if v == nil || v.Value().IsNull() {
		panic("sexp: unprotecting a nil value")
	}
	unprotect(1)
}

// Len returns the number of elements in the vector.
func (v *Integer) Len() int {
	if v == nil {
		return 0
	}
	return int(v.vecsxp.length)
}

// Info returns the information field of the SEXP value.
func (v *Integer) Info() Info {
	if v == nil {
		return NilValue.Info()
	}
	return *(*Info)(unsafe.Pointer(&v.sxpinfo))
}

// Value returns the generic state of the SEXP value.
func (v *Integer) Value() *Value {
	return (*Value)(unsafe.Pointer(v))
}

// Attributes returns the attributes of the SEXP value.
func (v *Integer) Attributes() *List {
	if v == nil {
		return nil
	}
	attr := (*List)(unsafe.Pointer(v.attrib))
	if attr.Value().IsNull() {
		return nil
	}
	return attr
}

// Vector returns a slice corresponding to the R vector.
func (v *Integer) Vector() []int32 {
	if v == nil {
		return nil
	}
	n := v.Len()
	return (*[1 << 30]int32)(v.base())[:n:n]
}

// Logical is an R logical vector.
type Logical struct {
	vector_sexprec
}

// NewLogical returns a logical vector with length n.
//
// The allocation is made by the R runtime. The returned value may need to
// call its Protect method.
func NewLogical(n int) *Logical {
	return (*Logical)(allocateVector(LGLSXP, n))
}

// Protect protects the SEXP value and returns it.
func (v *Logical) Protect() *Logical {
	return (*Logical)(protect(unsafe.Pointer(v)))
}

// Unprotect unprotects the SEXP. It is equivalent to UNPROTECT(1).
func (v *Logical) Unprotect() {
	if v == nil || v.Value().IsNull() {
		panic("sexp: unprotecting a nil value")
	}
	unprotect(1)
}

// Len returns the number of elements in the vector.
func (v *Logical) Len() int {
	if v == nil {
		return 0
	}
	return int(v.vecsxp.length)
}

// Info returns the information field of the SEXP value.
func (v *Logical) Info() Info {
	if v == nil {
		return NilValue.Info()
	}
	return *(*Info)(unsafe.Pointer(&v.sxpinfo))
}

// Value returns the generic state of the SEXP value.
func (v *Logical) Value() *Value {
	return (*Value)(unsafe.Pointer(v))
}

// Attributes returns the attributes of the SEXP value.
func (v *Logical) Attributes() *List {
	if v == nil {
		return nil
	}
	attr := (*List)(unsafe.Pointer(v.attrib))
	if attr.Value().IsNull() {
		return nil
	}
	return attr
}

// Vector returns a slice corresponding to the R vector.
func (v *Logical) Vector() []int32 {
	if v == nil {
		return nil
	}
	n := v.Len()
	return (*[1 << 30]int32)(v.base())[:n:n]
}

// Real is an R real vector.
type Real struct {
	vector_sexprec
}

// NewReal returns a real vector with length n.
//
// The allocation is made by the R runtime. The returned value may need to
// call its Protect method.
func NewReal(n int) *Real {
	return (*Real)(allocateVector(REALSXP, n))
}

// Protect protects the SEXP value and returns it.
func (v *Real) Protect() *Real {
	return (*Real)(protect(unsafe.Pointer(v)))
}

// Unprotect unprotects the SEXP. It is equivalent to UNPROTECT(1).
func (v *Real) Unprotect() {
	if v == nil || v.Value().IsNull() {
		panic("sexp: unprotecting a nil value")
	}
	unprotect(1)
}

// Len returns the number of elements in the vector.
func (v *Real) Len() int {
	if v == nil {
		return 0
	}
	return int(v.vecsxp.length)
}

// Info returns the information field of the SEXP value.
func (v *Real) Info() Info {
	if v == nil {
		return NilValue.Info()
	}
	return *(*Info)(unsafe.Pointer(&v.sxpinfo))
}

// Value returns the generic state of the SEXP value.
func (v *Real) Value() *Value {
	return (*Value)(unsafe.Pointer(v))
}

// Attributes returns the attributes of the SEXP value.
func (v *Real) Attributes() *List {
	if v == nil {
		return nil
	}
	attr := (*List)(unsafe.Pointer(v.attrib))
	if attr.Value().IsNull() {
		return nil
	}
	return attr
}

// Vector returns a slice corresponding to the R vector.
func (v *Real) Vector() []float64 {
	if v == nil {
		return nil
	}
	n := v.Len()
	return (*[1 << 30]float64)(v.base())[:n:n]
}

// Complex is an R complex vector.
type Complex struct {
	vector_sexprec
}

// NewComplex returns a complex vector with length n.
//
// The allocation is made by the R runtime. The returned value may need to
// call its Protect method.
func NewComplex(n int) *Complex {
	return (*Complex)(allocateVector(CPLXSXP, n))
}

// Protect protects the SEXP value and returns it.
func (v *Complex) Protect() *Complex {
	return (*Complex)(protect(unsafe.Pointer(v)))
}

// Unprotect unprotects the SEXP. It is equivalent to UNPROTECT(1).
func (v *Complex) Unprotect() {
	if v == nil || v.Value().IsNull() {
		panic("sexp: unprotecting a nil value")
	}
	unprotect(1)
}

// Len returns the number of elements in the vector.
func (v *Complex) Len() int {
	if v == nil {
		return 0
	}
	return int(v.vecsxp.length)
}

// Info returns the information field of the SEXP value.
func (v *Complex) Info() Info {
	if v == nil {
		return NilValue.Info()
	}
	return *(*Info)(unsafe.Pointer(&v.sxpinfo))
}

// Value returns the generic state of the SEXP value.
func (v *Complex) Value() *Value {
	return (*Value)(unsafe.Pointer(v))
}

// Attributes returns the attributes of the SEXP value.
func (v *Complex) Attributes() *List {
	if v == nil {
		return nil
	}
	attr := (*List)(unsafe.Pointer(v.attrib))
	if attr.Value().IsNull() {
		return nil
	}
	return attr
}

// Pointer returns an unsafe pointer to the SEXP value.
func (v *Complex) Pointer() unsafe.Pointer {
	return unsafe.Pointer(v)
}

// Vector returns a slice corresponding to the R vector.
func (v *Complex) Vector() []complex128 {
	if v == nil {
		return nil
	}
	n := v.Len()
	return (*[1 << 30]complex128)(v.base())[:n:n]
}

// String is an R character vector.
type String struct {
	vector_sexprec
}

// NewString returns a character vector with length n.
//
// The allocation is made by the R runtime. The returned value may need to
// call its Protect method.
func NewString(n int) *String {
	return (*String)(allocateVector(STRSXP, n))
}

// Protect protects the SEXP value and returns it.
func (v *String) Protect() *String {
	if v == nil {
		return nil
	}
	return (*String)(protect(unsafe.Pointer(v)))
}

// Unprotect unprotects the SEXP. It is equivalent to UNPROTECT(1).
func (v *String) Unprotect() {
	if v == nil || v.Value().IsNull() {
		panic("sexp: unprotecting a nil value")
	}
	unprotect(1)
}

// Len returns the number of elements in the vector.
func (v *String) Len() int {
	if v == nil {
		return 0
	}
	return int(v.vecsxp.length)
}

// Info returns the information field of the SEXP value.
func (v *String) Info() Info {
	if v == nil {
		return NilValue.Info()
	}
	return *(*Info)(unsafe.Pointer(&v.sxpinfo))
}

// Value returns the generic state of the SEXP value.
func (v *String) Value() *Value {
	return (*Value)(unsafe.Pointer(v))
}

// Attributes returns the attributes of the SEXP value.
func (v *String) Attributes() *List {
	if v == nil {
		return nil
	}
	attr := (*List)(unsafe.Pointer(v.attrib))
	if attr.Value().IsNull() {
		return nil
	}
	return attr
}

// Pointer returns an unsafe pointer to the SEXP value.
func (v *String) Pointer() unsafe.Pointer {
	return unsafe.Pointer(v)
}

// Vector returns a slice corresponding to the R vector.
func (v *String) Vector() []*Character {
	if v == nil {
		return nil
	}
	n := v.Len()
	return (*[1 << 30]*Character)(v.base())[:n:n]
}

// Character is the R representation of a string.
type Character struct {
	vector_sexprec
}

// NewCharacter returns a scalar string corresponding to s.
//
// The allocation is made by the R runtime. The returned value may need to
// call its Protect method.
func NewCharacter(s string) *Character {
	return (*Character)(allocateString(s))
}

// Protect protects the SEXP value and returns it.
func (v *Character) Protect() *Character {
	return (*Character)(protect(unsafe.Pointer(v)))
}

// Unprotect unprotects the SEXP. It is equivalent to UNPROTECT(1).
func (v *Character) Unprotect() {
	if v == nil || v.Value().IsNull() {
		panic("sexp: unprotecting a nil value")
	}
	unprotect(1)
}

// Len returns the number of elements in the vector.
func (v *Character) Len() int {
	if v == nil {
		return 0
	}
	return int(v.vecsxp.length)
}

// Info returns the information field of the SEXP value.
func (v *Character) Info() Info {
	if v == nil {
		return NilValue.Info()
	}
	return *(*Info)(unsafe.Pointer(&v.sxpinfo))
}

// Value returns the generic state of the SEXP value.
func (v *Character) Value() *Value {
	return (*Value)(unsafe.Pointer(v))
}

// Attributes returns the attributes of the SEXP value.
func (v *Character) Attributes() *List {
	if v == nil {
		return nil
	}
	attr := (*List)(unsafe.Pointer(v.attrib))
	if attr.Value().IsNull() {
		return nil
	}
	return attr
}

// Pointer returns an unsafe pointer to the SEXP value.
func (v *Character) Pointer() unsafe.Pointer {
	return unsafe.Pointer(v)
}

// Bytes returns the bytes held by the R SEXP value.
func (v *Character) Bytes() []byte {
	if v == nil {
		return nil
	}
	n := v.Len()
	return (*[1 << 30]byte)(v.base())[:n:n]
}

// String returns a Go string corresponding the the R characters.
// The returned string is allocated by the Go runtime.
func (v *Character) String() string {
	return string(v.Bytes())
}

// Raw is an R raw vector.
type Raw struct {
	vector_sexprec
}

// NewRaw returns a raw vector with length n.
//
// The allocation is made by the R runtime. The returned value may need to
// call its Protect method.
func NewRaw(n int) *Raw {
	return (*Raw)(allocateVector(RAWSXP, n))
}

// Protect protects the SEXP value and returns it.
func (v *Raw) Protect() *Raw {
	return (*Raw)(protect(unsafe.Pointer(v)))
}

// Unprotect unprotects the SEXP. It is equivalent to UNPROTECT(1).
func (v *Raw) Unprotect() {
	if v == nil || v.Value().IsNull() {
		panic("sexp: unprotecting a nil value")
	}
	unprotect(1)
}

// Len returns the number of elements in the vector.
func (v *Raw) Len() int {
	if v == nil {
		return 0
	}
	return int(v.vecsxp.length)
}

// Info returns the information field of the SEXP value.
func (v *Raw) Info() Info {
	if v == nil {
		return NilValue.Info()
	}
	return *(*Info)(unsafe.Pointer(&v.sxpinfo))
}

// Value returns the generic state of the SEXP value.
func (v *Raw) Value() *Value {
	return (*Value)(unsafe.Pointer(v))
}

// Attributes returns the attributes of the SEXP value.
func (v *Raw) Attributes() *List {
	if v == nil {
		return nil
	}
	attr := (*List)(unsafe.Pointer(v.attrib))
	if attr.Value().IsNull() {
		return nil
	}
	return attr
}

// Pointer returns an unsafe pointer to the SEXP value.
func (v *Raw) Pointer() unsafe.Pointer {
	return unsafe.Pointer(v)
}

// Bytes returns the bytes held by the R SEXP value.
func (v *Raw) Bytes() []byte {
	if v == nil {
		return nil
	}
	n := v.Len()
	return (*[1 << 30]byte)(v.base())[:n:n]
}

// Vector is a generic R vector.
type Vector struct {
	vector_sexprec
}

// NewVector returns a generic vector with length n.
//
// The allocation is made by the R runtime. The returned value may need to
// call its Protect method.
func NewVector(n int) *Vector {
	return (*Vector)(allocateVector(VECSXP, n))
}

// Protect protects the SEXP value and returns it.
func (v *Vector) Protect() *Vector {
	return (*Vector)(protect(unsafe.Pointer(v)))
}

// Unprotect unprotects the SEXP. It is equivalent to UNPROTECT(1).
func (v *Vector) Unprotect() {
	if v == nil || v.Value().IsNull() {
		panic("sexp: unprotecting a nil value")
	}
	unprotect(1)
}

// Len returns the number of elements in the vector.
func (v *Vector) Len() int {
	if v == nil {
		return 0
	}
	return int(v.vecsxp.length)
}

// Info returns the information field of the SEXP value.
func (v *Vector) Info() Info {
	if v == nil {
		return NilValue.Info()
	}
	return *(*Info)(unsafe.Pointer(&v.sxpinfo))
}

// Value returns the generic state of the SEXP value.
func (v *Vector) Value() *Value {
	return (*Value)(unsafe.Pointer(v))
}

// Attributes returns the attributes of the SEXP value.
func (v *Vector) Attributes() *List {
	if v == nil {
		return nil
	}
	attr := (*List)(unsafe.Pointer(v.attrib))
	if attr.Value().IsNull() {
		return nil
	}
	return attr
}

// Pointer returns an unsafe pointer to the SEXP value.
func (v *Vector) Pointer() unsafe.Pointer {
	return unsafe.Pointer(v)
}

// Vector returns a slice corresponding to the R vector.
func (v *Vector) Vector() []*Value {
	if v == nil {
		return nil
	}
	n := v.Len()
	return (*[1 << 30]*Value)(v.base())[:n:n]
}

// Expression is an R expression.
type Expression struct {
	vector_sexprec
}

// Len returns the number of elements in the vector.
func (v *Expression) Len() int {
	if v == nil {
		return 0
	}
	return int(v.vecsxp.length)
}

// Info returns the information field of the SEXP value.
func (v *Expression) Info() Info {
	if v == nil {
		return NilValue.Info()
	}
	return *(*Info)(unsafe.Pointer(&v.sxpinfo))
}

// Value returns the generic state of the SEXP value.
func (v *Expression) Value() *Value {
	return (*Value)(unsafe.Pointer(v))
}

// Attributes returns the attributes of the SEXP value.
func (v *Expression) Attributes() *List {
	if v == nil {
		return nil
	}
	attr := (*List)(unsafe.Pointer(v.attrib))
	if attr.Value().IsNull() {
		return nil
	}
	return attr
}

// Pointer returns an unsafe pointer to the SEXP value.
func (v *Expression) Pointer() unsafe.Pointer {
	return unsafe.Pointer(v)
}

// Expression returns a slice corresponding to the R expression.
func (v *Expression) Vector() []*Value {
	if v == nil {
		return nil
	}
	n := v.Len()
	return (*[1 << 30]*Value)(v.base())[:n:n]
}

// WeakReference is an R weak reference.
type WeakReference struct {
	vector_sexprec
}

// Len returns the number of elements in the vector.
func (v *WeakReference) Len() int {
	if v == nil {
		return 0
	}
	return int(v.vecsxp.length)
}

// Info returns the information field of the SEXP value.
func (v *WeakReference) Info() Info {
	if v == nil {
		return NilValue.Info()
	}
	return *(*Info)(unsafe.Pointer(&v.sxpinfo))
}

// Value returns the generic state of the SEXP value.
func (v *WeakReference) Value() *Value {
	return (*Value)(unsafe.Pointer(v))
}

// Attributes returns the attributes of the SEXP value.
func (v *WeakReference) Attributes() *List {
	if v == nil {
		return nil
	}
	attr := (*List)(unsafe.Pointer(v.attrib))
	if attr.Value().IsNull() {
		return nil
	}
	return attr
}

// Pointer returns an unsafe pointer to the SEXP value.
func (v *WeakReference) Pointer() unsafe.Pointer {
	return unsafe.Pointer(v)
}

// WeakReference returns the four elements of an R weak reference.
func (v *WeakReference) Vector() []*Value {
	if v == nil {
		return nil
	}
	n := v.Len()
	return (*[1 << 30]*Value)(v.base())[:n:n]
}
