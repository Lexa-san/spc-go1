package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var name string
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() { // Команда захвата потока ввода и сохранения в буфер (захват идет до символа окончания строки)
		name = input.Text() // Команда возвращения элементов, помещенных в буфер (отдаст string)
	}
	fmt.Println(name)

	fmt.Println("For loop started:")
	for {
		if input.Scan() {
			result := input.Text()
			if result == "" {
				break
			}
			fmt.Println("Input string is:", result)
		}
	}

	//Преоброазование строкового литерала к чему-нибудь числовому
	numStr := "10"
	numInt, _ := strconv.Atoi(numStr) // Atoi - Anything to Int (именно int)
	fmt.Printf("%d is %T\n", numInt, numInt)

	numInt64, _ := strconv.ParseInt(numStr, 10, 64)
	numFloat32, _ := strconv.ParseFloat(numStr, 32) // Но это 64-разрядное число будет без проблем ГАРАНТИРОВАНО ПРЕВОДИТЬСЯ К 32
	fmt.Println(numInt64, numFloat32)
	fmt.Printf("%.3f and %T\n", numFloat32, float32(numFloat32))

}
