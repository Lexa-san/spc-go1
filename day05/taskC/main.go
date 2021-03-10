package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var word string
	var lastRune rune
	scanner := bufio.NewScanner(os.Stdin)

	var runeSlice []rune
	for scanner.Scan() {
		word = scanner.Text()
		runeSlice = []rune(word)

		if lastRune == 0 || lastRune == runeSlice[0] {
			lastRune = runeSlice[len(runeSlice)-1]
			continue
		}

		break

	}

	fmt.Println(word)

}
