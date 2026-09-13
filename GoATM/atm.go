package main

import (
	"fmt"

	"atulb.com/atm/fileops"
)

func main() {

	var accountBalance float64
	accountBalance = fileops.GetFloatFromFile("balance.txt")

	for {
		ShowMenu()

		var yourChoice int
		fmt.Scan(&yourChoice)
		fmt.Println("Your Choice is:", yourChoice)

		// if your choice is 1/2/3/4 etc ..then execute certain logic. We can use switch-case here as well.
		if yourChoice == 1 {
			fmt.Println("Your current balance is:", accountBalance)
		} else if yourChoice == 2 {
			fmt.Println("How much you want to deposit")
			var newDeposit float64
			fmt.Scan(&newDeposit)
			if newDeposit <= 0 {
				fmt.Println("Invalid amount. Deposit must be greater than 0")
				continue // continue keyword skips the code after this line and starts from the top/ beginning of for loop.
			}
			fmt.Println("You have deposited", newDeposit)
			accountBalance = accountBalance + newDeposit
			fmt.Println("Balance Updated to:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, "balance.txt")

		} else {
			fmt.Println("Goodbye!")
			break // break keyword breaks out of loop and reaches the line after the loop ends.
		}

	}

	fmt.Println("Thanks for choosing our ATM!")

}
