// В Go существует 2 механизма сигнализирования анаомального поведения
// 1-ый механизм это ошибки Error (ЯВЛЯЕТСЯ КАНОНИЧНЫМ ИСПОЛНЕНИЕМ НА СЛУЧАЙ НЕНОРМАЛЬНОГО ПОВЕДЕНИЯ)
// 2-ой механизм - это паника (не лучший вариант, так как приводит сразу к аварийному завершению, и в принципе
// мог быть обычной ошибкой)

package main

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strconv"
)

func funcWuthError(a int) (string, error) {
	if a%2 == 0 {
		return "Even", nil
	}
	return "", errors.New("THIS IS ERRRO!")
}

func PanicRecover() {
	if r := recover(); r != nil {
		fmt.Println("Panic satisfied:", r)
		debug.PrintStack()
	}
}

func panicExample(firstName *string, lastName *string) {
	defer PanicRecover() // даже в случае возникновения паники - первым дело будут вызваны все deferred функци.
	if firstName == nil {
		panic("runtime errror: firstname can not be nil!")

	}

	if lastName == nil {
		panic("runtime error: lastname can not be nil!")
	}

	fmt.Println("Full name:", *firstName, *lastName)
}

func main() {
	numLiteral := "12"
	num, err := strconv.Atoi(numLiteral)
	if err != nil {
		fmt.Println("can not convert this value to int:", err)
		return
	}
	fmt.Println("Convertion done:", num)

	var name = "Bob"
	panicExample(&name, nil)

	ans, err := funcWuthError(5)
	if err != nil {
		fmt.Println("not use odd values", err)
		return
	}
	fmt.Println(ans)
}
