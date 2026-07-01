package cli

import (
	"expense-tracker-cli/models"
	"fmt"
	"os"
	"strings"
)

func Home(expense *models.ExpenseStore) {

	if len(os.Args) > 1 {
		command := os.Args[1:]

		execute(expense, command)

		return
	}

	for {
		input, _ := ReadInput(">> ")

		command := strings.Fields(input)

		if command[0] == "exit" {
			break
		}

		execute(expense, command)
	}
}

func execute(expense *models.ExpenseStore, command []string) {
	length := len(command)

	if length == 0 {
		return
	}

	switch command[0] {
	case "add":
		if length < 3 {
			fmt.Println("Error: Incomplete command")
		} else {
			money, err := ConvertInt(command[2])

			if err != nil {
				fmt.Println("Incorrect money input")
			} else {
				expense.AddExpense(command[1], money)
				fmt.Println("Expense added successfully")
			}
		}
	case "view":
		for _, expense := range expense.Expenses {
			fmt.Println(expense.ID, expense.Subject, expense.Money)
		}

	case "del":
		if length < 2 {
			fmt.Println("Error: Incomplete command")
		} else {
			id, err := ConvertInt(command[1])

			if err != nil {
				fmt.Println("Incorrect Expense ID")
			} else {
				expense.DeleteExpense(id)
				fmt.Println("Expense deleted successfully")
			}
		}

	case "help":
		fmt.Println("\nadd <subject> <amount> :- Add expense")
		fmt.Println("view :- Display all expenses")
		fmt.Println("del <id> :- Delete expenese")
		fmt.Println("exit :- Exit")

	case "max":
		max(expense)

	case "min":
		min(expense)

	default:
		fmt.Println("Invalid command: ", command[0])
	}
}

func max(expense *models.ExpenseStore) {
	if len(expense.Expenses) == 0 {
		fmt.Println("No expenses stored")
		return
	}

	maxExpense := expense.MaxExpense()

	fmt.Println("Maximum Expense: ", maxExpense.Money)
	fmt.Println("Subject: ", maxExpense.Subject)
}

func min(expense *models.ExpenseStore) {
	if len(expense.Expenses) == 0 {
		fmt.Println("No expenses stored")
		return
	}

	minExpense := expense.MinExpense()

	fmt.Println("Minimum Expense: ", minExpense.Money)
	fmt.Println("Subject: ", minExpense.Subject)
}
