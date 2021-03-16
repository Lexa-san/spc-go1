package main

import "fmt"

type Sequence struct {
	message string
}

func New(newMessage string) *Sequence {
	return &Sequence{newMessage}
}

//Calc max length and return this value
func (s *Sequence) FindMax() int {
	var max, current int

	for _, rn := range s.message {
		if rn == 'о' {
			current++
		}
		if current > max {
			max = current
		}
		if rn != 'о' {
			current = 0
		}
	}
	return max
}

func main() {
	var tmpS string
	fmt.Scan(&tmpS)
	var seq *Sequence = New(tmpS)
	fmt.Println(seq.FindMax())
}
