package main

import "fmt"

func main() {
	var (
		sideA float32
		sideB float32
	)

	fmt.Scan(&sideA, &sideB)

	fmt.Println("Периметр:", (sideA+sideB)*2)
	fmt.Println("Площадь:", sideA*sideB)
}
