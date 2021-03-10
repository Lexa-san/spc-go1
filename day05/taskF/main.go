package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	word := scanner.Text()
	runeSlice := []rune(word)

	fmt.Println(word)

	for i := 0; ; i++ {
		if len(runeSlice) <= 2 {
			break
		}
		if i%2 == 0 {
			runeSlice = runeSlice[2:]
			fmt.Println(string(runeSlice))
			continue
		}
		runeSlice = runeSlice[:len(runeSlice)-2]
		fmt.Println(string(runeSlice))
	}

}
