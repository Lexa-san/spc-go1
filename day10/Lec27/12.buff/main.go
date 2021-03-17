package main

import "fmt"

//1. Буферезированный канал - канал с буфером, в который можно напихать со стороны отправителя не 1 пак данных, а столько, сколько
// позволяет в буфер.

// ch := make(chan int, capacityIntValue)

func main() {
	ch := make(chan string, 2) // Создадим канал вместимостью 2
	ch <- "Bob"
	ch <- "Alex"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
