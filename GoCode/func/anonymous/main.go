package main

import "fmt"

// define operate function taking Anomynoys function as a Argument. Ano fun takes int, int and returns int.
func operate(fn func(int, int) int) {
	fmt.Println(fn(10, 20))
}

func main() {
	// creating Anonymous func and Assigning to a variable.
	myfn := func(num int) int {
		return num * num
	}
	// another anonymous function inside a function and calling it there itself immediately.
	x := func(mynum int) int {
		return mynum * mynum
	}(5) // this ( ) calls Anonymous function right there by passing 5 as value inside the function.
	fmt.Println(myfn(4))
	fmt.Println(x)

	// calling operate function inside  main
	operate(func(a, b int) int {
		return a + b
	})
}
