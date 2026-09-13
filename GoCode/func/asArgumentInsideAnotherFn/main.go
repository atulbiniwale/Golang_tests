package main

import "fmt"

func test1(myfunc func(int) int) {
	fmt.Println(myfunc(7))
}

func callFunc(callable func(int) int) int {
	return callable(10)
}

func doubleNumber(num int) int {
	return 2 * num
}

func tripleNumber(num int) int {
	return 3 * num
}

func main() {

	fmt.Println(" function passed as argument inside another function")
	fmt.Println("**************************************************")

	myfn := func(num int) int {
		return num * num
	}

	test1(myfn)

	value1 := callFunc(doubleNumber)
	value2 := callFunc(tripleNumber)

	fmt.Println("value from doubleNumber function is", value1)
	fmt.Println("value from tripleNumber function is", value2)

}
