package main

import "fmt"

// creating a function here
func testfun(x int) {
	fmt.Println(" some number", x)
}

func main() {
	fmt.Println("First class functions - functions stored in a variable")

	x := testfun // function is stored in variable x.
	x(5)         // function is called using just the variable.

	// anonymous function inside a function (inside main in this case)
	var1 := func() {
		fmt.Println(" another anonymous function")
	}
	var1() // calling a function stored in a variable.
}
