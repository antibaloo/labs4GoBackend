package main

import "fmt"

func main() {
	var list = []int{0, 1, 2, 3, 4, 5}
	fmt.Println(Insert(list, 20, -1))
	fmt.Println(Delete(list, 1))
}

// функция возвращает срез, в котором элемент el вставлен по позицию index, если индек < 0,
// то возвращаем исходный срез; если index > длины среза, то элемент добавляется в конец среза
func Insert(l []int, el int, index int) []int {
	/*result := []int{}
	for i, r := range l {
		if i == index {
			result = append(result, el)
		}
		result = append(result, r)
	}
	return result*/
	if index < 0 {
		return l
	}
	if index > len(l) {
		return append(l, el)
	}
	result := make([]int, len(l[0:index]))
	copy(result, l[0:index])
	return append(append(result, el), l[index:]...)
}
func Delete(l []int, index int) []int {

	return append(l[0:index], l[index+1:]...)
}
