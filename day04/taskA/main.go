package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var (
		seriesNum   int
		interestNum int
		i           int
	)
	scanner := bufio.NewScanner(os.Stdin)

	// or fmt.Scanf("%q", &name)
	fmt.Scan(&seriesNum, &interestNum)

	var seriesArr [100]string
	i = 0
	for scanner.Scan() {
		seriesArr[i] = scanner.Text()

		i++
		if i == seriesNum {
			break
		}
	}

	var seriesName string
	var found bool

	i = 0
	for scanner.Scan() {
		seriesName = scanner.Text()

		found = false
	inner:
		for _, val := range seriesArr {
			//if val == seriesName {
			if strings.Compare(val, seriesName) == 0 {
				fmt.Println("ЕСТЬ")
				found = true
				break inner
			}
		}

		if !found {
			fmt.Println("НЕТ В НАЛИЧИИ")
		}

		i++
		if i == interestNum {
			break
		}
	}

}
