package main

import (
	"fmt"
	"strconv"
)

func main() {

	var fname string = "atul"
	var age int = 42
	var height float64 = 5.9542
	var isAwesome bool = true

	// some built-in fns using Printf to format output (data type checking, formatting float etc.)
	fmt.Printf("%v is %v years old, and this is %v.\n", fname, age, isAwesome)
	fmt.Printf("height is %.2f feet.\n", height)          // formating the float.
	fmt.Printf("Type of variable height is %T\n", height) // type of variable.

	// data type conversion
	var v2 int = 10
	var v3 string = "sampletext"

	v4 := float64(v2)         // converted int to float64
	v5 := strconv.Itoa(v2)    // converted int to string
	v6, _ := strconv.Atoi(v3) // converted string to int
	fmt.Println(v4, v5, v6)
	fmt.Println(v5)

}
