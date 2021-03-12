package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var (
		cppNum  int
		rustNum int
		count   int
	)

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	cppNum, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	rustNum, _ = strconv.Atoi(scanner.Text())

	cppStudents := make([]string, cppNum, cppNum)
	rustStudents := make([]string, rustNum, rustNum)
	//var rustStudent string

	for i := 0; i < cppNum; i++ {
		scanner.Scan()
		cppStudents[i] = scanner.Text()
	}
	for i := 0; i < rustNum; i++ {
		scanner.Scan()
		rustStudents[i] = scanner.Text()
	}

	var flag bool
	for _, cppStudent := range cppStudents {
		flag = false

	inner1:
		for _, rustStudent := range rustStudents {

			if rustStudent == cppStudent {
				flag = true
				//fmt.Println("BINGO", rustStudent)
				break inner1
			}
		}

		if !flag {
			count++
		}
	}

	for _, rustStudent := range rustStudents {
		flag = false

	inner2:
		for _, cppStudent := range cppStudents {

			if rustStudent == cppStudent {
				flag = true
				//fmt.Println("BINGO", rustStudent)
				break inner2
			}
		}

		if !flag {
			count++
		}
	}

	fmt.Println(count)

	//fmt.Println(cppStudents, rustStudents)
}
