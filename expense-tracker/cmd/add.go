package cmd

import (
	"flag"
	"fmt"
	"os"
	"time"

	"expense-tracker/models"
	"expense-tracker/storage"
)

// Add creates a new expense and saves it.
func Add(args []string) {
	// Create a new flag set for the "add" subcommand
	cmd := flag.NewFlagSet("add", flag.ExitOnError)

	// Define flags for this subcommand
	amount := cmd.Float64("amount", 0, "Amount spent (required)")
	category := cmd.String("category", "", "Category like food, transport (required)")
	note := cmd.String("note", "", "Optional note")

	// Parse the flags from the arguments
	cmd.Parse(args)

	// Validate required fields
	if *amount <= 0 {
		fmt.Fprintln(os.Stderr, "Error: --amount must be greater than 0")
		cmd.Usage()
		os.Exit(1)
	}
	if *category == "" {
		fmt.Fprintln(os.Stderr, "Error: --category is required")
		cmd.Usage()
		os.Exit(1)
	}

	// Load existing expenses
	expenses, err := storage.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading data: %v\n", err)
		os.Exit(1)
	}

	// Create the new expense
	newExpense := models.Expense{
		ID:       storage.NextID(expenses),
		Date:     time.Now(), // current date and time
		Category: *category,
		Amount:   *amount,
		Note:     *note,
	}

	// Add to the slice
	expenses = append(expenses, newExpense)

	// Save back to file
	err = storage.Save(expenses)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error saving data: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Added expense #%d: %.2f for %s\n", newExpense.ID, newExpense.Amount, newExpense.Category)
}
