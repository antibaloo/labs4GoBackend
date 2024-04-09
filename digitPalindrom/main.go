package main

import "fmt"

func main() {
	var number int

	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	if isPalindrom(number) {
		fmt.Println("Число:", number, "палиндром")
	} else {
		fmt.Println("Число:", number, "Не палиндром")
	}
}
func isPalindrom(n int) bool {
	var rn int
	wp := n
	for wp != 0 {
		rn = rn*10 + wp%10
		wp /= 10
	}
	return n == rn
}
