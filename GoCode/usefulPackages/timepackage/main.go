package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("current date:", time.Now())

	current_time := time.Now()

	formatted_currenttime := current_time.Format("02-01-2026")
	fmt.Println("formatted current date:", formatted_currenttime)

	//add more days in current time
	new_date := current_time.AddDate(0, 2, 10)

	fmt.Println("new date after adding few days in current date:", new_date)

	fmt.Println("line 1 printed")
	time.Sleep(5 * time.Second) // add time delay
	fmt.Println("line 2 printed")

	// time ticker
	ticker := time.NewTicker(3 * time.Second)

	for range ticker.C {
		println(ticker)
	}

}
