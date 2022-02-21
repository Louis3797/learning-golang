package main

import (
	"fmt"
	//"math"
)

// example for iota

const (
	a = iota
	b = iota
	c = iota

	// same as
	// a = iota
	// b
	// c

	// the value of iota is scoped to this constant block so it resets if we declare a new block
)

const (
	a2 = iota
	b2
	c2
)

func main() {

	// Naming convention
	// If you want to export the constant than make the first char to uppercase and use camelcase for naming'
	// If you wont export the constant than begin with a lowercase character

	// const MyConstant int = 1
	// const myConstant int = 1

	// typed constant
	const myConst int = 1
	fmt.Printf("%v, %T\n", myConst, myConst)

	// Constants need to be initalized by compile time. That means we need to assign a value to every constant

	// Doesnt work
	// const a int

	// Thah also doesnt work because the functions needs to execute
	// const res float64 = math.Sin(1)

	// iota Enumerated constants

	// iota can be used for incrementing numbers, starting from 0
	// the initial value of iota is an int

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)

	fmt.Printf("%v\n", a2)
	fmt.Printf("%v\n", b2)
	fmt.Printf("%v\n", c2)

	fmt.Printf("%v\n", a == a2)
}
