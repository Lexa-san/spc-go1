package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	word := scanner.Text()
	runeSlice := []rune(word)
	var result string

	for index, oneRune := range runeSlice {
		if index%2 == 0 && oneRune != ' ' {
			result += strings.Repeat(string(oneRune), 3)
		}
	}

	fmt.Println(result)

}
