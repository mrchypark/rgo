// Copyright ©2020 The rgonomic Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sexp

import (
	"bytes"
	"unsafe"
)

// List is an R linked list.
type List struct {
	list_sexprec
}

// NewList returns a list with length n.
func NewList(n int) *List {
	return (*List)(allocateList(n))
}

// Protect protects the SEXP value and returns it.
func (v *List) Protect() *List {
	return (*List)(protect(unsafe.Pointer(v)))
}

// Unprotect unprotects the SEXP. It is equivalent to UNPROTECT(1).
func (v *List) Unprotect() {
	if v == nil || v.Value().IsNull() {
		panic("sexp: unprotecting a nil value")
	}
	unprotect(1)
}

// Info returns the information field of the SEXP value.
func (v *List) Info() Info {
	if v == nil {
		return NilValue.Info()
	}
	return *(*Info)(unsafe.Pointer(&v.sxpinfo))
}

// Value returns the generic state of the SEXP value.
func (v *List) Value() *Value {
	return (*Value)(unsafe.Pointer(v))
}

// Attributes returns the attributes of the SEXP value.
func (v *List) Attributes() *List {
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
func (v *List) Pointer() unsafe.Pointer {
	return unsafe.Pointer(v)
}

// Head returns the first element of the list (CAR/lisp in R terminology).
func (v *List) Head() *Value {
	if v == nil || v.Value().IsNull() {
		return nil
	}
	car := (*Value)(unsafe.Pointer(v.list_sxp.carval))
	if car.IsNull() {
		return nil
	}
	return car
}

// Tail returns the remaining elements of the list (CDR/lisp in R terminology).
func (v *List) Tail() *Value {
	if v == nil || v.Value().IsNull() {
		return nil
	}
	cdr := (*Value)(unsafe.Pointer(v.list_sxp.cdrval))
	if cdr.IsNull() {
		return nil
	}
	return cdr
}

// Tag returns the list's tag value.
func (v *List) Tag() *Symbol {
	if v == nil || v.Value().IsNull() {
		return nil
	}
	tag := (*Symbol)(unsafe.Pointer(v.list_sxp.tagval))
	if tag.Value().IsNull() {
		return nil
	}
	return tag
}

// Get returns the the Value associated with the given tag in the list.
func (v *List) Get(tag []byte) *Value {
	curr := v
	for !curr.Value().IsNull() {
		t := curr.Tag()
		if t != nil {
			if bytes.Equal(t.Name().Bytes(), tag) {
				return (*Value)(curr.Pointer())
			}
		}
		tail := curr.Tail()
		if tail, ok := tail.Value().Interface().(*List); ok {
			curr = (*List)(tail.Pointer())
			continue
		}
		break
	}
	return nil
}

// tags returns all the tags for the list. The []string is allocated
// by the Go runtime.
func (v *List) tags() []string {
	var tags []string
	curr := v
	for !curr.Value().IsNull() {
		t := curr.Tag()
		if t != nil {
			tag := t.String()
			tags = append(tags, tag)
		}
		tail := curr.Tail()
		if tail, ok := tail.Value().Interface().(*List); ok {
			curr = (*List)(tail.Pointer())
			continue
		}
		break
	}
	return tags
}

// Lang is an R language object.
type Lang struct {
	list_sexprec
}

// Info returns the information field of the SEXP value.
func (v *Lang) Info() Info {
	if v == nil {
		return NilValue.Info()
	}
	return *(*Info)(unsafe.Pointer(&v.sxpinfo))
}

// Value returns the generic state of the SEXP value.
func (v *Lang) Value() *Value {
	return (*Value)(unsafe.Pointer(v))
}

// Attributes returns the attributes of the SEXP value.
func (v *Lang) Attributes() *List {
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
func (v *Lang) Pointer() unsafe.Pointer {
	return unsafe.Pointer(v)
}

// Head returns the first element of the list (CAR/lisp in R terminology).
func (v *Lang) Head() *Value {
	if v == nil || v.Value().IsNull() {
		return nil
	}
	car := (*Value)(unsafe.Pointer(v.list_sxp.carval))
	if car.IsNull() {
		return nil
	}
	return car
}

// Tail returns the remaining elements of the list (CDR/lisp in R terminology).
func (v *Lang) Tail() *Value {
	if v == nil || v.Value().IsNull() {
		return nil
	}
	cdr := (*Value)(unsafe.Pointer(v.list_sxp.cdrval))
	if cdr.IsNull() {
		return nil
	}
	return cdr
}

// Tag returns the object's tag value.
func (v *Lang) Tag() *Symbol {
	if v == nil || v.Value().IsNull() {
		return nil
	}
	tag := (*Symbol)(unsafe.Pointer(v.list_sxp.tagval))
	if tag.Value().IsNull() {
		return nil
	}
	return tag
}

// Dot is an R pairlist of promises.
type Dot struct {
	list_sexprec
}

// Info returns the information field of the SEXP value.
func (v *Dot) Info() Info {
	if v == nil {
		return NilValue.Info()
	}
	return *(*Info)(unsafe.Pointer(&v.sxpinfo))
}

// Value returns the generic state of the SEXP value.
func (v *Dot) Value() *Value {
	return (*Value)(unsafe.Pointer(v))
}

// Attributes returns the attributes of the SEXP value.
func (v *Dot) Attributes() *List {
	if v == nil {
		return nil
	}
	attr := (*List)(unsafe.Pointer(v.attrib))
	if attr.Value().IsNull() {
		return nil
	}
	return attr
}

// Head returns the first element of the list (CAR/lisp in R terminology).
func (v *Dot) Head() *Value {
	if v == nil || v.Value().IsNull() {
		return nil
	}
	car := (*Value)(unsafe.Pointer(v.list_sxp.carval))
	if car.IsNull() {
		return nil
	}
	return car
}

// Tail returns the remaining elements of the list (CDR/lisp in R terminology).
func (v *Dot) Tail() *Value {
	if v == nil || v.Value().IsNull() {
		return nil
	}
	cdr := (*Value)(unsafe.Pointer(v.list_sxp.cdrval))
	if cdr.IsNull() {
		return nil
	}
	return cdr
}

// Tag returns the object's tag value.
func (v *Dot) Tag() *Symbol {
	if v == nil || v.Value().IsNull() {
		return nil
	}
	tag := (*Symbol)(unsafe.Pointer(v.list_sxp.tagval))
	if tag.Value().IsNull() {
		return nil
	}
	return tag
}

// Symbol is an R name value.
type Symbol struct {
	sym_sexprec
}

// Info returns the information field of the SEXP value.
func (v *Symbol) Info() Info {
	if v == nil {
		return NilValue.Info()
	}
	return *(*Info)(unsafe.Pointer(&v.sxpinfo))
}

// Value returns the generic state of the SEXP value.
func (v *Symbol) Value() *Value {
	return (*Value)(unsafe.Pointer(v))
}

// Attributes returns the attributes of the SEXP value.
func (v *Symbol) Attributes() *List {
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
func (v *Symbol) Pointer() unsafe.Pointer {
	return unsafe.Pointer(v)
}

// Value returns the value of the symbol.
func (v *Symbol) SymbolValue() *Value {
	if v == nil {
		return nil
	}
	val := (*Value)(unsafe.Pointer(v.sym_sxp.value))
	if val.IsNull() {
		return nil
	}
	return val
}

// Name returns the name of the symbol
func (v *Symbol) Name() *Character {
	if v == nil {
		return nil
	}
	name := (*Character)(unsafe.Pointer(v.sym_sxp.pname))
	if name.Value().IsNull() {
		return nil
	}
	return name
}

// String returns a Go string of the symbol name.
// The returned string is allocated by the Go runtime.
func (v *Symbol) String() string {
	return v.Name().String()
}

// Internal returns a pointer if the symbol is a .Internal function.
func (v *Symbol) Internal() *Value {
	if v == nil {
		return nil
	}
	intern := (*Value)(unsafe.Pointer(v.sym_sxp.internal))
	if intern.IsNull() {
		return nil
	}
	return intern
}