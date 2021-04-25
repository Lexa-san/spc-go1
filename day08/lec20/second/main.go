package main

import "fmt"

//1. Объявляем интерфейс, декларирующий поведенческий -паттерн (набор методов под реализацию)
type Worker interface {
	Work()
}

//2. Объявляем структуру. Данная структура - претендент на удовлетворение интерфейса Worker
type Employee struct {
	name string
	age  int
}

//3. Реализуем метод Work(), чтобы структура Employee удовлетворяла интерфейсу Worker
func (e Employee) Work() {
	fmt.Println("Now Employee with name", e.name, "is working!")
}

//4. А давайте создадим функцию, которая будет принимать аргументы типа Worker и что-то с ними делать?
func Describer(w Worker) {
	fmt.Printf("Interface with type %T and value %v\n", w, w)
}

//6. Объявляем структура. Данная структура - второй пренедент на удовлетворение интерфейса Worker
type Student struct {
	name         string
	courseNumber int
}

func (s Student) WantToEat() {
	fmt.Println("student with name", s.name, "want to eat!")
}

func (s Student) Work() {
	fmt.Println("Student with name", s.name, "is working!")
}

func main() {
	//5. Создадим экземпляр Employee
	emp := Employee{"Bob", 34}
	var workerEmployee Worker = emp // Присваиваем сотрудника в переменную типа Worker
	workerEmployee.Work()
	Describer(workerEmployee) // В резульатте видим, что тип интерфейса определяется структурой, его реализующей,
	//а значение интерфейса - это соответственно значение состояний структуры

	//7. Создадим экземпляр Student
	std := Student{"Mike", 19}
	var workerStudent Worker = std
	workerStudent.Work()
	Describer(workerStudent) // Приянл внутренний тип Student, а значение - равно значению полей экземпляра

	//8. Созаддим набор тех, кто умеет  Work()
	var workers []Worker = []Worker{workerStudent, workerEmployee}
	for _, worker := range workers {
		Describer(worker) // Данная функция вызывается у разных экземпляров благодря тому, кто для ее вызова
		//экземпляру нужно удовлетворить некому контракту (поведенческому паттерну). Если структура экземпляра поддерживает
		// данный паттерн - то у экземпляра 100% можно вызвать все необходимые для этого методы.
	}
}
