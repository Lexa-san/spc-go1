package main

import (
	"fmt"
	"time"
)

//1. Select - это инструмент, позволяющий выбирать из множества канальных операций (чтение/запись) для множества каналов.
// Если из 10 каналов что-то пришло в один - select выбирает его
// Если из 10 каналов что-то пришло сразу в два и более - select выбирает случайный

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"

}
func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2: // выбирается этот кейс, т.к. в этот канал будут помещены данные быстрее
		fmt.Println(s2)
	}
}
