package main

//1. Select как инструмент защиты от deadlock
// Добавление default страхует от появление deadlock в ходе выполнения и берет работу на себя (попробуйте убрать default и все умрет)

import "fmt"

func main() {
	var ch chan string
	select {
	case v := <-ch:
		fmt.Println("received value", v)
	default:
		fmt.Println("default case executed")

	}
}
