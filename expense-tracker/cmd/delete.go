package cmd

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"expense-tracker/models"
	"expense-tracker/storage"
)

// Delete removes an expense by its ID.
func Delete(args []string) {
	cmd := flag.NewFlagSet("delete", flag.ExitOnError)
	cmd.Parse(args)

	if cmd.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "Error: expense ID required")
		fmt.Fprintln(os.Stderr, "Usage: delete <id>")
		os.Exit(1)
	}

	idStr := cmd.Arg(0)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: '%s' is not a valid ID\n", idStr)
		os.Exit(1)
	}

	expenses, err := storage.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading data: %v\n", err)
		os.Exit(1)
	}

	found := false
	var newExpenses []models.Expense

	for _, e := range expenses {
		if e.ID == id {
			found = true
			fmt.Printf("Deleted expense #%d: %.2f for %s\n", e.ID, e.Amount, e.Category)
			continue
		}
		newExpenses = append(newExpenses, e)
	}

	if !found {
		fmt.Fprintf(os.Stderr, "Error: no expense found with ID %d\n", id)
		os.Exit(1)
	}

	err = storage.Save(newExpenses)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error saving data: %v\n", err)
		os.Exit(1)
	}
}
