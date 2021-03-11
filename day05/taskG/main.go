package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)

	var count int
	var max int
	for _, char := range str {
		if char == 'о' {
			count++
			if count > max {
				max = count
			}
			continue
		}

		if count > 0 {
			count = 0
		}

	}
	fmt.Println(max)
}
