package main

import "fmt"

func main() {

	// first way of for loop
	for i := 0; i < 3; i++ {
		println(" line number is", i)
	}

	// second way- using range (iterate over slice, map, strings etc.) e.g. for i := range 7 {...}
	names := []string{"apple", "mango"}
	for index, value := range names {
		println("index:", index, "value:", value)
		for index, char := range value {
			println("index:", index, "char:", string(char))
		}
	}

	// third way- like a while loop (infinite loop with break, continue)
	// break infinite loop with Ctrl + C
	counter := 0
	for {
		println("now counter is", counter)
		counter++
		if counter == 3 {
			break // breaks the infinite loop when counter is 3.
			// similarly using 'continue' instead of break will skip everthing after this & continue loop for next iteration.
		}

		fmt.Printf("iteration %v \n", counter)
	}

}
