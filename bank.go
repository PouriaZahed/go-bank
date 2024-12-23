package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64 , error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000 , errors.New("failed to parse stored balance value")
	}

	return balance , nil
}

func writeBalanceToFile(balance float64)  {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText) , 0644)
}

func main()  {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("---------------")
		// panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposite money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int 
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is: ", accountBalance)
	} else if choice == 2 {
		fmt.Print("Your deposit: ") 
		var depositeAmount float64
		fmt.Scan(&depositeAmount)

		if depositeAmount <= 0 {
			fmt.Println("Invalid amount.Must be greater than 0.")
			continue
		}

		accountBalance += depositeAmount
		fmt.Println("Balance ypdated! New Amount:", accountBalance)
		writeBalanceToFile(accountBalance)
	} else if choice == 3 {
		fmt.Println("How much do you want to withdraw!: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)

		if withdrawAmount <= 0 {
			fmt.Println("Invalid amount.Must be greater than 0.")
			continue

		}

		if withdrawAmount > accountBalance {
			fmt.Println("Invalid amount.You can't withdraw more than you have.")
			continue

		}

		accountBalance -= withdrawAmount
		fmt.Println("Balance ypdated! New Amount:", accountBalance)
		writeBalanceToFile(accountBalance)
	} else {
		fmt.Println("GoodBye!")
		break
	}
	}
	
	fmt.Println("Thanks for choosing our bank")
}