package main

import (
	"fmt"
	"strings"
)

func main() {
	// loop
	for i := 0; i <= 5; i++ {
		fmt.Printf("Current value: %d\n", i)
	}

	// break
	for i := 0; i <= 15; i++ {
		if i > 12 {
			break
		}
		fmt.Printf("Current value: %d\n", i)
	}
	fmt.Println("After for loop with BREAK")

	// continue
	for i := 0; i <= 15; i++ {
		if i > 5 && i <= 12 {
			continue
		}
		fmt.Printf("Current value: %d\n", i)
	}
	fmt.Println("After for loop with CONTINUE")

	// included loop
	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("See triangle")

	// sometimes it is going wrong
outer:
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 3; j++ {
			fmt.Printf("i:%d and j:%d and sum i+j=%d\n", i, j, i+j)

			// хочу, чтобы все циклы остановились, включая и внешний цикл
			if i == j {
				break outer
			}
		}
	}

	// loop modifications
	var loopVar1 int = 0
	for loopVar1 < 10 {
		fmt.Printf("Current value: %d\n", loopVar1)
		loopVar1++
	}
	// infinity loop
	var password string
outer2:
	for {
		fmt.Print("Enter password: ")
		fmt.Scan(&password)
		if strings.Contains(password, "1234") {
			fmt.Println("Weak password. Try again")
			continue
		} else {
			fmt.Println("Password accepted")
			break outer2
		}

	}

	// множественные переменные циклв
	for x, y := 0, 1; x <= 10 && y <= 12; x, y = x+1, y+2 {
		fmt.Printf("%d + %d = %d\n", x, y, x+y)
	}
}
