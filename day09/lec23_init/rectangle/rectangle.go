package rectangle

import "fmt"

func init() {
	fmt.Println("Init function form rectangle package!")
	fmt.Println("Name:", name, "Age:", age)
}

//1. Инициализируем переменные уровня пакета
var (
	name string = "John"
	age  int    = 33
)

type Rectangle struct {
	A, B int
}

func New(newA, newB int) *Rectangle {
	return &Rectangle{newA, newB}
}
