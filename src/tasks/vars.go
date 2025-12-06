package tasks

import "fmt"

var x, y = "x", "y"

func GlobalAndLocalVars() (string, string, string) {
	z := "z"
	return x, y, z
}

func EmptyInitialize() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
