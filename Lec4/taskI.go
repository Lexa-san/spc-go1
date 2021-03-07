package main

import (
	"fmt"
)

func main() {
	var (
		a string
		b string
		c string
	)

	fmt.Scan(&a, &b, &c)

	fmt.Println(c + b + a)
}
