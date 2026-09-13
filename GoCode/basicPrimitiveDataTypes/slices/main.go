package main

import "fmt"

//slices are flexible & dynamic data structs which grow/shrink (unlike arrays). They're' built on top of arrays.

func main() {

	// create slices
	slice1 := []int{1, 2, 3, 4, 5}
	empty_slice := []int{}   //empty slice with no elements.
	sl3 := make([]int, 5, 8) // using make to create slice with specific length and capacity.

	// slicing arrays
	fmt.Println(slice1)
	fmt.Println(slice1[0])
	fmt.Println(slice1[0:3])
	fmt.Printf("data type is : %T\n", slice1) // underlying data type of slice is still array.

	// append to a slice.
	slice3 := append(slice1, 10, 20, 30)
	slice4 := append(slice1, sl3...) // ... is used to unpack elements of existing slice before appending.

	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(empty_slice)
	fmt.Println(sl3)

	// you can use for loop thr slice/map:  for index, value := range myslice {.....}

}
