package main

import "fmt"

type Rectangle struct {
	length, width int
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}

func Area(r *Rectangle) int {
	return r.length * r.width
}

func (r *Rectangle) SetWidth(newWidth int) {
	r.width = newWidth
}

func main() {
	//Значение исходное
	rectangleAsValue := Rectangle{10, 10}
	//Ссылка на исходное значение
	rectangleAsPointer := &rectangleAsValue
	fmt.Println("Call Area function for &rectangle:", Area(rectangleAsPointer))
	fmt.Println("Call Area method for &rectangle:", rectangleAsPointer.Area())

	//1. Вызываем метод у value - исходного значения
	fmt.Println("Call Area method for rectangle:", rectangleAsValue.Area())
	//2. Вызываем функцию с параметром value - исходное значение
	//Area(rectangleAsValue)

	//3. Распечатаем исходный прямоугольник
	fmt.Println("Before changing width:", rectangleAsValue)

	//4. Вызываю метод SetWidth у &rectangle
	rectangleAsPointer.SetWidth(999)
	fmt.Println("After change via method SetWidth for &rectangle:", rectangleAsValue)

	//5. Вызов метода SetWidth у rectangle
	rectangleAsValue.SetWidth(888)
	fmt.Println("After change via method SetWidth for rectangle:", rectangleAsValue)
}
