package main

import "fmt"

func main() {
	// Слайсы (они же срезы)
	// Слайс - это динамическая обвязка над массивом.
	startArr := [4]int{10, 20, 30, 40}
	// Слайс инициализируется пустыми скобками.
	// Создали слайс на основе существующего массива.
	var startSlice []int = startArr[0:2]
	fmt.Println("Slice[0:2]:", startSlice)

	//Создание слала без явно	 инициализации массива
	secondSlice := []int{15, 25, 35, 45}
	fmt.Println("Second slice:", secondSlice)

	//Изменение элемнтв среза
	originArr := [...]int{15, 25, 35, 45, 55, 65, 75}
	firstSlice := originArr[1:4] // это набор ссылок на элементы нижележащего массива
	for i, _ := range firstSlice {
		firstSlice[i]++
	}
	fmt.Println("OriginArr:", originArr)
	fmt.Println("FirstSlice:", firstSlice)

	// один массив и два производных среза
	fSlice := originArr[:]
	sSlice := originArr[2:5]
	fmt.Println("Before modification: Arr:", originArr, "fSlice:", fSlice, "sSlice", sSlice)
	fSlice[3]++
	sSlice[1]++
	fmt.Println("Before modification: Arr:", originArr, "fSlice:", fSlice, "sSlice", sSlice)

	// Срез как встроенный тип
	//type slice struct {
	//	Length int
	//	??? int
	//	ZeroElement *byte
	//}

	// Длина и емкость массива
	wordSlice := []string{"one", "two", "three"}
	fmt.Println("slice:", wordSlice, "Length:", len(wordSlice), "Capacity:", cap(wordSlice))
	wordSlice = append(wordSlice, "four")
	fmt.Println("slice:", wordSlice, "Length:", len(wordSlice), "Capacity:", cap(wordSlice))
	wordSlice = append(wordSlice, "five")
	fmt.Println("slice:", wordSlice, "Length:", len(wordSlice), "Capacity:", cap(wordSlice))
	// Capacity (емкость) слайла - это целочисленного значение, показывающее сколько элементов можно добавить в срез
	// без выделения дополнительной памяти под нижележащий массив.

	// Допустим у нас есть срез на 3 элемента.
	// Компилятор по создании этог осреза сначала создал массив ровно на 3 элемента.
	// После этого компилятор вернул адрес, где этот массив живет.
	// Срез запомнил этот адрес и стал ссылаться на него.
	// Потом начинаем деформировать слайл (добавляем элементы)
	// Проблема - в нижележащем массиве все 3 ячейки заняты. Что делать?
	// Комплиятор ищет в памяти место для массива размера 3*2 (n*2, где n - первоначальный размер массива)
	// Как место найдено, в него компируется старые элементы на свои позици.
	// На 4-ю позицию добавляется новый элемент.
	// После компилятор возвращаяет нашего слайсу новый адрес в памяти, где находтся массив под 6 элементов.

	numerics := []int{1, 2}
	for i := 0; i < 200; i++ {
		if i%10 == 0 {
			fmt.Printf("Current len: %d Current cap: %d\n", len(numerics), cap(numerics))
		}
		numerics = append(numerics, i)
	}

	// Важно: после выделения памяти под новый массив, ссылки со старым будут перенесены в новый.
	numArr := [2]int{1, 2}
	numSlice := numArr[:]

	numSlice = append(numSlice, 3) //в этот момен numSlice больше не ссылается на numArr
	numSlice[0] = 10
	fmt.Println(numArr)
	fmt.Println(numSlice)

	numArr2 := [4]int{1, 2}
	numSlice2 := numArr2[:]

	numSlice2 = append(numSlice2, 3) //в этот момен numSlice больше не ссылается на numArr
	numSlice2[0] = 10
	fmt.Println(numArr2)
	fmt.Println(numSlice2)

	// Как создавать слайсы наиболее эффективно?
	// make() -  это функция, позволяющая более детально создавать срезы
	s1 := make([]int, 10, 15)
	// []int - тип коллекции
	// 10 - длина
	// 15 - емкость
	// Сперва инициализируем app = [15]int
	// Затем по нему делается срез arr[0:10]
	// После чего он возвращается
	fmt.Println(s1)
	fmt.Printf("s1 len: %d s1 cap: %d\n", len(s1), cap(s1))

	// Добавление элементо в срез
	myWords := []string{"one", "two", "three"}
	fmt.Println("myWords:", myWords)
	//myWords = append(myWords, "four")
	//fmt.Println("myWords:", myWords)
	anotherSlice := []string{"four", "five", "six"}
	myWords = append(myWords, anotherSlice...)
	fmt.Println("myWords:", myWords)
	myWords = append(myWords, "seven", "eight")
	fmt.Println("myWords:", myWords)

	// Многомерный срез
	mSlice := [][]int{
		{1, 2},
		{3, 4, 5, 6},
		{10, 20, 30},
		{},
	}
	fmt.Println("mSlice:", mSlice)

}
