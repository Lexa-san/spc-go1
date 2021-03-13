package main

import (
	"fmt"
	"math"
)

func main() {
	//Простейший вывод на косноль. println - это вывод аргумента + '\n'
	fmt.Println("Hello world", "Hello another")
	fmt.Println("Second line")
	//Функция print - простой вывод аргумента
	fmt.Print("First")
	fmt.Print("Second")
	fmt.Print("Third")
	//Форматированный вывод: Printf - стандартный вывод os.Stdout с флагами форматирования
	fmt.Printf("\nHello, my name is %s\nMy age is %d\n", "Bob", 42)
	///////////////////////////////////
	///////////////////////////////////
	//Декларирование переменных
	var age int
	fmt.Println("My age is:", age)
	age = 32
	fmt.Println("Age after assignment:", age)

	//Декларирование и инициализация пользовательским значением
	var height int = 183
	fmt.Println("My height is:", height)

	//В чем "полустрогость" типзации? Можно опускать тип переменной
	var uid = 12345
	fmt.Println("My uid:", uid)
	//Декларирование и инициализация переменных одного типа (множественный случай)
	var firstVar, secondVar = 20, 30
	fmt.Printf("FirstVar:%d SecondVar:%d\n", firstVar, secondVar)
	//Декларирвоание блока переменных
	var (
		personName string = "Bob"
		personAge         = 42
		personUID  int
	)

	fmt.Printf("Name: %s\nAge %d\nUID: %d\n", personName, personAge, personUID)

	//Немного странного
	var a, b = 30, "Vova"
	fmt.Println(a, b)
	a = 300
	//Немного хорошего. Повторное декларирование переменной приводит к ошибке компиляции
	//var a = 200

	//Короткая декларция (корткое объявление)
	count := 10
	fmt.Println("Count:", count)
	count = count + 1
	fmt.Println("Count:", count)
	//Множественное присваивание через :=
	aArg, bArg := 10, 30
	fmt.Println(aArg, bArg)
	aArg, bArg = 30, 40
	fmt.Println(aArg, bArg)
	// aArg, bArg := 10, 30
	// fmt.Println(aArg, bArg)

	//Исключение из этого правила
	bArg, cArg := 300, 400
	fmt.Println(aArg, bArg, cArg)

	//Пример
	width, length := 20.5, 30.2
	fmt.Printf("Min dimensional of rectangle is : %.2f\n", math.Min(width, length))

}
