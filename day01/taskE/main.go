package main

import "fmt"

func main() {
	var (
		drink   string
		meal    string
		subMeal string
		time    string
	)

	fmt.Scan(&drink)
	fmt.Scan(&meal)
	fmt.Scan(&subMeal)
	fmt.Scan(&time)

	fmt.Printf("I wanna some '%s', my favorite meal - '%s'. Give me also '%s'. I will come soon... '%s'\n", drink, meal, subMeal, time)
}
