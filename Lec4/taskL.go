package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var (
		login        string
		email        string
		loginIsValid bool
		emailIsValid bool
	)

	fmt.Scan(&login, &email)

	loginIsValid = utf8.RuneCountInString(login) >= 10 && !strings.Contains(login, "@")
	emailIsValid = strings.Contains(email, "@") && strings.Contains(email, ".")
	// fmt.Println(loginIsValid, emailIsValid)

	if !loginIsValid {
		fmt.Println("Некорректный логин")
		return
	} else if !emailIsValid {
		fmt.Println("Некорректная почта")
		return
	}

	fmt.Println("ОК")
}
