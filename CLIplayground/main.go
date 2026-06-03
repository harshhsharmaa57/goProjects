package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
)

const appName = "Go Playground"

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func printSortedMap(m map[string]string) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("  %s → %s\n", k, m[k])
	}
}

func helloWorld() {
	fmt.Println("Hello, World!")
}

func valuesDemo() {
	// Strings
	s := "Go is strongly typed"
	fmt.Println("String:", s)

	// Ints & Floats
	i := 42
	f := 3.14159
	fmt.Printf("Int: %d | Float: %.5f\n", i, f)

	// Booleans
	t, fa := true, false
	fmt.Printf("Booleans: %v and %v | AND: %v | OR: %v\n", t, fa, t && fa, t || fa)
}

func variablesDemo() {
	name := "playground"
	count := 10

	count = 20

	fmt.Printf("name=%s, count=%d\n", name, count)
}

func constantsDemo() {
	const maxRetries = 3
	const greeting = "Welcome"

	fmt.Printf("Constant maxRetries=%d, greeting=%s\n", maxRetries, greeting)
}

func loopsDemo() {
	// Array (value type, fixed size)
	arr := [3]int{10, 20, 30}
	fmt.Println("Array iteration (for range):")
	for i, v := range arr {
		fmt.Printf("  index=%d value=%d\n", i, v)
	}

	// Slice (reference type, dynamic)
	slice := []string{"alpha", "beta", "gamma"}
	fmt.Println("Slice iteration (for range):")
	for i, v := range slice {
		fmt.Printf("  index=%d value=%s\n", i, v)
	}

	// Classic for loop
	fmt.Print("Classic for: ")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func ifElseDemo() {
	val := 7
	if val%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

func switchDemo(mode string) {
	switch mode {
	case "hello":
		helloWorld()
	case "values":
		valuesDemo()
	case "vars":
		variablesDemo()
	case "const":
		constantsDemo()
	case "loops":
		loopsDemo()
	case "ifelse":
		ifElseDemo()
	case "map":
		dictionary := map[string]string{
			"slice":  "A dynamically-sized, flexible view into an array",
			"map":    "An unordered collection of key-value pairs",
			"struct": "A composite data type that groups fields together",
		}
		fmt.Println("Dictionary (sorted by key):")
		printSortedMap(dictionary)
	case "func":
		fmt.Printf("sum(1, 2, 3, 4) = %d\n", sum(1, 2, 3, 4))
		q, err := divide(17, 3)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("divide(17, 3) = %d\n", q)
		}
	case "fib":
		fmt.Println("Fibonacci sequence:")
		for i := 0; i < 10; i++ {
			fmt.Printf("  fib(%d) = %d\n", i, fib(i))
		}
	default:
		fmt.Printf("Unknown mode: %q\n", mode)
	}
}

func printHelp() {
	fmt.Printf(`%s — A CLI tour of Go basics

Usage:
  go run . [mode]

Available modes:
  hello   Print Hello, World!
  values  Print typed values (strings, ints, floats, bools)
  vars    Demonstrate variables (:= vs =)
  const   Demonstrate constants
  loops   Iterate arrays and slices with for / range
  ifelse  Demonstrate if/else branching
  map     Build and print a sorted map
  func    Call variadic sum() and divide()
  fib     Recursive fibonacci sequence
  help    Show this message

Examples:
  go run . hello
  go run . func
  go run . map
`, appName)
}

func main() {
	// ---------- flag package stretch goal ----------
	helpFlag := flag.Bool("help", false, "Show help message")
	flag.Parse()

	if *helpFlag {
		printHelp()
		return
	}

	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("No mode provided. Run with --help for usage.")
		os.Exit(1)
	}

	mode := args[0]
	if mode == "help" {
		printHelp()
		return
	}

	fmt.Printf("=== Mode: %s ===\n", mode)
	switchDemo(mode)
}