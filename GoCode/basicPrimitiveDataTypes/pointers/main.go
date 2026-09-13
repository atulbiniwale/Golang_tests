package main

func main() {

	num1 := 10

	//pointer declaration and assigning a memory location of another var
	var mypointer *int
	mypointer = &num1

	println("original variable:", num1)
	println("value of pointer:", *mypointer)
	println("address/ memory location of pointer:", mypointer)

	println("************************************")

	*mypointer = 15

	println("original variable is now changed to :", num1)
	println("value of pointer changed to:", *mypointer) // * is used for dereference (get value of pointer).
	println("address/ memory location of pointer remains the same:", mypointer)

}
