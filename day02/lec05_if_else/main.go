package main

import (
	"fmt"
	"strings"
)

func main() {
	var value int
	fmt.Scan(&value)

	if value%2 == 0 {
		fmt.Println("The number", value, "is even")
	} else {
		fmt.Println("The number", value, "is odd")
	}

	var color string
	fmt.Scan(&color)
	if strings.Compare(color, "green") == 0 {
		fmt.Println("Color is green")
	} else if strings.Compare(color, "red") == 0 {
		fmt.Println("Color is red")
	} else {
		fmt.Println("Unknown color")
	}

	// Инициализация в блоке условного оператора
	// Блок присвания - только :=
	// Инициализируемая переменная видна ТОЛЬКО внутри области видимости условного оператора.
	if num := 10; num%2 == 0 {
		fmt.Println("EVEN")
	} else {
		fmt.Println("ODD")
	}

	// if data, err := someFunc(); err == nil {
	// 	fmt.Println("EVEN")
	// }

	// Плохо: автоподстановка точки с запятой ; в конце строки может сломать конструкцию - следите за переносами строк
	// var age int = 10
	// if age >= 7 {
	// 	fmt.Println("Go to school")
	// }
	// else {
	// 	fmt.Println("Go to mammy")
	// }

	// НЕ ИДЕОМАТИЧНО
	if width := 100; width > 100 {
		fmt.Println("width > 100")
	} else {
		fmt.Println("width <= 100")
	}
	// Странное правило номкр 1: старайтесь избегать блоков ELSE
	// ИДЕОМАТИЧНО:
	if height := 100; height > 100 {
		fmt.Println("height > 100")
		return
	}

	fmt.Println("height <= 100")
}
