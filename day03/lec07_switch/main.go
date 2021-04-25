package main

import "fmt"

func main() {
	// Switch
	var price int
	fmt.Scan(&price)

	// В switch-case запрещены дублирующиеся условия в case-ах
	switch price {
	case 100:
		fmt.Println("First case")
	case 110:
		fmt.Println("Second case")
	case 120:
		fmt.Println("Third case")
	default:
		fmt.Println("Default case")
	}

	// multiple variaples
	var ageGroup string = "q"
	switch ageGroup {
	case "a", "b", "c":
		fmt.Println("First case")
	case "d", "e":
		fmt.Println("Second case")
	case "f":
		fmt.Println("Third case")
	default:
		fmt.Println("Default case")
	}

	//
	var age int
	fmt.Scan(&age)
	switch {
	case age < 18:
		fmt.Println("Too young")
	case age >= 18 && age <= 30:
		fmt.Println("OK")
	case age > 30 && age <= 100:
		fmt.Println("Too old")
	default:
		fmt.Println("Good")
	}

	// В момент выполнения fallthrough условия следующего case не проверяется
	var number int
	fmt.Scan(&number)
outer:
	switch {
	case number < 100:
		fmt.Printf("%d is less then 100\n", number)
		if number%2 == 0 {
			break outer
		}
		fallthrough
	case number > 200:
		fmt.Printf("%d is greater then 200\n", number)
		// fallthrough
	case number > 1000:
		fmt.Printf("%d is greater then 1000\n", number)
		// fallthrough
	default:
		fmt.Printf("%d default\n", number)
	}

	// гадость с терминацией чикла for из switch
	var i int
uberloop:
	for {
		fmt.Scan(&i)
		switch {
		case i%2 == 0:
			fmt.Printf("value is %d and it's even\n", i)
			break uberloop
		}
	}
	fmt.Println("END")
}
