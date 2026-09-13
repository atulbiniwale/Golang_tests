package main

import "fmt"

func main() {

	// if else if
	var1 := 10

	if var1 > 10 {
		fmt.Println("var1 is greateer than 10")
	} else if var1 < 10 {
		fmt.Println("var1 is less than 10")
	} else {
		fmt.Println("var1 is 10")
	}

	// switch case
	var var2 string

	switch var2 {
	case "melbourne":
		fmt.Println("var2 is Melbourne")
	case "brisbane":
		fmt.Println("var2 is Brisbane")
	default:
		fmt.Println("var2 is not mel or bris, its Sydney by defaault")
	}
}
