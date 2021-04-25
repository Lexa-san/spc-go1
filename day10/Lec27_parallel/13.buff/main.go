package main

import (
	"fmt"
	"time"
)

//1. Как блокируется буферезированный канал?
func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

//2. Как только буфер заполняется - канал блокируется до тех пор, пока не будет освобождено место (буфер может быть освобожден не до конца!)

func main() {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)

	}
}

// 3. Длина и вместимость канала.
// Длина кала len(ch) - сколько сейчас элементов в канале
// Вместимость cap(ch) - какой размер буфера у канала
