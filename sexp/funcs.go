// Copyright ©2020 The rgonomic Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sexp

import "unsafe"

// Promise is an R promise.
type Promise struct {
	prom_sexprec
}

// Info returns the information field of the SEXP value.
func (v *Promise) Info() Info {
	if v == nil {
		return NilValue.Info()
	}
	return *(*Info)(unsafe.Pointer(&v.sxpinfo))
}

// Value returns the generic state of the SEXP value.
func (v *Promise) Value() *Value {
	return (*Value)(unsafe.Pointer(v))
}

// Attributes returns the attributes of the SEXP value.
func (v *Promise) Attributes() *List {
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
func (v *Promise) Pointer() unsafe.Pointer {
	return unsafe.Pointer(v)
}

// Value is value of the promise.
func (v *Promise) PromiseValue() *Value {
	if v == nil || v.Value().IsNull() {
		return nil
	}
	val := (*Value)(unsafe.Pointer(v.prom_sxp.value))
	if val.IsNull() {
		return nil
	}
	return val
}

// Expression is the expression to be evaluated.
func (v *Promise) Expression() *Value {
	if v == nil || v.Value().IsNull() {
		return nil
	}
	expr := (*Value)(unsafe.Pointer(v.prom_sxp.expr))
	if expr.IsNull() {
		return nil
	}
	return expr
}

// Environment returns the environment in which to evaluate the expression.
func (v *Promise) Environment() *Value {
	if v == nil || v.Value().IsNull() {
		return nil
	}
	env := (*Value)(unsafe.Pointer(v.prom_sxp.env))
	if env.IsNull() {
		return nil
	}
	return env
}

// Closure is an R closure.
type Closure struct {
	clo_sexprec
}

// Info returns the information field of the SEXP value.
func (v *Closure) Info() Info {
	if v == nil {
		return NilValue.Info()
	}
	return *(*Info)(unsafe.Pointer(&v.sxpinfo))
}

// Value returns the generic state of the SEXP value.
func (v *Closure) Value() *Value {
	return (*Value)(unsafe.Pointer(v))
}

// Attributes returns the attributes of the SEXP value.
func (v *Closure) Attributes() *List {
	if v == nil {
		return nil
	}
	attr := (*List)(unsafe.Pointer(v.attrib))
	if attr.Value().IsNull() {
		return nil
	}
	return attr
}

// Formals returns the formal arguments of the function.
func (v *Closure) Formals() *Value {
	if v == nil || v.Value().IsNull() {
		return nil
	}
	formals := (*Value)(unsafe.Pointer(v.clos_sxp.formals))
	if formals.IsNull() {
		return nil
	}
	return formals
}

// Body returns the body of the function.
func (v *Closure) Body() *Value {
	if v == nil || v.Value().IsNull() {
		return nil
	}
	body := (*Value)(unsafe.Pointer(v.clos_sxp.body))
	if body.IsNull() {
		return nil
	}
	return body
}

// Environment returns the environment in which to evaluate the function.
func (v *Closure) Environment() *Value {
	if v == nil || v.Value().IsNull() {
		return nil
	}
	env := (*Value)(unsafe.Pointer(v.clos_sxp.env))
	if env.IsNull() {
		return nil
	}
	return env
}

// Environment is a current execution environment.
type Environment struct {
	env_sexprec
}

// Info returns the information field of the SEXP value.
func (v *Environment) Info() Info {
	if v == nil {
		return NilValue.Info()
	}
	return *(*Info)(unsafe.Pointer(&v.sxpinfo))
}

// Value returns the generic state of the SEXP value.
func (v *Environment) Value() *Value {
	return (*Value)(unsafe.Pointer(v))
}

// Attributes returns the attributes of the SEXP value.
func (v *Environment) Attributes() *List {
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
func (v *Environment) Pointer() unsafe.Pointer {
	return unsafe.Pointer(v)
}

// Frame returns the current frame.
func (v *Environment) Frame() *Value {
	if v == nil || v.Value().IsNull() {
		return nil
	}
	frame := (*Value)(unsafe.Pointer(v.env_sxp.frame))
	if frame.IsNull() {
		return nil
	}
	return frame
}

// Enclosing returns the enclosing environment.
func (v *Environment) Enclosing() *Value {
	if v == nil || v.Value().IsNull() {
		return nil
	}
	enc := (*Value)(unsafe.Pointer(v.env_sxp.enclos))
	if enc.IsNull() {
		return nil
	}
	return enc
}

// HashTable returns the environment's hash table.
func (v *Environment) HashTable() *Value {
	if v == nil || v.Value().IsNull() {
		return nil
	}
	tbl := (*Value)(unsafe.Pointer(v.env_sxp.hashtab))
	if tbl.IsNull() {
		return nil
	}
	return tbl
}

// Builtin is an R language built-in function.
type Builtin struct {
	prim_sexprec
}

// Info returns the information field of the SEXP value.
func (v *Builtin) Info() Info {
	if v == nil {
		return NilValue.Info()
	}
	return *(*Info)(unsafe.Pointer(&v.sxpinfo))
}

// Value returns the generic state of the SEXP value.
func (v *Builtin) Value() *Value {
	return (*Value)(unsafe.Pointer(v))
}

// Attributes returns the attributes of the SEXP value.
func (v *Builtin) Attributes() *List {
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
func (v *Builtin) Pointer() unsafe.Pointer {
	return unsafe.Pointer(v)
}

// Offset returns the offset into the table of language primitives.
func (v *Builtin) Offset() int32 {
	return int32(v.prim_sxp.offset)
}

// Special is an R language built-in function.
type Special struct {
	prim_sexprec
}

// Info returns the information field of the SEXP value.
func (v *Special) Info() Info {
	if v == nil {
		return NilValue.Info()
	}
	return *(*Info)(unsafe.Pointer(&v.sxpinfo))
}

// Value returns the generic state of the SEXP value.
func (v *Special) Value() *Value {
	return (*Value)(unsafe.Pointer(v))
}

// Attributes returns the attributes of the SEXP value.
func (v *Special) Attributes() *List {
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
func (v *Special) Pointer() unsafe.Pointer {
	return unsafe.Pointer(v)
}

// Offset returns the offset into the table of language primitives.
func (v *Special) Offset() int32 {
	return int32(v.prim_sxp.offset)
}
