package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	var investment float64
	var yrs float64
	var irr float64
	var futurevalue float64

	fmt.Println("Investment calc - printed from app.go file > using package main > inside main function")
	fmt.Println("**************************************")

	fmt.Println("Enter your investment amount:")
	fmt.Scan(&investment)

	fmt.Println("Enter time period:")
	fmt.Scan(&yrs)

	fmt.Println("Enter rate of return:")
	fmt.Scan(&irr)

	fmt.Println("\nThe result is :")
	futurevalue = float64(investment) * math.Pow(1+irr/100, yrs)
	fmt.Printf("Our investment: %v\n", investment)
	fmt.Printf("Time period %v\n", yrs)
	fmt.Printf("IRR: %v \n", irr)
	fmt.Printf("Calculated Future Value would be: %v \n", futurevalue)

	time.Sleep(5 * time.Second)
	fmt.Println("**************************************")
	time.Sleep(2 * time.Second)

}
