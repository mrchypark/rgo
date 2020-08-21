// Copyright ©2019 The rgonomic Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package codegen

import (
	"text/template"
)

func MakevarsTemplate() *template.Template {
	return template.Must(template.New("Makevars").Parse(`# Code generated by rgnonomic/rgo; DO NOT EDIT.

.PHONY: all

CGO_CFLAGS = "$(ALL_CPPFLAGS)"
CGO_LDFLAGS = "$(PKG_LIBS) $(SHLIB_LIBADD) $(LIBR)"

all: go docs

docs:

go:
	rm -f *.h
	CGO_CFLAGS=$(CGO_CFLAGS) CGO_LDFLAGS=$(CGO_LDFLAGS) go build -o $(SHLIB) -buildmode=c-shared ./rgo
`))
}
