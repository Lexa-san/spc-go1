package main

import "fmt"

func main() {
	var (
		personFirstName string
		personLastName  string
		personAge       int
	)

	fmt.Scan(&personFirstName)
	fmt.Scan(&personLastName)
	fmt.Scan(&personAge)

	fmt.Printf("Имя: %s , Фамилия: %s , Возраст: %d . Студент BPS\n", personFirstName, personLastName, personAge)
}
