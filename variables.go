package main

import "fmt"

// Naming conventions and scopes
// Lowercase package scoped variable (var i int) -> any file in the package can access this variable
// Uppercase package scoped variable (var I int) -> exported from the package and globally visible
// Blocked scope olny visible in the block -> example see main func

// Naming conventions
// 1. the name of the variable should reflect the life span of the variable

// Grouped declaration, package scoped
var (
	i int = 5
	x int
	y string = "Hello World"
)

func main() {
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)

	var temp int // declaration without initialization
	fmt.Printf("%v, %T\n", temp, temp)

	var y int = 42 // declaration with initialization
	fmt.Printf("%v, %T\n", y, y)

	var i, x int = 42, 1302 // declare and init multiple vars at once
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", x, x)

	var foo = 42 // type omitted, will be inferred
	fmt.Printf("%v, %T\n", foo, foo)

	k := 42 // shorthand, only in func bodies, omit var keyword, type is always implicit
	fmt.Printf("%v, %T\n", k, k)

	const constant = "This is a constant"
	fmt.Printf("%v, %T", constant, constant)

}
