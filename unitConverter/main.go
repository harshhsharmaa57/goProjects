package main

import (
	"flag"
	"fmt"
	"os"
	"converter/units"
)

func main() {
	from := flag.String("from", "", "Source unit (e.g., m, ft, c, f, kg, lb, l, gal, USD, EUR)")
	to := flag.String("to", "", "Target unit (e.g., m, ft, c, f, kg, lb, l, gal, USD, EUR)")
	value := flag.Float64("value", 0, "Value to convert")
	list := flag.Bool("list", false, "List all supported conversions")

	flag.Parse()

	if *list {
		fmt.Println("Supported conversions:")
		for _, c := range units.ListConversions() {
			fmt.Printf("  %s\n", c)
		}
		return
	}

	if *from == "" || *to == "" {
		fmt.Fprintln(os.Stderr, "Error: --from and --to flags are required")
		flag.Usage()
		os.Exit(1)
	}

	// Convert returns (float64, error) — never panic for user input.
	result, err := units.Convert(*from, *to, *value)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// fmt.Printf for direct output; fmt.Sprintf would return a string.
	fmt.Printf("%.4f %s = %.4f %s\n", *value, *from, result, *to)
}