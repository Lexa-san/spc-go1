package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	numberSlice := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numberSlice[i])
	}
	var a, b int
	fmt.Scan(&a, &b)

	var sum int
	for _, val := range numberSlice[a-1 : b] {
		sum += val
	}
	fmt.Println(sum)
}
