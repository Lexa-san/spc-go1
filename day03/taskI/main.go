package main

import (
	"fmt"
)

func main() {
	var (
		n        int
		isSimple bool = true
	)

	fmt.Scan(&n)
	if n == 1 {
		fmt.Println("1")
		return
	}

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Printf("%d ", i)
			if i != 1 && i != n {
				isSimple = false
			}
		}
	}
	fmt.Println("")

	if isSimple {
		fmt.Println("ACHTUNG")
	}
}
