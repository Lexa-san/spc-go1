package main

import (
	"fmt"
	"math"
)

//1. Константы - это неизменяемые переменные, которые служат для:
//	1) Более строго опнимания кода
//	2) Для того, чтобы случайно не поменять значение (предполагается что значение констатнты не изменно)
//	3) Для удобных преобразований

const (
	MAIN_PORT = "8001"
)

func main() {
	//2. Объвление одной константы
	const a = 10
	fmt.Println(a)
	//3. Объявление блока констант с областью видимости внутри функции main
	const (
		ipAddress string = "127.127.00.03"
		port             = "8000"
		dbName           = "postgres"
	)
	fmt.Println("ipAddress value:", ipAddress)
	fmt.Println(checkPortIsValid())

	//4. Константу никак нельзя поменять в ходе работы программы
	// const b = 200
	// b = 30

	//5. Значения констант ДОЛЖНЫ БЫТЬ ИЗВЕСТНЫ на момент компиляции
	var sqrt = math.Sqrt(25)
	//const sqrt = math.Sqrt(25) //Нельзя присвоить в константу что-либо, что является результатом вызова функции, метода
	fmt.Println("Var sqrt:", sqrt)

	//6. Типизированные и нетипизированные константы
	const ADMIN_EMAIL string = "admin@admin.com" // Указание типа - повышение читабельности кода

	//7. Нетипизирвоанные константы и их профит
	//При использовании нетипизированных констант мы разрешаем компилятору
	//исопльзовать неявное приведение типов в момент присваивания значеий констант в перменные
	const NUMERIC = 10
	var numInt8 int8 = NUMERIC
	var numInt32 int32 = NUMERIC
	var numInt64 int64 = NUMERIC
	var numFloat32 float32 = NUMERIC
	var numComplex complex64 = NUMERIC

	fmt.Printf("numInt8 value %v type %T\n", numInt8, numInt8)
	fmt.Printf("%v + %v is %v\n", numInt8, NUMERIC, numInt8+NUMERIC)
	fmt.Printf("numInt32 value %v type %T\n", numInt32, numInt32)
	fmt.Printf("numInt64 value %v type %T\n", numInt64, numInt64)
	fmt.Printf("numFloat32 value %v type %T\n", numFloat32, numFloat32)
	fmt.Printf("numComplex value %v type %T\n", numComplex, numComplex)
	//8. Константы в Go зашиваются в момент компиляции в RUNTIME программы и не выбрасываются до ее окончания

}

func checkPortIsValid() bool {
	if MAIN_PORT == "8001" {
		return true
	}
	return false
}
