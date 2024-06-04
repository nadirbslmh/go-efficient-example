package main

import "testing"

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(10)
	}
}

func BenchmarkFibonacciDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciDP(10)
	}
}
