package main

import (
	"fmt"
)

func main() {
	//	Массивы. Основы
	//	1. Определение массива

	//важно передать информацию о кол-ве элементов в массиве
	var arr [5]int
	fmt.Println("This is my array: ", arr)
	arr[0] = 10
	arr[1] = 20
	arr[3] = 30
	arr[4] = 40
	fmt.Println("After elements init: ", arr)

	newArr := [5]int{10, 20, 30}
	fmt.Println("Short declaration and init:", newArr)

	arrEmpty := [...]int{}
	fmt.Println("Empty array:", arrEmpty, "Length:", len(arrEmpty))

	arrWithValues := [...]int{10, 20, 30, 40}
	fmt.Println("Array declaration with:", arrWithValues, "Length:", len(arrWithValues))
	arrWithValues[0] = 10000
	fmt.Println("Arr declaration with [...]", arrWithValues, "Length:", len(arrWithValues))

	//	Массив - это набор значений. При всех манипуляциях массив копируется (на уровне компилятора).
	first := [...]int{1, 2, 3}
	second := first
	second[0] = 10000
	fmt.Println("First array:", first)
	fmt.Println("Second array:", second)

	// Массив и его размер - две составляющие одного типа. Размер - часть типа
	//var aArr [5]int
	//var bArr [6]int
	//aArr[0] = 100
	//bArr = aArr

	// Итерирование по массиву
	floatArr := [...]float64{12.3, 23.4, 34.5, 45, 6}
	for i := 0; i < len(floatArr); i++ {
		fmt.Printf("%d element of array is %.2f\n", i, floatArr[i])
	}

	// Итерирование по масиву через оператор range
	var sum float64
	for id, val := range floatArr {
		fmt.Printf("%d element of array is %.2f\n", id, val)
		sum += val
	}
	fmt.Printf("Total sum is: %.2f\n", sum)

	// Игнорирование id в range based for цикле
	for _, val := range floatArr {
		fmt.Printf("%.2f value WO id\n", val)
	}

	// Многомерные массивы
	words := [2][2]string{
		{"Bob", "Alice"},
		{"Victor", "Jo"},
	}
	fmt.Println("Multidimentional array:", words)
	for _, val1 := range words {
		for _, val2 := range val1 {
			fmt.Printf("%s ", val2)
		}
		fmt.Println()
	}

	// Работа со срезом
	slice := []int{10, 20, 30}
	slice[0] = slice[0] * 10
	slice[1] = 200
	slice = append(slice, 200) // добавление нового элемента
	for i, v := range slice {
		fmt.Println(i, v)
	}

	emptySlice := []int{}
	fmt.Println("Empty slice:", emptySlice)

	emptySlice = append(emptySlice, 10)
	fmt.Println("Not emty slice:", emptySlice)

}
