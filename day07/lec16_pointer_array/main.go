package main

import "fmt"

//1. Указатели на массивы. Почему так делать не надо
func mutation(arr *[3]int) {
	// (*arr)[1] = 909
	// (*arr)[2] = 100000
	//Можно написать и так, т.к. Go сам разыменует указатель на массив (из-за того,что функци принимает *arr)
	arr[1] = 909
	arr[2] = 10000
}

//2. Используйте лучше слайсы (это идеоматично с точки зрения Go)
func mutationSlice(sls []int) {
	sls[1] = 909
	sls[2] = 10000
}

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println("Arr before mutation:", arr)
	mutation(&arr)
	fmt.Println("Arr after mutation:", arr)

	newArr := [3]int{1, 2, 4}
	fmt.Println("newArr before mutationSlice:", newArr)
	mutationSlice(newArr[:])
	fmt.Println("newArr after mutationSlcie:", newArr)
}
