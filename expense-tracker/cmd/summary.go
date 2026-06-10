package cmd

import (
	"flag"
	"fmt"
	"os"

	"expense-tracker/storage"
)

// Summary shows total spending grouped by category.
func Summary(args []string) {
	cmd := flag.NewFlagSet("summary", flag.ExitOnError)
	cmd.Parse(args)

	expenses, err := storage.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading data: %v\n", err)
		os.Exit(1)
	}

	if len(expenses) == 0 {
		fmt.Println("No expenses recorded yet.")
		return
	}

	// Map to accumulate totals by category
	// map[string]float64 means: keys are strings, values are float64s
	totals := make(map[string]float64)
	var grandTotal float64

	for _, e := range expenses {
		totals[e.Category] += e.Amount
		grandTotal += e.Amount
	}

	// Print results
	fmt.Println("\n=== Spending Summary ===")
	fmt.Printf("%-15s %12s\n", "Category", "Amount")
	fmt.Println("------------------------------")

	// Sort categories alphabetically for consistent output
	// (maps iterate in random order)
	var categories []string
	for cat := range totals {
		categories = append(categories, cat)
	}

	// Sorting strings alphabetically
	for i := 0; i < len(categories); i++ {
		for j := i + 1; j < len(categories); j++ {
			if categories[i] > categories[j] {
				categories[i], categories[j] = categories[j], categories[i]
			}
		}
	}

	for _, cat := range categories {
		fmt.Printf("%-15s %12.2f\n", cat, totals[cat])
	}

	fmt.Println("------------------------------")
	fmt.Printf("%-15s %12.2f\n", "TOTAL", grandTotal)
}
