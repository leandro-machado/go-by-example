package main

import (
	"fmt"
)

func main() {
	var a = "initial"
	fmt.Println(a, "string variable")

	var b = 1
	fmt.Println(b, "int variable")

	var c, d int = 100, 150
	fmt.Println(c, d, "multiple variables")

	var e int
	fmt.Println(e, "declaring var type is int but not have value, the default value is zero")

	name := "leandro machado"
	fmt.Println(name, ":= is a shortand to declaring and initialization a var")
}
