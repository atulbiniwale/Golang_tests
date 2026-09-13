package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Enter your input:")

	//method 1: using Scan method from fmt package.
	// Scan method takes user input and stores in variable. Watch for'&' here.
	// Scan methods only reads first word of input (till white space). To read whole user input, we  bufio package.
	var user_input1 string
	fmt.Scan(&user_input1)
	fmt.Printf("you entered: %v\n", user_input1)

	// method 2: using ReadString method from bufio package.
	reader := bufio.NewReader(os.Stdin)
	user_input2, _ := reader.ReadString('\n')
	fmt.Printf("this time you entered: %v", user_input2)

	//ReadString method reads whole input incl. spaces until user hits enter key. It returns the input and error. ( we are ignoring the error here by using _).

}
