package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("you can't divide by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("lets see Error handling")

	// error handling is done by returning error value from the function.
	ans, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Answer of division is", ans)

	// if we want to ignore error, we can use _ to ignore it.
	// ans, _ := divide(10,0)
	// fmt.Println (ans)
}
