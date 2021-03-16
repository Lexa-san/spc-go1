package main

import "fmt"

func main() {
	var (
		n int
		m int
	)

	fmt.Scan(&n, &m)

	fmt.Println(combination(n, m))
}

func factorial(x int) uint64 {
	if x == 0 {
		return 1
	}

	return uint64(x) * factorial(x-1)
}

func factorialLimit(x int, limit int, step int) uint64 {
	if step == limit {
		return 1
	}

	return uint64(x) * factorialLimit(x-1, limit, step+1)
}

func combination(n int, m int) int {
	if n < m {
		return 0
	}
	if n == m || m == 0 {
		return 1
	}
	return int(factorialLimit(n, m, 0) / factorial(m))
}
