package main

import (
	"fmt"
)

func main() {
	var (
		a float64
		b float64
	)

	fmt.Scan(&a, &b)

	if int64((a+b))%2 == 0 {
		fmt.Println("ЧЁТНОЕ")
	} else {
		fmt.Println("НЕЧЁТНОЕ")
	}
}
