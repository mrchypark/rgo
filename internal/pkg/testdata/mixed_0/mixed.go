// Code generated by "go generate github.com/rgonomic/rgo/internal/pkg/testdata"; DO NOT EDIT.

package mixed_0

type (
	T  int
	S1 string
)

//{"in":["int"],"out":["int"]}
func Test0(par0 int) (int, int) {
	var res0 int
	var res1 int
	return res0, res1
}

//{"in":["int"],"out":["float64","int"]}
func Test1(par0 int) (res0 float64, res1 int) {
	return res0, res1
}

//{"in":["int"]}
func Test2(par0 int) {
}

//{"in":["github.com/rgonomic/rgo/internal/pkg/testdata/mixed_0.S1","github.com/rgonomic/rgo/internal/pkg/testdata/mixed_0.T","int","string"],"out":["github.com/rgonomic/rgo/internal/pkg/testdata/mixed_0.S1","string"]}
func Test3(par0 T, par1 S1) S1 {
	var res0 S1
	return res0
}
