package main

import "fmt"

func main() {

	i := 1

	// Basic

	if i > 10 {
		fmt.Println(i)
	} else if i == 10 {
		fmt.Println(i, "is 10")
	} else {
		fmt.Println(i, "is less than 10")
	}

	// if with statement before condition
	if i = 10; i > 10 {
		fmt.Println(i)
	} else if i == 10 {
		fmt.Println(i, "is equal 10")
	} else {
		fmt.Println(i, "is less than 10")
	}

	fmt.Println("")
	/* ------------------------------------------------------ */

	// Switch

	os := "Windows"
	switch os {
	case "darwin":
		fmt.Println("Mac OS Hipster")
		// cases break automatically, no fallthrough by default
	case "linux":
		fmt.Println("Linux Geek")
	default:
		fmt.Println("Other")
	}

	// Comparisons in switch cases
	number := 42
	switch {
	case number < 42:
		fmt.Println("Smaller")
	case number == 42:
		fmt.Println("Equal")
	case number > 42:
		fmt.Println("Greater")
	}

	fmt.Println("")
	/* ------------------------------------------------------ */

	// Type Switch

	var foo interface{} = 1

	switch foo.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("Float64")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("other type")
	}

}
