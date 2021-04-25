package main

import "fmt"

//0. Структура -  явно декларированный заименованный набор СОСТОЯНИЙ.
//Структра , сходя из своего описания, отвечает на вопрос - из ЧЕГО я должен состоять,
// чтобы считаться тем ТИПОМ, который декларируется структурой?
// Структура - описывает ПАТТЕР СОСТОЯНИЯ.

//1. Интерфейсы - явно декларированный набор сигнатур ПОВЕДЕНИЙ (чаще всег ов виде набора методов), удовлетворив который,
// можно считаться типом, котоырй декларирует интерфейс.
// Интерфейсы в Go декларируют ПОЛУ-АБСТРАКТНЫЙ тип.
// Отвечает на вопрос - что я должен УМЕТЬ ДЕЛАТЬ, чтобы считаться тем ТИПОМ, котоырй декларирует интерфейс?
// Интерфейс - описывает ПАТТЕРН ПОВЕДЕНИЯ.

//2. Объявление интерфейса
type FigureBuilder interface {
	//Набор сигнатур методов, которые необходимо реализовать в структуре-претенденте
	//Во-первых , у претендента должен быть метод Area() возвращающий int
	Area() int
	//Во-вторых, у претендента должен быть метод Perimeter() возвращающий int
	Perimeter() int
}

//3. Декларируем претендентов
//3.1 Первый претендент - это прямоугольник
// У него есть 2 метода -
//Area() int
//Perimter() int
//Когда эти методы реализованы , говорят, что RECTANGLE УДОВЛЕТВОРЯЕТ УСЛОВИЯМ ИНТЕРФЕЙСА FigureBuilder
//RECTANGLE РЕАЛИЗУЕТ ИНТЕРФЕЙС FigureBuilder
type Rectangle struct {
	length, width int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

// 3.2 Второй претендент - это окружность
// У нее есть 2 метода -
//Area() int
//Perimter() int
//Когда эти методы реализованы , говорят, что CIRCLE УДОВЛЕТВОРЯЕТ УСЛОВИЯМ ИНТЕРФЕЙСА FigureBuilder
//CIRCLE РЕАЛИЗУЕТ ИНТЕРФЕЙС FigureBuilder
type Circle struct {
	radius int
}

func (c Circle) Area() int {
	return 3 * c.radius * c.radius
}

func (c Circle) Perimeter() int {
	return 2 * 3 * c.radius
}

func main() {

	//4. Создадим по 3 экземпляра
	r1 := Rectangle{10, 20}
	r2 := Rectangle{30, 50}
	r3 := Rectangle{1, 1}
	c1 := Circle{5}
	c2 := Circle{10}
	c3 := Circle{15}

	//5. Посчитаем общую площадь этих  фигур
	rectangles := []Rectangle{r1, r2, r3}
	totalSumAreaOfRectangles := 0
	for _, rect := range rectangles {
		totalSumAreaOfRectangles += rect.Area()
	}

	circles := []Circle{c1, c2, c3}
	totalSumAreaOfCircles := 0
	for _, circ := range circles {
		totalSumAreaOfCircles += circ.Area()
	}

	fmt.Println("Total area is:", totalSumAreaOfRectangles+totalSumAreaOfCircles)

	//6 . Теперь воспользуемся интерфейсом, указанным выше
	figures := []FigureBuilder{r1, r2, r3, c1, c2, c3} //Объявляю слайс экземпляров, удовлетворяющих интерфейсу FigureBuilder
	// С другой стороны, кажется, что это слайс - каких-то определенных типов

	//7. Посчитаем обзую площадь
	total := 0
	for _, fig := range figures {
		total += fig.Area()
	}
	fmt.Println("Total area:", total)
	//8 . Пояснение - так как каждый экземпляр удовлетворяет интерфейсу FigureBuilder (объявляющий пол-абстрактный тип)
	// у каждого из слайса figures можно 100% вызывать метод Area() (который точно вернет int), или Perimter()
	// кототрый тоже 100% вернет int

}
