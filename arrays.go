package main

import "fmt"

func main() {

	// declartion
	var arr [3]int

	fmt.Printf("%v, %T\n", arr, arr)

	fmt.Println("")
	/* ------------------------------------------------------ */

	// initialize array

	var a = [...]int{1, 2, 3, 4} // Compiler figures out array length

	// var a = [4]int{1, 2, 3, 4} with specific length

	//a := []int{1, 2, 3, 4}	shorthand
	fmt.Println(a)

	fmt.Println("")
	/* ------------------------------------------------------ */

	// assign value to an index
	a[0] = 100
	fmt.Println(a)

	fmt.Println("")
	/* ------------------------------------------------------ */

	// get length
	fmt.Println(len(a))

	fmt.Println("")
	/* ------------------------------------------------------ */

	// matrix
	matrix := [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	// var matrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}} that works also

	fmt.Println(matrix)

	fmt.Println("")
	/* ------------------------------------------------------ */

	// array are values
	// arrays dont point to its values in golang
	foo := [...]int{1, 2, 3}
	b := foo

	b[0] = 100
	fmt.Println("foo:", foo)
	fmt.Println("b:", b)

	fmt.Println("")
	/* ------------------------------------------------------ */

	// get capacity
	test := [100]int{1, 2, 3}
	fmt.Println(test)
	fmt.Println(len(test))
	fmt.Println(cap(test))

	fmt.Println("")
	/* ------------------------------------------------------ */

	// slices
	s := []int{1, 2, 3, 4}
	temp := s[:] // slice of all elements
	fmt.Println(temp)
	temp = s[1:4] // slice from index 1 to 3
	fmt.Println(temp)
	temp = s[:3] // missing low index implies 0
	fmt.Println(temp)
	temp = s[3:]
	fmt.Println(temp)

	fmt.Println("")
	/* ------------------------------------------------------ */

	// create a slice with make
	m := make([]int, 5, 5) // first arg length, second capacity
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))

	fmt.Println("")
	/* ------------------------------------------------------ */

	// append
	array := []int{1, 2}
	fmt.Println(array)
	array = append(array, 3, 4)
	fmt.Println(array)

	// concatenate slices
	x := append(array, s...) // ... -> spread operator
	fmt.Println(x)
}
