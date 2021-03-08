package main

import (
	"fmt"
)

func main() {
	var x int

	for {
		fmt.Scan(&x)
		if x == 0 {
			return
		}
		fmt.Println(x)
	}
}
