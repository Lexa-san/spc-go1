package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var (
		cppNum  int
		rustNum int
		count   int
	)

	scanner.Scan()
	cppNum, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	rustNum, _ = strconv.Atoi(scanner.Text())

	students := make([]string, cppNum+rustNum, cppNum+rustNum)
	//cppStudents := make([]string, cppNum, cppNum)
	//rustStudents := make([]string, rustNum, rustNum)
	//var rustStudent string

	for i := 0; i < cppNum+rustNum; i++ {
		scanner.Scan()
		students[i] = scanner.Text()
	}

	//fmt.Println(students)

	var flag bool
	for id1, student1 := range students {
		flag = false

	inner1:
		for id2, student2 := range students {
			//fmt.Println(id1, id2, student1, student2)
			if id1 != id2 && student1 == student2 {
				flag = true
				break inner1
			}
		}
		if !flag {
			count++
		}
	}

	fmt.Println(count)
}
