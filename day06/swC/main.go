package main

import (
	"fmt"
	"strings"
)

func main() {

	var height, width int
	var s string
	fmt.Scan(&height, &width, &s)

	for i := 1; i <= height; i++ {
		if i == 1 || i == height {
			fmt.Println(strings.Repeat(s, width))
			continue
		}

		fmt.Printf("%s%s%s\n", s, strings.Repeat(" ", width-2), s)

	}

}
