package sexp

import (
	"fmt"
	"unsafe"

	"github.com/rgonomic/rgo/sexp"
)

// PrintSEXP prints the SEXP value passed to it and returns it unaltered.
// If the value is an atomic vector, its contents are printed.
func PrintSEXP(p unsafe.Pointer) unsafe.Pointer {
	sxp := (*sexp.Value)(p)
	info := sxp.Info()
	switch sxp := sxp.Interface().(type) {
	case *sexp.Integer:
		fmt.Printf("%s %#v\n", info, sxp.Vector())
	case *sexp.Logical:
		fmt.Printf("%s %#v\n", info, sxp.Vector())
	case *sexp.Real:
		fmt.Printf("%s %#v\n", info, sxp.Vector())
	case *sexp.Complex:
		fmt.Printf("%s %#v\n", info, sxp.Vector())
	case *sexp.String:
		fmt.Printf("%s %q\n", info, sxp.Vector())
	case *sexp.Character:
		fmt.Printf("%s %s\n", info, sxp)
	case *sexp.Raw:
		fmt.Printf("%s %#v\n", info, sxp.Bytes())
	default:
		fmt.Println(info)
	}
	return p
}

// Gophers returns n gophers.
func Gophers(n int) unsafe.Pointer {
	c := sexp.NewString(n).Protect()
	defer c.Unprotect()
	vec := c.Vector()
	for i := range vec {
		vec[i] = sexp.NewCharacter(fmt.Sprintf("Gopher %d", i+1))
	}
	return c.Pointer()
}
