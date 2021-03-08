package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		sideA float64
		sideB float64
	)

	fmt.Scan(&sideA, &sideB)

	fmt.Println(math.Pow(sideA+sideB, 2))
}
