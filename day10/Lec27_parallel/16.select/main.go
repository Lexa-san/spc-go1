package main

import (
	"fmt"
	"time"
)

//1. На практике select чаще всего используется для того, чтобы предпринимать какие-то действия,
// пока в каналы еще не пришли данные

//Пример
func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func main() {
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}
	}

}
