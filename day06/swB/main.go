package main

import (
	"fmt"
	"strings"
)

func main() {

	var n int
	fmt.Scan(&n)

	var suffix string
	var main string
	for i := 1; i <= n; i++ {
		suffix = strings.Repeat(" ", n-i)
		main = strings.Repeat("@", 2*(i-1)+1)
		fmt.Printf("%s%s%s\n", suffix, main, suffix)
	}

}
