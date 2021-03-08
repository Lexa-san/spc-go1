package main

import (
	"fmt"
	"os"
)

func main() {

	var (
		age  int
		name string
	)

	fmt.Scan(&name, &age)

	fmt.Printf("My name is: %s\nMy age is: %d\n", name, age)

	// Для ручного использования потока ввода
	fmt.Fscan(os.Stdin, &age)
	fmt.Println("New age:", age)
	
}
