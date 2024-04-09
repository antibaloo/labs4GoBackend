package main

import (
	"calc"
	"fmt"
)

func main() {
	var a, b float64
	var oper string
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Printf("Ошибка ввода первого значения: %v", err)
		return
	}
	_, err = fmt.Scanln(&oper)
	if err != nil {
		fmt.Printf("Ошибка ввода оператора: %v", err)
		return
	}
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Printf("Ошибка ввода второго значения: %v", err)
		return
	}
	calculator := calc.NewCalculator()
	result, err := calculator.Calculate(a, b, oper)
	if err != nil {
		fmt.Printf("Произошла ошибка: %v\n", err)
	} else {
		fmt.Printf("Результат вычислений: %v\n", result)
	}
}
