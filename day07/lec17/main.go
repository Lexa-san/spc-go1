package main

import "fmt"

//1. Структура - заименованный набор полей (состояний), определяющий новый тип данных.

//2. Определение структуры (явное определение)
type Student struct {
	firstName string
	lastName  string
	age       int
}

//3. Если имеется ряд состояний одного типа, можно сделать так
type AnotherStudent struct {
	firstName, lastName, groupName string
	age, courseNumber              int
}

//11. Структура с анонимными полями
type Human struct {
	firstName string
	lastName  string
	string
	int
	bool
}

func PrintStudent(std Student) {
	fmt.Println("==================")
	fmt.Println("FirstName:", std.firstName)
	fmt.Println("LastName:", std.lastName)
	fmt.Println("Age:", std.age)
}

func main() {

	//4. Создание представителей структуры
	stud1 := Student{
		firstName: "Fedya",
		age:       21,
		lastName:  "Petrov",
	}
	PrintStudent(stud1)
	stud2 := Student{"Petya", "Ivanov", 19} // Порядок указания свойств - такой же как в структуре
	PrintStudent(stud2)

	//5. Что если не все поля структуры определить?
	stud3 := Student{
		firstName: "Vasya",
	}
	PrintStudent(stud3)

	//6. Анонимные структуры (структура без имени)
	anonStudent := struct {
		age           int
		groupID       int
		proffesorName string
	}{
		age:           23,
		groupID:       2,
		proffesorName: "Alexeev",
	}
	fmt.Println("AnonStudent:", anonStudent)

	//7. Доступ к состояниям и их модфикация
	studVova := Student{"Vova", "Ivanov", 19}
	fmt.Println("firstName:", studVova.firstName)
	fmt.Println("lastName:", studVova.lastName)
	fmt.Println("age:", studVova.age)
	studVova.age += 2
	fmt.Println("new age:", studVova.age)

	//8. Инициализация пустой структуры
	emptyStudent1 := Student{}
	var emptyStudent2 Student
	PrintStudent(emptyStudent1)
	PrintStudent(emptyStudent2)

	//9. Указатели на экземпляры структур
	studPointer := &Student{
		firstName: "Igor",
		lastName:  "Sidorov",
		age:       22,
	}
	fmt.Println("Value studPointer:", studPointer)
	secondPointer := studPointer
	(*secondPointer).age += 20
	fmt.Println("Value afterPointerModify:", studPointer)
	studPointerNew := new(Student)
	fmt.Println(studPointerNew)

	//10. Работа с доступ к полям структур через указатель
	fmt.Println("Age via (*...).age:", (*studPointer).age)
	fmt.Println("Age via .age:", studPointer.age) //Неявно происходит разыменование указателя studpointer и запрос соотв поля

	//12. Создание экземпляра с анонимными полями структуры
	human := &Human{
		firstName: "Bob",
		lastName:  "Johnson",
		string:    "Additional Info",
		int:       -1,
		bool:      true,
	}

	fmt.Println(human)
	fmt.Println("Anon field string:", human.string)
}
