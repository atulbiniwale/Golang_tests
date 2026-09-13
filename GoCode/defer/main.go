package main

import "fmt"

func main() {
	fmt.Println("start of code")

	// note the order of multiple defer statements. Its LIFO (last in - first out)

	defer fmt.Println("Defer1. this is middle of code, but because of defer it's executed at the last.")

	defer fmt.Println("Defer2. this is middle of code, but because of defer it's executed at the last.")

	defer fmt.Println("Defer3.this is middle of code, but because of defer it's executed at the last.")

	fmt.Println("end of code")

}
