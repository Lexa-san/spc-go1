package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var str string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str = scanner.Text()
	//fmt.Println(str)

	runeSlice := []rune(str)
	//fmt.Println(runeSlice)

	var newStr1 = fmt.Sprintf("%c%c", runeSlice[0], runeSlice[len(runeSlice)-1])
	var newStr2 = fmt.Sprintf("%c%c", runeSlice[len(runeSlice)-1], runeSlice[0])
	//fmt.Println(newStr1)

	switch newStr1 {
	case "да", "дА", "Да", "ДА":
		fmt.Println("СОГЛАСЕН")
		return
	default:
		switch newStr2 {
		case "да", "дА", "Да", "ДА":
			fmt.Println("СОГЛАСЕН")
			return
		default:
			fmt.Println("НЕ СОГЛАСЕН")
			return
		}
		fmt.Println("НЕ СОГЛАСЕН")
		return
	}
}
