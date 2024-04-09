package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkBrackets(input string) bool {
	stack := make([]rune, 0)
	for _, v := range input {
		switch v {
		case '[', '(', '{':
			stack = append(stack, v)
		case ']', ')', '}':
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if (v == ']' && last != '[') || (v == ')' && last != '(') || (v == '}' && last != '{') {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var text string
	fmt.Print("Введите строку для проверки баланса скобок: ")
	text, _ = reader.ReadString('\n')
	fmt.Printf("Результата проверки: %t", checkBrackets(text))

}
