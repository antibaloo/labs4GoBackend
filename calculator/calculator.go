package main

import "fmt"

func main() {
	var a, b float32
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
	switch oper {
	case "+":
		fmt.Print("Результат сложения: ", a+b)
	case "-":
		fmt.Print("Результат вычитания: ", a-b)
	case "*":
		fmt.Print("Результат умножения: ", a*b)
	case "/":
		if b == 0 {
			fmt.Print("Ошибка деления на ноль")
			return
		}
		fmt.Print("Результат деления: ", a/b)
	default:
		fmt.Print("Несуществующий арифметический оператор.")
	}
}
