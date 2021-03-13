package main

import "fmt"

//1. Указатели - переменная, хранящая в качестве значения - адрес в памяти другой переменной

func main() {

	//2. Определение указателя на что-то
	var variable int = 30
	var pointer *int = &variable //&.... - операция взятия адреса в памяти
	//Выше у нас создан указатель на переменную variable
	//В pointer лежит 18293xcd000132 - это место в памяти, где хранится int значение 30
	fmt.Printf("Type of pointer %T\n", pointer)
	fmt.Printf("Value of pointer %v\n", pointer)

	//3. А какое нулевое значение для указатели?
	var zeroPointer *int //zeroValue имеет значение nil (это указатель в никуда)
	fmt.Printf("Type %T and value %v\n", zeroPointer, zeroPointer)
	if zeroPointer == nil {
		zeroPointer = &variable
		fmt.Printf("After initializatoin type %T and value %v\n", zeroPointer, zeroPointer)
	}

	//4. Разыменование указателя (получение значения): *pointer - возвращает значение, хранимое по адресу
	var numericValue int = 32
	pointerToNumeric := &numericValue

	fmt.Printf("Value in numericValue is %v\nAddress is %v\n", *pointerToNumeric, pointerToNumeric)

	//5. Создание указателей на нулевые занчения типов
	// var zeroVar int
	// var zeroPoint *int = &zeroVar
	zeroPoint := new(int) // Создает под капотом zeroValue для int, и возвращает адрес, где этот 0 хранится
	fmt.Printf("Value in *zeroPointer %v\nAddress is %v\n", *zeroPoint, zeroPoint)

	//6. Изменение значения хранимого по адресу через указатель
	zeroPointerToInt := new(int)
	fmt.Printf("Addres is %v and Value in zeroPointerToInt is %v\n", zeroPointerToInt, *zeroPointerToInt)
	*zeroPointerToInt += 40
	fmt.Printf("Addres is %v and New Value in zeroPointerToInt is %v\n", zeroPointerToInt, *zeroPointerToInt)

	b := 345
	a := &b
	c := &b
	*a++
	*c += 100
	fmt.Println(b)
	//7. Указательная арфиметика ОТСУТСТВУЕТ ПОЛНОСТЬЮ
	// У вас на руках адрес одной ячейки - вы можете через этот адрес продвинуться в другие ячейки
	//

	//8. Передача указателей в функции
	// Колоссальный прирост производительности засчет того, что передается не значение (которые должно копироваться)
	// а передается лишь адрес в памяти, за которым уже хранится какое-то значение
	sample := 1
	//samplePointer := &sample

	fmt.Println("Origin value of sample:", sample)
	changeParam(&sample)
	fmt.Println("After changing sample is:", sample)

	//9. Возврат поинтера из функции (В С++ результат работы такого механизма - неопределен)
	ptr1 := returnPointer()
	ptr2 := returnPointer()
	fmt.Printf("Ptr1: %T and address %v and value %v\n", ptr1, ptr1, *ptr1)
	fmt.Printf("Ptr2: %T and address %v and value %v\n", ptr2, ptr2, *ptr2)

}

//9.1  Инициализация функции, возвращающей указатель
func returnPointer() *int {
	var numeric int = 321
	return &numeric //В момент возврата Go перемещает данную переменную в кучу
}

//8.1 Определдение фукнции, принимающей параметр как указатель
func changeParam(val *int) {
	*val += 100
}
