package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	// var x string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if utf8.RuneCountInString(scanner.Text()) == 0 {
			return
		}
		fmt.Println(scanner.Text())
	}
	// for {
	// 	fmt.Scan(&x)
	// 	if utf8.RuneCountInString(x) == 0 {
	// 		return
	// 	}
	// 	fmt.Println(x)
	// }
}
