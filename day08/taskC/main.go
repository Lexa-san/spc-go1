package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"unicode/utf8"
)

type SmartList struct {
	words []string
}

func New(newWords []string) *SmartList {
	return &SmartList{newWords}
}

//Sort words and printed it
func (sl *SmartList) GetAnswer() {
	sort.Strings(sl.words)

	sort.Slice(sl.words, func(i, j int) bool {
		return utf8.RuneCountInString(sl.words[i]) < utf8.RuneCountInString(sl.words[j])
	})
	for _, word := range sl.words {
		fmt.Println(word)
	}
}

//Add word to words
func (sl *SmartList) AddWord(word string) {
	sl.words = append(sl.words, word)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	SL := New(make([]string, 0, n))

	for i := 0; i < n; i++ {
		scanner.Scan()
		SL.AddWord(scanner.Text())
	}

	SL.GetAnswer()

}
