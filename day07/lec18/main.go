package main

import "fmt"

//1. Вложенные структуры (вложение структур). Это использование одной структуры, как тип поля
//в другйо структуре
type University struct {
	age       int
	yearBased int
	infoShort string
	infoLong  string
}

type Student struct {
	firstName  string
	lastName   string
	university University
}

//4. Встроенные структуры (когда мы добавляем поля одной структуры к другой)
type Professor struct {
	firstName string
	lastName  string
	age       int
	greatWork string
	//papers     map[string]string - добавление этого поля делает структуру несравнимой
	University // В этом месте происходит добавление всех полей структуры Uni в Professor
}

func main() {
	//2. Создание экземпляров структур с вложением
	stud := Student{
		firstName: "Fedya",
		lastName:  "Petrov",
		university: University{
			yearBased: 1991,
			infoShort: "cool University",
			infoLong:  "very cool University",
		},
	}

	//3. Получение доступа к вложенным полям структур
	fmt.Println("FirstName:", stud.firstName)
	fmt.Println("LastName:", stud.lastName)
	fmt.Println("Year based Uni:", stud.university.yearBased)
	fmt.Println("Long info:", stud.university.infoLong)

	//5. Создание экземпляра с встраиванием структур
	prof := Professor{
		firstName: "Anatoly",
		lastName:  "Smirnov",
		age:       125,
		greatWork: "Ultimate C programming",
		University: University{
			yearBased: 1734,
			infoShort: "short Info",
			age:       2021 - 1734,
		},
	}
	//6. Обращение к состояниям с встроенной структурой
	fmt.Println("FirstName:", prof.firstName)
	fmt.Println("Year based:", prof.yearBased)
	fmt.Println("Info Short:", prof.infoShort)
	fmt.Println("Age:", prof.University.age) //prof.age - получим доступ к полю ВЫШЕЛЕЖАЩЕЙ СТРУКТУРЫ

	//7. Сравнение экземпляров ==
	//При сравнении экзмеляров происходит сравнение всех их полей друг с другом
	profLeft := Professor{}
	profRight := Professor{}

	fmt.Println(profLeft == profRight)
	//8. Если ХОТЯ БЫ ОДНО ИЗ ПОЛЕЙ СТРУКТУР - НЕ СРАВНИМО - то и вся структура несравнима
}
