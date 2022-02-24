package main

import "fmt"

// Struct Declaration

type Dog struct {
	name string
	age  int
}

// Inheritance (Embedding) in Struct*

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	CanFly   bool
	SpeedKMH float32
}

func main() {

	// Maps are unordered

	// Ways to create an map

	m := make(map[string]int)

	m = map[string]int{
		"Tom": 1,
		"Bob": 2,
	}

	fmt.Println(m)

	fmt.Println("")
	/* ------------------------------------------------------ */

	// Add to map

	m["John"] = 1

	fmt.Println(m)

	fmt.Println("")
	/* ------------------------------------------------------ */

	// Delete item in map

	delete(m, "Tom")

	fmt.Println(m)

	fmt.Println("")
	/* ------------------------------------------------------ */

	// Get specific value from map

	fmt.Println(m["John"])

	fmt.Println("")
	/* ------------------------------------------------------ */

	// Check for presence

	person, ok := m["Jessy"]

	fmt.Println(person, ok) // 0 false

	fmt.Println("")
	/* ------------------------------------------------------ */

	// Get len

	fmt.Println(len(m))

	fmt.Println("")
	/* ------------------------------------------------------ */

	// Structs

	// Creating

	// fieldname syntax
	myDog := Dog{
		name: "Bello",
		age:  2,
	}

	// positinal syntax
	// myDog := Dog{
	// 	"Bello",
	// 	2,
	// }

	fmt.Println(myDog)
	fmt.Println(myDog.name)
	fmt.Println(myDog.age)

	fmt.Println("")
	/* ------------------------------------------------------ */

	// Anonymous Struct

	myCat := struct{ name string }{name: "Garfield"}
	fmt.Println(myCat)

	fmt.Println("")
	/* ------------------------------------------------------ */

	// Embedding

	b := Bird{}
	b.Name = "Tweety"
	b.Origin = "USA"
	b.CanFly = true
	b.SpeedKMH = 40
	fmt.Println(b)
}
