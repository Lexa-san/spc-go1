package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "Hello world"
	fmt.Println(name)

	//1. Строка - это байтовый слайс со своими особенносятми при обращении
	//к нижелажащему массиву
	word := "Тестовая строка"
	fmt.Printf("String %s\n", word)
	// Какие значения байт сейчас находятся в слайсе word?
	fmt.Printf("Bytes: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%x ", word[i]) //%x - формат представления 16-ти ричного байта
	}
	fmt.Println()
	// Каким образом получать доступ к отдельно стоящим символам?
	fmt.Printf("Characters: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%c ", word[i]) //%c - формат представления символа
	}
	fmt.Println()
	//2. Строки в Go - хранятся как наборы UTF-8символов. Каждый символ, вообще говоря, может занимать
	// больше чем 1 байт

	//3. Руна (Rune) - это стандартный встроенный тип в Go (alias над int32), позволяющий хранить
	//единый неделимый UTF символ ВНЕ ЗАВИСИМОСТИ ОТ ТОГО сколько байт он занимает
	fmt.Printf("Runes: ")
	runeSlice := []rune(word) // Преобразование слайса байт к слайсу рун []byte(sliceRune)
	for i := 0; i < len(runeSlice); i++ {
		fmt.Printf("%c ", runeSlice[i])
	}
	fmt.Println()
	//4. Итерирование по строке с использованием рун
	for id, runeVal := range word { // id - это индекс байта, с КОТОРОГО НАЧИНАЕТСЯ РУНА runeVal
		fmt.Printf("%c starts at postion %d\n", runeVal, id)
	}

	//5. Создание строки из слайса байт
	myByteSlice := []byte{0x40, 0x41, 0x42, 0x43} // Исходное представление байтов
	myStr := string(myByteSlice)
	fmt.Println(myStr)

	myDecimalByteSlice := []byte{100, 101, 102, 103} // Синтаксический сахар - можно использовать десятичное представление байтов
	myDecimalStr := string(myDecimalByteSlice)
	fmt.Println(myDecimalStr)

	//6. Создание строки из слайса рун
	// Руны как hex
	runeHexSlice := []rune{0x45, 0x46, 0x47, 0x48}
	myStrFromRune := string(runeHexSlice)
	fmt.Println("From Runes(hex):", myStrFromRune)
	// Руны как литералы
	runeLiteralSlice := []rune{'V', 'a', 's', 'y', 'a'} // '' - таким образом обозначается руна
	myStrFromRuneLiterals := string(runeLiteralSlice)
	fmt.Println("From Runes(literals):", myStrFromRuneLiterals)

	fmt.Printf("%s and %T\n", myStrFromRuneLiterals, myStrFromRuneLiterals)

	//7. Длина и емкость строки
	// Длина len() - количество байт в слайсе
	fmt.Println("Length of Вася:", len("Вася"), "bytes")
	// Длина RuneCounter - количество !рун!
	fmt.Println("Length of Вася:", utf8.RuneCountInString("Вася"), "runes")
	// Вычисление емкости строки - бессмысленно, т.к. строка базовый тип
	fmt.Println(cap([]rune("Вася")))

	//8. Сравнение строк == и !=. Начиная с go 1.6
	word1, word2 := "Вася", "Петя"
	if word1 == word2 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}

	//9. Конкатенация
	word3 := word1 + word2
	fmt.Println(word3)

	//10. Строитель строк (java -> StringBuiler)
	firstName := "Alex"
	secondName := "Johnson"
	fullName := fmt.Sprintf("%s #### %s", firstName, secondName)
	fmt.Println("FullName:", fullName)

	//11. Строки не изменяемые
	// fullName[0] = "Q"

	//12. А слайсы изменяемые :)
	fullNameSlice := []rune(fullName)
	fullNameSlice[0] = 'F'
	fullName = string(fullNameSlice)
	fmt.Println("String mutation:", fullName)

	//13. Сравнение рун
	if 'Q' == 'Q' {
		fmt.Println("Runes equal")
	} else {
		fmt.Println("Runes not equal")
	}

	//14. Где живут полезные методы работы со строками?
	// import "strings"

}
