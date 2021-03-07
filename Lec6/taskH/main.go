package main

import (
	"fmt"
)

func main() {
	var (
		grainCount  int
		damageNum   int
		damageDenum int
		numerator   int
		denominator int
		lcm         int
		gcd         int
	)

	fmt.Scan(&grainCount)

	for i := 0; i < grainCount; i++ {
		fmt.Scan(&damageNum, &damageDenum)

		// in the First iteration
		if numerator == 0 && denominator == 0 {
			numerator = damageNum
			denominator = damageDenum
			continue
		}

		// in another case
		lcm = calcLCM(denominator, damageDenum)
		numerator = int(lcm/denominator)*numerator + int(lcm/damageDenum)*damageNum
		denominator = lcm
	}

	// fmt.Printf("%d/%d\n", numerator, denominator)
	gcd = calcGCD(numerator, denominator)
	if gcd > 1 {
		numerator /= gcd
		denominator /= gcd
	}
	fmt.Printf("%d/%d\n", numerator, denominator)
}

// Calculate the Least Common Multiple: (x*y)/gcd(x,y)
func calcLCM(x1 int, y1 int) int {
	return int(x1 * y1 / calcGCD(x1, y1))
}

// Calculate the Greatest Common Denominator
func calcGCD(x int, y int) int {
	for x != 0 && y != 0 {
		if x > y {
			x = x % y
			continue
		}
		y = y % x
	}
	return x + y
}
