package main

import "fmt"

func main() {
	//yearDetection()
	dayDetection()
	ages := map[string]int{"John": 21, "Mary": 19, "Bob": 80}
}

func yearDetection() {
	var year uint
	fmt.Println("введите год: ")
	fmt.Scan(&year)
	if year%400 == 0 || ((year%100 != 0) && (year%4 == 0)) {
		fmt.Println("Это високосный год")
	} else {
		fmt.Println("Это невисокосный год")
	}
	fmt.Printf("%d \n", year)
}
func dayDetection() {
	var day uint
	fmt.Print("введите номер дня недели (1-7): ")
	fmt.Scan(&day)
	var dayText string
	switch day {
	case 1:
		dayText = "понедельник"
	case 2:
		dayText = "вторник"
	case 3:
		dayText = "среда"
	case 4:
		dayText = "четверг"
	case 5:
		dayText = "пятница"
	case 6:
		dayText = "суббота"
	case 7:
		dayText = "воскресенье"
	default:
		dayText = "неправильно введен номер дня"
	}
	fmt.Println("День недели: ", dayText)
}
