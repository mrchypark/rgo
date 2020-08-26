# Wrapping a function with multiple return values

Go allows multiple return values from functions. This is handled by `rgo` by wrapping the returned tuple into a named list with the names being extracted from the Go function signature if possible.

The package is generated by defining a `go.mod` file with
```
$ go mod init ioutil
```
running `rgo init` with an argument pointing to the [Go standard library io/ioutil package](https://pkg.go.dev/io/ioutil?tab=doc),
```
$ rgo init io/ioutil
```
and then editing the `rgo.json` file to regex-match the desired functions in the Go sort package, `sort.Float64s`, `sort.Float64sAreSorted` and `sort.SearchFloat64s`. To help the R code generation to properly break word in Go's camel-case naming covention, we also give "Float64s" as a word (otherwise the camel-case splitter gives back "float_64s").
```
{
	"PkgPath": "io/ioutil",
	"AllowedFuncs": "^ReadFile$",
	"Words": null,
	"LicenseDir": "LICENSE",
	"LicensePattern": "(LICEN[SC]E|COPYING)(\\.(txt|md))?$"
}
```

Once this is done, the wrapper code can be generated by running the build subcommand.
```
$ rgo build
```
This will generate the Go, C and R wrapper code for the R package, and collate all the licenses in the source package into the `LicenseDir` directory. At this stage the `DESCRIPTION` file should be edited and non-relevant licenses should be removed.

The package can now be installed.
```
$ R CMD INSTALL .
```

With the package installed, we can now read files into a `raw` vector.

```
> library(ioutil)
> a <- ioutil::read_file("go.mod")
> a
$r0
 [1] 6d 6f 64 75 6c 65 20 69 6f 75 74 69 6c 0a 0a 67 6f 20 31 2e 31 35 0a

$r1
NULL

> rawToChar(a$r0)
[1] "module ioutil\n\ngo 1.15\n"
> ioutil::read_file("missing_file")
$r0
raw(0)

$r1
[1] "open missing_file: no such file or directory"
```

Return values in Go may have names, and these are used by `rgo` to give names to the returned list.

The function signature for `ReadFile` is `func ReadFile(filename string) ([]byte, error)` which is standard Go idiom when returned values' contents are obvious. If the signature were `func ReadFile(filename string) (data []byte, err error)` then the output would look like this.

```
> ioutil::read_file("go.mod")
$data
 [1] 6d 6f 64 75 6c 65 20 69 6f 75 74 69 6c 0a 0a 67 6f 20 31 2e 31 35 0a

$err
NULL
```