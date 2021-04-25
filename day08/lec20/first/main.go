package main

import (
	"fmt"
	"unicode/utf8"
)

//1. Описание интерфейса (описание того, что должен уметь претендент)
type BigWord interface {
	IsBig() bool
}

//2. Наш претендент, которого надо научить делать IsBig() bool
type MySuperString string

//3. Реализация IsBig() у претендента MySuperString
func (mss MySuperString) IsBig() bool {
	if utf8.RuneCountInString(string(mss)) > 10 {
		return true
	}
	return false
}

func main() {
	sample := MySuperString("akj")
	var interfaceSample BigWord // Объявление переменной, типа интерфейса BigWord
	interfaceSample = sample    // присваивание значения для переменной тип BigWord возможно,
	//потому что sample (типа MySuperString ) удовлетворяет интерфейсу BigWord
	fmt.Println("IsBig?", interfaceSample.IsBig())

	//4. Попытка присвоить значение с типажом, неудовлетворяющему интерфейсу
	// var interfaceBadSample BigWord
	// interfaceBadSample = "abcdef" // тип string не имеет реализации метода IsBig , поэтому не удовлетворяет интерфейсу

}
