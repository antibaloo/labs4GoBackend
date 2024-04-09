package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Ффункция ввода общих элементов 2-х массивов
func intercsection(first, second []int) {
	check := make(map[int]bool, len(first)) // создаем map длиной равной длине первого массива
	intercsection := []int{}
	for _, el := range first { //Записываем элементы массива в map, а каччестве ключей значения элементов
		check[el] = true
	}
	for _, el := range second { //Переибираем второй массив
		//если существует ключи в map равный элементу второго массива и он еще не проверялся, добавялем к результату
		if unique, ok := check[el]; ok && unique {
			check[el] = false
			intercsection = append(intercsection, el)
		}
	}
	fmt.Println(intercsection) //выводим результат
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите размер первого массива: ")
	Len, _ := reader.ReadString('\n')
	fLen, err := strconv.Atoi(strings.TrimSpace(Len))

	if err != nil {
		fmt.Println(err)
		fmt.Println("Размер массива должен быть целым положительным числом")
		return
	}

	if fLen < 0 {
		fmt.Println("Размер массива должен быть целым положительным числом")
	}

	fmt.Print("Введите размер второго массива: ")
	Len, _ = reader.ReadString('\n')
	sLen, err := strconv.Atoi(strings.TrimSpace(Len))

	if err != nil {
		fmt.Println("Размер массива должен быть целым положительным числом")
		return
	}

	if sLen < 0 {
		fmt.Println("Размер массива должен быть целым положительным числом")
	}

	fArr := make([]int, fLen)
	sArr := make([]int, sLen)
	fmt.Println("Введите элементы первого массива:")
	for i := 0; i < fLen; i++ {
		number, _ := reader.ReadString('\n')
		fArr[i], err = strconv.Atoi(strings.TrimSpace(number))
		if err != nil {
			fmt.Println("Элемент массива должен быть целым числом")
			return
		}

	}
	fmt.Println("Введите элементы второго массива:")
	for i := 0; i < sLen; i++ {
		number, _ := reader.ReadString('\n')
		sArr[i], err = strconv.Atoi(strings.TrimSpace(number))
		if err != nil {
			fmt.Println("Элемент массива должен быть целым числом")
			return
		}
	}
	intercsection(fArr, sArr)
}
