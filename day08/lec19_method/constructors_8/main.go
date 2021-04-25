package main

import "fmt"

//1. Создадим структуру Rectangle
type Rectangle struct {
	length, width int
}

//2. Создадим конструктор для Rectangle
// Для имен конструкторов в Go договорились использовать функцию с следующим названием:
// * New() если данный конструткор на файл один (в файле присутствует описание только одной структуры)
// * New<StructName>() - если в файле присутсвуют еще какие-то структуры

//3. В Go принято возвращать из конструктора не сам экземпляр, а ссылку на него
func NewRectangle(newLength, newWidth int) *Rectangle {
	return &Rectangle{newLength, newWidth}
}

//4. Добавим 2 метода
func (r *Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}

func main() {
	r := NewRectangle(10, 20)
	fmt.Printf("Type as %T and value %v\n", r, r)
	fmt.Println("Perimeter:", r.Perimeter())
	fmt.Println("Area:", r.Area())

}

type Circle struct {
	radius float64
}

func NewCircle(newRadius float64) *Circle {
	return &Circle{newRadius}
}
