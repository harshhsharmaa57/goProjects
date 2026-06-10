package cmd

import (
	"flag"
	"fmt"
	"os"

	"expense-tracker/models"
	"expense-tracker/storage"
)

// List shows all expenses in a table.
func List(args []string) {
	cmd := flag.NewFlagSet("list", flag.ExitOnError)

	// Optional filter by category
	categoryFilter := cmd.String("category", "", "Filter by category")

	cmd.Parse(args)

	// Load data
	expenses, err := storage.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading data: %v\n", err)
		os.Exit(1)
	}

	// Filter if requested
	var filtered []models.Expense
	for _, e := range expenses {
		if *categoryFilter == "" || e.Category == *categoryFilter {
			filtered = append(filtered, e)
		}
	}

	if len(filtered) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	// Print table
	fmt.Println(models.Separator())
	fmt.Println(models.TableHeader())
	fmt.Println(models.Separator())

	for _, e := range filtered {
		fmt.Println(e) // automatically uses the String() method!
	}

	fmt.Println(models.Separator())
	fmt.Printf("Total: %d expenses\n", len(filtered))
}
