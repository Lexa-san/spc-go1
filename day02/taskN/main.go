package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		x1        int32
		y1        int32
		x2        int32
		y2        int32
		isAllowed bool = true
	)

	fmt.Scan(&x1, &y1, &x2, &y2)

	if math.Abs(float64(x1-x2)) > 2 || math.Abs(float64(y1-y2)) > 2 {
		isAllowed = false
	} else if math.Abs(float64(x1-x2))+math.Abs(float64(y1-y2)) != 3 {
		isAllowed = false
	}

	if isAllowed {
		fmt.Println("ДА")
	} else {
		fmt.Println("НЕТ")
	}
}
