package main

import "fmt"

type Rectangle struct {
	length, width int
}

//1. Реализуем функцию и метод для подсчет периметра прямогуольника
// ВАЖНО: Все параметры передаем как копии

//4. При вызове этого метода неважно, будет ли он вызываться у экземпляра или у его ссылки
func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

//5. Данную функцию можно вызывать ТОЛЬКО у копии прямоугольника (но не у его ссылки)
func Perimeter(r Rectangle) int {
	return 2 * (r.length + r.width)
}

//6. Допустим будет метод, который меняет значение состояния структуры на новое, но этот метод - Value Reciever
func (r Rectangle) SetLength(newLength int) {
	r.length = newLength
}

func main() {
	//2. Создаем экземпляр прямоугольника
	rectangleAsValue := Rectangle{10, 10}
	fmt.Println("Call function for rectangleAsValue:", Perimeter(rectangleAsValue))
	fmt.Println("Call method for rectangleAsValue:", rectangleAsValue.Perimeter())

	//3. Создадим указатель на прямоугольник
	rectangleAsPointer := &rectangleAsValue
	fmt.Println("Call method for rectangleAsPointer:", rectangleAsPointer.Perimeter())
	//Perimeter(rectangleAsPointer) -  Иллюстрация к пункту 5

	//7. Вызываем метод SetLength у экземпляра rectangleAsValue
	fmt.Println("Before call method SetLength:", rectangleAsValue)
	rectangleAsValue.SetLength(9999)
	fmt.Println("Aftet call method SetLength:", rectangleAsValue)

	//8. Вызываем метод SetLength у ссылки на rectangleAsValue
	rectangleAsPointer.SetLength(999999999)
	fmt.Println("After call method SetLength for &rectangle:", *rectangleAsPointer)

}
