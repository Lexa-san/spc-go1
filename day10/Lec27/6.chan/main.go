package main

import (
	"fmt"
	"time"
)

//1. Немного перепишем последнюю программу, чтобы лучше увидеть как устроен процесс блокирования

func newGoRoutine(done chan bool) {
	fmt.Println("Hey, I'm newGoRoutine and I'm going sleep!")
	time.Sleep(4 * time.Second)
	fmt.Println("newGoRoutine awake and going to send data to channel")
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Println("I'm main goroutine and I wanna call newGoRoutine")
	go newGoRoutine(done)
	<-done
	fmt.Println("Ok, Main goroutine recieved data and gonna die!")
}
