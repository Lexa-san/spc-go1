package main

import "fmt"

//0. Почему интерфейс -полу-абстрактный тип в Go?

//1. Создадим интерфейс Ездилка
type Rider interface {
	Ride()
	Gas()
	Stop()
}

func main() {
	//2. Создаю экземпляр ездилки
	var r Rider   // ZeroValue - nil, ZeroType - nil
	if r == nil { // Попробуем сравнить сравнить с nil
		fmt.Printf("r is nil. Value %v and type %T\n", r, r)
	}

	r.Ride() // Здесь будет паника, т.к. у экземпляра интерфейса можно вызвать метод Ride()
	//но т.к. значение, которое там лежит - nil - получается nil.Ride()
	//Мораль - если код падает с memory-dereferncing - в 99% случаев - это попытка обратиться к
	//экземпляру интерфейса без присвоенного претендента!

}
