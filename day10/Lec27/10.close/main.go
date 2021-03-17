package main

import "fmt"

//1. Закрытие каналов и итерирование
// Со стороны получателя можно использоваться синтаксис
// val, ok := <- ch
//где val - значение помещенное в канал отправителем
// ok - true/false в зависимости от того, открыт канал или уже закрыт отправителем.
// если канал закрыт то в val будет помещено zeroValue для типа канала

func generator(ch chan int) {
	for i := 0; i < 25; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go generator(ch)
	for {
		val, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Recieved from channel", val)
	}

	// Конструкцию можно упростить и использовать
	// for val := range ch {
	// 	fmt.Println("Recieved from channel:", val)
	// }
}
