package main

import (
	"encoding/json"
	"expense-tracker-cli/cli"
	"expense-tracker-cli/models"
	"fmt"
	"os"
)

func main() {
	var expenseStore *models.ExpenseStore
	var err error

	expenseStore, err = loadExpenseStoreFromFile("storage/data.json")

	if err != nil {
		expenseStore = models.NewExpenseStore()
	}

	cli.Home(expenseStore)

	if err := saveExpensesToFile(expenseStore, "storage/data.json"); err != nil {
		fmt.Println("Failed to save to file")
	}
}

func saveExpensesToFile(ExpenseStore *models.ExpenseStore, filePath string) error {
	data, err := json.MarshalIndent(ExpenseStore, "", "  ")

	if err != nil {
		return fmt.Errorf("marhsalling failed: %w", err)
	}

	return os.WriteFile(filePath, data, 0666)
}

func loadExpenseStoreFromFile(filePath string) (*models.ExpenseStore, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var ExpenseStore models.ExpenseStore
	err = json.Unmarshal(data, &ExpenseStore)
	if err != nil {
		return nil, err
	}

	return &ExpenseStore, nil
}
