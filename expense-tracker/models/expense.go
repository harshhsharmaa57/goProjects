package models

import (
	"fmt"
	"time"
)

// Expense is a single spending record.
// Think of a struct as a container that groups related data together.
type Expense struct {
	ID       int       // unique number for each expense
	Date     time.Time // when the expense happened
	Category string    // e.g., "food", "transport"
	Amount   float64   // how much money
	Note     string    // optional description
}

// String makes Expense printable in a nice format.
// This method is called automatically when you fmt.Println(anExpense)
func (e Expense) String() string {
	// Format the date as YYYY-MM-DD
	dateStr := e.Date.Format("2006-01-02")
	return fmt.Sprintf("| %-4d | %-10s | %-10s | %8.2f | %-20s |",
		e.ID, dateStr, e.Category, e.Amount, e.Note)
}

// TableHeader returns the column names for our list view.
func TableHeader() string {
	return "| ID   | Date       | Category   |   Amount | Note                 |"
}

// Separator returns a line to make the table look neat.
func Separator() string {
	return "+------+------------+------------+----------+----------------------+"
}