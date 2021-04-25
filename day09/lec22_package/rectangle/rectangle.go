package rectangle // Т.к. проект main на весь проект один, тут мы создаем новый сторонний пакет

type Rectangle struct {
	A, B  int    //Начинается с большой буквы
	color string //Начинается с маленькой буквы. Данное поле больше НЕ ЭКСПОРТИРУЕМО
}

func New(newA int, newB int, newColor string) *Rectangle {
	return &Rectangle{newA, newB, newColor}
}

func (r *Rectangle) Perimeter() int {
	return 2 * (r.A + r.B)
}
