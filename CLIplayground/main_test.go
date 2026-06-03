package main

import (
	"fmt"
	"testing"
)

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}
}

func BenchmarkFib10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(10)
	}
}

func BenchmarkFib20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(20)
	}
}

func Example() {
	fmt.Println(sum(1, 2, 3))
	// Output: 6
}