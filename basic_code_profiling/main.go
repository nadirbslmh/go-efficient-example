package main

var memo map[int]int = map[int]int{}

func main() {

}

//go:noinline
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

//go:noinline
func FibonacciDP(n int) int {
	res, exist := memo[n]

	if exist {
		return res
	}

	if n <= 1 {
		return n
	}

	memo[n] = FibonacciDP(n-1) + FibonacciDP(n-2)
	return memo[n]
}
