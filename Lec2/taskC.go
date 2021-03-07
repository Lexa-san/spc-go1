package main

import "fmt"

func main() {

	var (
		word1 string = "имя"
		word2 string = "твое"
		word3 string = "мне"
		word4 string = "знакомо"
	)

	fmt.Printf("%s %s %s %s\n", word4, word3, word2, word1)
	fmt.Printf("%s %s %s %s\n", word3, word1, word4, word2)
	fmt.Printf("%s %s %s %s\n", word2, word4, word1, word3)

}
