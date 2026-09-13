package main

import (
	"errors"
	"fmt"
)

// creating a struct type.
type Person struct {
	personFname      string
	personProfession string
}

// assigning methods to struct type using receiver varible (p) which points (using pointer) to Person type/struct.
func (p *Person) Talkfunc() {
	fmt.Println(p.personFname, "says he is a", p.personProfession)
}

// constructor func helps in faster creation of instances of struct/class. Notice the use of pointers here.
func createPerson(firstname, profession string) (*Person, error) {
	// data validation logic added below.
	if firstname == "" || profession == "" {
		return nil, errors.New("firstname and profession are required fields.")
	}
	return &Person{
		personFname:      firstname,
		personProfession: profession,
	}, nil
}

// passing struct values as Arguments to function,
func main() {

	// one way to create object/ instance.
	var person1 Person
	person1.personFname = "atul"
	person1.personProfession = "engineer"

	person1.Talkfunc()

	// another way to create object/ instance.
	person2 := Person{
		personFname:      "rish",
		personProfession: "sportsman",
	}
	person2.Talkfunc()

	// using constructor function to create a instance object.
	person3, err := createPerson("ab", "actor")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(person3.personFname, person3.personProfession)

}
