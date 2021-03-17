package main

import "fmt"

//1. Создадим чуть более полезную программу, которая будет делать следующие действия:
// берем число, например 456
// Подсчитаем сумму квадартов цифр и сумму кубов, а затем сложим полученные результаты
//Делать будем следующим образом:
// * main gorutine получает число и вызывает 2 другие горутины, по итогу, получив от них результаты,
// она просто их сложит и выведет на консоль
// ** squaresGoRoutine - будет запущена main и подсчитает сумму квадратов всех цифр, результат положит в канал
// ** cubesGoRoutine - будет запущена main и подсчиатет сумму кубов всех цифр , результат полужит в канал

func squaresGoRoutine(num int, squareChan chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit
		num /= 10
	}
	squareChan <- sum
}

func cubesGoRoutine(num int, cubeChan chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit * digit
		num /= 10
	}
	cubeChan <- sum
}

func main() {
	number := 456
	squareChan := make(chan int)
	cubeChan := make(chan int)
	go squaresGoRoutine(number, squareChan)
	go cubesGoRoutine(number, cubeChan)
	squaresSum, cubesSum := <-squareChan, <-cubeChan
	fmt.Println("Total result is:", squaresSum+cubesSum)
}
