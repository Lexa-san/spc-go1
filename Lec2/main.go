package main

import (
	"fmt"
	"math"
)

func main() {
	// Простейший вывод на консоль
	fmt.Println("Hello world")
	fmt.Println("Second line")

	// Простейший вывод без переноса строки
	fmt.Print("First")
	fmt.Print("Second")
	fmt.Print("Third")

	// форматированный вывод& Printf - стандартный вывод os.Stdout с флагами форматирования
	fmt.Printf("\nHello, my name is %s\nMy age is %d\n", "Bob", 42)

	//////////////////////////////
	// Декларирование переменных

	var age int
	fmt.Println("My age is:", age)
	age = 123
	fmt.Println("var `age` after assignment:", age)

	// Декларирование и инициализация пользовательским значением
	var height int = 234
	fmt.Println("My height is:", height)

	// В чем полустрогость типизации
	uid := 12345
	fmt.Printf("My uid:%d. Type is %T\n", uid, uid)
	// Декларирование и инициализация переменный одного типа (множественный случай)
	firstVar, secondVar := 1, 2
	fmt.Printf("firstVar:%d, secondVar:%d\n", firstVar, secondVar)
	// Декларирование блока переменных - наш выбор
	var (
		personName string = "Bob"
		personAge  int    = 25
		personUid  int
	)

	fmt.Printf("Name: %s, Age: %d, Uid: %d\n", personName, personAge, personUid)

	// Немного странного
	a, b := 123, "qwe"
	fmt.Println(a, b)

	// Немного хорошего. Повторное декларирование переменной приводит к ошибки компиляции
	// vat a = 200

	// Короткая декларация
	count := 10
	fmt.Println("Count:", count)
	count++
	fmt.Println("Count:", count)
	// Множественное присваивание через :=
	aArg, bArg := 10, "Vova"
	fmt.Println(aArg, bArg)
	aArg, bArg = 20, "Pasha"
	fmt.Println(aArg, bArg)
	// исключение из правила - достаточно одной новой переменной
	aArg, cArg := 30, 123123
	fmt.Println(aArg, bArg, cArg)

	// Пример
	width, height := 20.5, 30.2
	fmt.Printf("Min dimentionl of rectangle is: %.2f\n", math.Min(width, height))
}
