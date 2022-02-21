package main

import (
	"fmt"
	"strconv"
)

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr -> unsigned integer

// byte -> alias for uint8

// rune -> alias for int32 ~= a character (Unicode code point) - very Viking

// float32 float64

// complex64 complex128

func main() {

	// booleans
	var i bool = true
	fmt.Println(i)
	i = false
	fmt.Println(i)
	fmt.Println("")

	// integer
	a := 10
	b := 3

	fmt.Println(a + b) // 13
	fmt.Println(a - b) // 7
	fmt.Println(a * b) // 30
	fmt.Println(a / b) // 3
	fmt.Println(a % b) // 1

	var c int8 = 3
	//fmt.Println(a + c) error
	fmt.Println(a + int(c)) // this works fine

	// a = 10 -> 1010
	// b = 3  -> 0011
	fmt.Println(a & b)  // 2 -> bitwise and, output in bit = 0010 = 2
	fmt.Println(a | b)  // 11 -> logical not, output int bit = 1011 = 11
	fmt.Println(a ^ b)  // 9 -> bitwise xor, output int bit = 1001 = 9
	fmt.Println(a &^ b) // 8 -> bit clear, output int bit = 0100 = 8

	fmt.Println(a << b) // 80 -> bit shift left
	fmt.Println(a >> b) // 1 -> bit shift right

	fmt.Println("")
	// float
	n := 3.
	fmt.Printf("%v, %T\n", n, n)

	n = 3.14
	fmt.Printf("%v, %T\n", n, n)

	n = 2.1e14
	fmt.Printf("%v, %T\n\n", n, n)

	// complex

	x := 2i

	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", real(x), real(x))
	fmt.Printf("%v, %T\n\n", imag(x), imag(x))

	// string

	s := "Hello World"

	fmt.Println(s)
	fmt.Printf("%v, %T\n", s, s)

	var foo int = 60
	fmt.Println(string(foo)) // output = < because its converted to ascii
	// to convert it to a string we need this function
	fmt.Println(strconv.Itoa(foo)) // output = 60

	var p string = "!"
	fmt.Println(s + p)

	fmt.Println("")

	// rune

	r := 'a'
	fmt.Printf("%v, %T\n", r, r)
}
