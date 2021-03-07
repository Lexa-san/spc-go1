package main

import "fmt"

func main() {
	var (
		word1  string
		word2  string
		word3  string
		word4  string
		author string
	)

	fmt.Scan(&word1, &word2, &word3, &word4, &author)

	fmt.Printf("%s - %s\n", word4, author)
	fmt.Printf("%s - %s\n", word3, author)
	fmt.Printf("%s - %s\n", word2, author)
	fmt.Printf("%s - %s\n", word1, author)
}
