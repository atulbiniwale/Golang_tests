package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// randoms (from math/rand)
	nums := rand.Perm(19) // genarates 19 random numbers in a slice.
	fmt.Println(nums)

}
