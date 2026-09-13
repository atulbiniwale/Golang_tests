package main

import "fmt"

// arrays are fixed length data structues, size can not be change after defining. So in Go, we mostly use Slices instead of arrays.

func main() {

	// create arrays
	array1 := [5]int{1, 2, 3, 4, 5} // create an array of 5 integers (same data type)
	array2 := [5]int{}              //empty array of integers
	array3 := [5]string{}           //empty array of strings
	array4 := [5]bool{}             //empty array of booleans
	fmt.Println(array1)

	// slicing arrays
	fmt.Println(array1[0]) // slicing arrays
	fmt.Println(array1[1])
	fmt.Println(array1[1:3]) // returning array
	fmt.Println(len(array1)) // length of array

	// empty arrays of diff data types.
	fmt.Println(array2) // empty array of integers
	fmt.Println(array3) // empty array of strings
	fmt.Println(array4) // empty array of booleans
}
