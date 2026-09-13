package main

import "fmt"

// map is like a dictionary with key-value pair (also called hashes.)

func main() {

	//create maps
	grades := make(map[string]int)
	grades["atul"] = 90 // adding key-value to the map.
	grades["rishin"] = 95

	//easier way to create maps (inline)
	grade2 := map[string]int{"neha": 85, "b": 89}

	fmt.Println(grade2)
	fmt.Println(grades)
	fmt.Println("Grades of atul", grades["atul"])

	// modify value in place.
	grades["atul"] = 91

	fmt.Println("updated Grades of atul", grades["atul"])

	// adding a key-value to the map.
	grades["crazy"] = 45
	fmt.Println(grades)

	delete(grades, "crazy") // deleting an item in the map.

	fmt.Println("crazy now deleted")
	fmt.Println(grades)
	fmt.Println("**************************************************************")

	// checking if a key exists
	index, Exists := grades["atul"] // Exists function returns true or false.
	fmt.Println("grades of atul", index)
	fmt.Println("does key exists?", Exists)

	// iterate over a map using for loop and := range
	for index, value := range grades {
		fmt.Printf("key is %s and value is %d\n", index, value)
	}

}
