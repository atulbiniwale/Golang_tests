package main

import (
	"fmt"
	"strings"
)

// declaring custom types
type Email string // can also be used for float64/bool/int/struct/array/slice/maps/any other custom type.

// adding methods to custom types declared above.
func (e Email) emailDomainFunc() string {
	parts := strings.Split(string(e), "@")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""

}

func main() {

	myemail := Email("john@gmail.com")     // create the object of custom type Email.
	fmt.Println(myemail.emailDomainFunc()) // use the method of custom type Email created above.

}
