package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

func main() {
	var (
		str            string
		oneSymbolPrice int32 = 23
		fullPrice      int32
		fracPart       int32
		intPart        int32
	)

	fmt.Scan(&str)

	fullPrice = oneSymbolPrice * int32(utf8.RuneCountInString(str))

	if fullPrice < 100 {
		fmt.Printf("%d коп.\n", fullPrice)
		return
	}

	intPart = int32(math.Ceil(float64(fullPrice / 100)))
	fracPart = int32(math.Mod(float64(fullPrice), 100))
	fmt.Printf("%d р. %d коп.\n", intPart, fracPart)
}
