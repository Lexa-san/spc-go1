package main

import "fmt"

func main() {

	var weight int
	fmt.Scan(&weight)

	if weight != 2 && weight%2 == 0 {
		fmt.Println("YES")
		return
	}

	fmt.Println("NO")

}
