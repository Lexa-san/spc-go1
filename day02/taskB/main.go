package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		a  float64
		b  float64
		c  float64
		D  float64
		x1 float64
		x2 float64
	)

	fmt.Scan(&a, &b, &c)

	if a == 0 && b == 0 {
		fmt.Println("корней нет")
		return
	}

	if a == 0 {
		fmt.Println("один корень")
		return
	}

	D = math.Pow(b, 2) - 4*a*c

	if D < 0 {
		fmt.Println("корней нет")
		return
	}

	if D == 0 {
		fmt.Println("один корень")
		return
	}

	if D > 0 {
		x1 = (-b + math.Sqrt(D)) / 2 / a
		x2 = (-b - math.Sqrt(D)) / 2 / a
		if x1 == x2 {
			fmt.Println("один корень")
			return
		}

		fmt.Println("два корня")
		return
	}
}
