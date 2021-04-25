package main

import "fmt"

//1. Вопрос - имеет ли для интерфейса значение тот факт, что метод, реализованный для претендента
// , в качестве получателя принимает значение или указатель?

//0. В Go принято называть интерфейсы , с окончанием на `er` (Describer, Worker, Pooller,....)
type Describer interface {
	Describe()
}

type Student struct {
	name string
	age  int
}

func (std Student) Describe() {
	fmt.Printf("Student name: %s and age %d\n", std.name, std.age)
}

type Professor struct {
	name string
	age  int
}

func (prof *Professor) Describe() {
	fmt.Printf("Professor name %s and age %d\n", prof.name, prof.age)
}

func main() {
	var descr1 Describer
	stud1 := Student{"Alex", 23}
	descr1 = stud1 //Student реализует интерфейс Describer
	descr1.Describe()

	stud2 := &Student{"Bob", 21} // Поскольку экземпляр -ссылка, разыменовать ее имеет право кто-угодно (в том числе и компилятор)
	descr1 = stud2
	descr1.Describe()

	var descr2 Describer
	prof1 := &Professor{"Viktor Wann", 72}
	descr2 = prof1 // &Professor реализует интерфейс Describer
	descr2.Describe()

	prof2 := Professor{"Bob Brown", 65}
	prof2.Describe() // Здесь ссылку на &prof берет компилятор
	descr2 = prof2   //Professor не реализует интерфейс Describer
	// Дело в том, что сам по себе интерфейс - не референсный тип
	// Выливается это в следующее следствие:
	// Когда мы работаем с обычным методом, то взять референс на экземпляр  может компилятор
	// Но когда мы пытаемся сделать это через интерфейс - в нем в принципе комплятор не видит никаких ссылок!
}
