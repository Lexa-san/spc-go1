//0. Интерфейсы (с точки зрения ООП) - увеличивают уровень абстракции вашего кода
// Засчет увеличения уровня абстракции - можно решать много сторонних проблем, связанных с поддержкой
//понимание и реюзабельностью кода.
// С другой стороны - добавление интерфейсов не решают проблему уменьшения абстрактности.

//1. Что делать, если хочется скрестить 2 интерфейса и создать единый уровень абстракции в коде?
package main

import "fmt"

type PerimeterCalculator interface {
	Perimeter() int
}

type AreaCalculator interface {
	Area() int
}

//2. Соберем новый интерфейс из двух старых
type FigureFeatureCalculator interface {
	PerimeterCalculator // встраиваем интерфейсы
	AreaCalculator
	// Наслдеование интерфейсов
	//Perimeter() int
	//Area() int
}

type Rectangle struct {
	a, b  int
	color string
}

//2. Реализуем интерфейс PerimeterCalculator
func (r Rectangle) Perimeter() int {
	return 2 * (r.a + r.b)
}

//3. Реализуем второй интерфейс AreaCalculator
func (r Rectangle) Area() int {
	return r.a * r.b
}

func main() {
	var pCalc PerimeterCalculator
	fmt.Printf("Value %v and Type %T\n", pCalc, pCalc)
	var aCalc AreaCalculator
	rect := Rectangle{10, 20, "green"}
	pCalc = rect // Стурктура Rectangle удовлетворяет интерфейсу PerimeterCalculator
	fmt.Printf("Value %v and Type %T\n", pCalc, pCalc)
	fmt.Println("Perimeter:", pCalc.Perimeter())

	aCalc = rect // Структура Rectangle удовлетворяет интерфейсу AreaCalculator
	fmt.Println("Area:", aCalc.Area())

	var ffCalc FigureFeatureCalculator
	ffCalc = rect // Структура Rectangle удовлетворяет FigureFeatureCalculator
	fmt.Println(ffCalc.Area())
	fmt.Println(ffCalc.Perimeter())

}
