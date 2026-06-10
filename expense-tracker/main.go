package main

import (
	"fmt"
	"os"

	"expense-tracker/cmd"
)

func main() {
	// os.Args contains the command-line arguments
	// os.Args[0] is the program name
	// os.Args[1] is the subcommand (add, list, summary, delete)
	// os.Args[2:] are the arguments to that subcommand

	if len(os.Args) < 2 {
		fmt.Println("Expense Tracker — track your spending")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  expense-tracker add    --amount 50 --category food --note 'Lunch'")
		fmt.Println("  expense-tracker list   [--category food]")
		fmt.Println("  expense-tracker summary")
		fmt.Println("  expense-tracker delete <id>")
		fmt.Println()
		fmt.Println("Commands:")
		fmt.Println("  add      Add a new expense")
		fmt.Println("  list     List all expenses")
		fmt.Println("  summary  Show spending by category")
		fmt.Println("  delete   Remove an expense by ID")
		os.Exit(1)
	}

	// Determine which subcommand to run
	switch os.Args[1] {
	case "add":
		cmd.Add(os.Args[2:])
	case "list":
		cmd.List(os.Args[2:])
	case "summary":
		cmd.Summary(os.Args[2:])
	case "delete":
		cmd.Delete(os.Args[2:])
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}
