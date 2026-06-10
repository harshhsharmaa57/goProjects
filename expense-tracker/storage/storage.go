package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"expense-tracker/models"
)

// FileName is where we save our data.
// It's a constant — it never changes while the program runs.
const FileName = "expenses.json"

// Load reads all expenses from the JSON file.
// If the file doesn't exist yet, it returns an empty slice (not an error).
func Load() ([]models.Expense, error) {
	// Check if file exists
	_, err := os.Stat(FileName)
	if os.IsNotExist(err) {
		// File doesn't exist yet — that's fine for a first run
		return []models.Expense{}, nil
	}

	// Open the file for reading
	file, err := os.Open(FileName)
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %w", err)
	}
	// CRITICAL: Always close files when done.
	// defer means "run this line when the function finishes"
	defer file.Close()

	// Create an empty slice to hold the data
	var expenses []models.Expense

	// decoder reads JSON and turns it into Go structs
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&expenses)
	if err != nil {
		return nil, fmt.Errorf("cannot read JSON: %w", err)
	}

	return expenses, nil
}

// Save writes all expenses to the JSON file.
func Save(expenses []models.Expense) error {
	// os.Create opens the file, or creates it if missing.
	// WARNING: It also DELETES everything in the file first!
	file, err := os.Create(FileName)
	if err != nil {
		return fmt.Errorf("cannot create file: %w", err)
	}
	defer file.Close()

	// encoder turns Go structs into JSON
	encoder := json.NewEncoder(file)
	// Indent makes the JSON pretty (with newlines and spaces)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(expenses)
	if err != nil {
		return fmt.Errorf("cannot write JSON: %w", err)
	}

	return nil
}

// NextID finds the highest ID and adds 1.
// This ensures every new expense gets a unique number.
func NextID(expenses []models.Expense) int {
	max := 0
	for _, e := range expenses {
		if e.ID > max {
			max = e.ID
		}
	}
	return max + 1
}
