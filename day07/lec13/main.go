package main

import "fmt"

//1. Явные функции (в принципе любая функция в Go) - является
//экземлпяром 1-го уровня (функцию можно присваивать в перменную, ее можно передавать в
//качестве параметра и возвращать из других функций)

//2. Возврат функции в качестве значения
func calcAndReturnValidFunc(command string) func(a, b int) int {
	if command == "addition" {
		return func(a, b int) int { return a + b }
	} else if command == "substraction" {
		return func(a, b int) int { return a - b }
	} else {
		return func(a, b int) int { return a * b }
	}
}

//3. Функция как парметр в другой функции
func recieveFuncAndReturnValue(f func(a, b int) int) int {
	var intVarA, intVarB int
	intVarA = 100
	intVarB = 200

	return f(intVarA, intVarB)
}

func add(a, b int) int {
	return a + b
}
func main() {

	var command string
	command = "substraction"
	res := calcAndReturnValidFunc(command)
	fmt.Println("Result with command :", command, "value:", res(10, 20))
	fmt.Println(res(30, 40))

	//4. Тип функции
	fmt.Printf("Type of func is %T\n", res)
	fmt.Printf("Type of calcAndReturnValidFunc is %T\n", calcAndReturnValidFunc)
	//5. Тип функции в Go определяется как входными парамтерами, так и выходными

	fmt.Println("recieveFuncAndReturnValue(add):", recieveFuncAndReturnValue(add))
	fmt.Println(recieveFuncAndReturnValue(func(a, b int) int {
		return a*a + b*b + 2*a*b
	}))

}
