package main

import "fmt"

// create interface
type bird interface {
	move() string
}

// create struct 1
type parrot struct{}

func (p parrot) move() string {
	return "parrot flies!"
}

// create struct 2
type duck struct{}

func (d duck) move() string {
	return "duck swims! "
}

// create common interface function which will be used by child ducts with same method signature.
func interfacefunc(b bird) {
	fmt.Printf("\n move function: %v", b.move())
}

func main() {
	var p1 parrot
	var d1 duck

	// here same function gives different outputs for different objects passed into it.
	interfacefunc(p1)
	interfacefunc(d1)
}
