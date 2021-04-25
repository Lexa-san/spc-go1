package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	// Boolean => default false
	var fisrtBoolean bool
	fmt.Println(fisrtBoolean)

	// Boolean operands
	aBool, bBool := true, false
	fmt.Println("AND:", aBool && bBool)
	fmt.Println("OR:", aBool || bBool)
	fmt.Println("NOT:", !aBool)
	// not acceptable
	// fmt.Println("+:", aBool+bBool)

	// Numeric, Integers
	// int8, int16, int32, int64, int
	// uint8, uint16, uint32, uint64, uint
	var a int = 32
	b := 92
	fmt.Println("a:", a, ", b:", b, ", Summ of a+b:", a+b)
	// Вывод тимпа через %T
	fmt.Printf("Type is %T\n", a)
	// Узнаем, сколько байт занимает переменная типа int
	fmt.Printf("Type %T size of %d bytes\n", a, unsafe.Sizeof(a))

	// Эксперимент 1. При использовании короткого объявления - тип для целого числа - int платформо-зависимоый
	fmt.Printf("Type %T size of %d bytes\n", a, unsafe.Sizeof(b))

	// Эксперимент 2. Используйте явное приведение типов
	var first32 int32 = 12
	var second64 int64 = 13
	fmt.Println(int64(first32) + second64)

	// Эксперимент 3. int != int32 and int != int32 - используйте явное приведение типов
	var third64 int64 = 16123414
	var forthInt int = 156234
	fmt.Println(third64 + int64(forthInt))
	// Аналогично поступать с unit-типами

	// Numeric. Float
	// float32, float64
	floatFirst, floatSecond := 1.23, 4.56
	fmt.Printf("type of a %T and type of %T\n", floatFirst, floatSecond)
	sum := floatFirst + floatSecond
	sub := floatFirst - floatSecond
	fmt.Printf("Sum: %.3f and Sub: %.3f\n", sum, sub)

	// Numeric. Complex
	c1 := complex(5, 7)
	c2 := 12 + 34i
	fmt.Println(c1 + c2)
	fmt.Println(c1 * c2)

	// Strings. Строка - набор БАЙТ
	name := "Вася"
	lastname := "Pupkin"
	concat := name + " " + lastname
	fmt.Println("Full name:", concat)
	// функция len() возвращает кол-во элементов в наборе
	fmt.Println("Length of string:", name, len(name))
	fmt.Println("Length of string:", lastname, len(lastname))
	// Rune - руна, один utf-символ
	fmt.Println("Length of string:", name, utf8.RuneCountInString(name))
	fmt.Println("Length of string:", lastname, utf8.RuneCountInString(lastname))

	// Поиск подстроки в строке
	totalString, subString := "ABCDEFG", "asd"
	fmt.Println(strings.Contains(totalString, subString))
	// rune -> alias int32
	var sampleRune rune
	// для инициализации руны символьным значенем используйте одинарные кавычки
	var anotherRune rune = 'Q'
	var thirdRune rune = 234
	fmt.Println(sampleRune)
	fmt.Printf("Rune as char %c\n", sampleRune)
	fmt.Printf("Rune as char %c\n", anotherRune)
	fmt.Printf("Rune as char %c\n", thirdRune)

	// "A" < "abcde"
	// -1: first < second
	// 0: first == second
	// 1: first > second
	fmt.Println(strings.Compare("abcde", "a"))

	// byte is alias uint8
	var myByte byte
	fmt.Printf("Type of byte is %T\n", myByte)
}
