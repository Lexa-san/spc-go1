package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string

	fmt.Scan(&str)

	if strings.Contains(str, "чат") {
		fmt.Println("БОТ")
		return

	}
	fmt.Println("НЕ БОТ")
}
