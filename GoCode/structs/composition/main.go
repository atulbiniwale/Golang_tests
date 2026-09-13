package main

import "fmt"

type Person struct {
	fname string
	lname string
}

// this is Struct embedding & Composition i.e. Employee struct is composed of other structs.
type Employee struct {
	empid      string
	empDetails Person
}

func main() {
	var emp1 Employee // declaring a instance object.
	emp1.empid = "1001"
	emp1.empDetails = Person{
		fname: "atul",
		lname: "b",
	}

	// another way to create a object with embedded struct.
	emp2 := Employee{
		empid: "101",
		empDetails: Person{
			fname: "a",
			lname: "b:",
		},
	}

	fmt.Println("Emp id", emp1.empid, "is", emp1.empDetails.fname)
	fmt.Println("Emp id", emp2.empid, "is", emp2.empDetails.fname)

}
