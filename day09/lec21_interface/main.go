package main

import "fmt"

type Describer interface {
	Describe()
}

type Student struct {
	name string
	age  int
}

func (std Student) Describe() {
	fmt.Printf("%s and %d y.o\n", std.name, std.age)
}

func TypeFinder(i interface{}) {
	switch v := i.(type) { //Присовим переменной v значение лежащие под предполагаемым интерфейсом
	case string:
		fmt.Println("This is string")
	case int:
		fmt.Println("This is int")
	case Describer: // Вывод - с типом интерфейса можно сравниваться точно так же, как и с любым другим типажом языка
		// это как раз и говорит о том, что интерфейсы - полу-абстрактные типы.
		fmt.Println("I'm describer!")
		v.Describe()
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	std := Student{"Alex", 23}
	TypeFinder(10)
	TypeFinder("Hello")
	TypeFinder(nil)
	TypeFinder(std)
}
