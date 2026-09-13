package main

import "fmt"

func factorialfn(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorialfn(n-1)

}

func main() {
	n := 4
	fmt.Println(factorialfn(n))

}
