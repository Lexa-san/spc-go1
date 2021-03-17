package main

import (
	"fmt"
	"sync"
	"time"
)

//1. Еще один инструмент для оркестрирования горутинами - это WaitGroup
// По сути WaitGroup - это счетчик горутин.
// Когда горутина запускается делается WaitGroup++
// Когда горутина завершается делается WaitGroup--
//Таким образом когда WaitGroup == 0 делаем вывод, что все горутины отработали!

//Пример

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done() //WaitGroup--
}

func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1) // WaitGroup++
		go process(i, &wg)
	}
	wg.Wait() // if WaitGroup == 0 ? До тех пор, пока это условие не выполнено - мы заблокированы в данной строке для main горутины
	fmt.Println("All go routines finished executing")
}
