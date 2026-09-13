package main

import (
	"fmt"
	package1 "myproject/mypackage"
)

func main() {
	fmt.Println("hi from main package- entry point of the program.")

	package1.Myfunc() // calling exported function from another package 'package1'.

}
