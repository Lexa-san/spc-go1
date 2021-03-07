package main

import (
	"fmt"
)

func main() {
	var (
		count  int32
		x      int32
		result int32
	)

	fmt.Scan(&count)
	for i := 0; int32(i) < count; i++ {
		fmt.Scan(&x)
		if i%2 == 0 {
			result += x
			continue
		}
		result -= x
	}
	fmt.Println(result)
}
