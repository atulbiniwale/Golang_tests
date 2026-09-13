package main

import (
	"fmt"
)

// 1. functions returning multiple values
func myfunc1(a int, b int) (int, int) {
	sum := a + b
	difference := a - b
	return sum, difference
}

func main() {
	fmt.Println(myfunc1(5, 4))    // calling function 1
	fmt.Println(myfunc2(2, 3, 4)) // calling function 2
	myfunc3()                     // calling function 3
	myfunc4("hello")              // calling function 4
	fmt.Println(myfunc5())        // calling function 5

}

// 2. variadic fn - fn with variable arguments.
func myfunc2(numbers ...int) int {
	sum := 0
	// here _ is used to ignore error.
	for _, num := range numbers {
		sum = sum + num
	}
	return sum
}

// 3. function with no args, no return
func myfunc3() {
	fmt.Println("This function just runs and logs a message!")
}

// 4. function with no return but takes args.
func myfunc4(a string) {
	fmt.Println("prints the argument", a)
}

// 5. function with no args, but returning something.
func myfunc5() string {
	return "hello"
}
