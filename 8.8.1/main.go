package main

import (
	"electronic"
	"fmt"
)

func main() {
	applePhone := electronic.NewApplePhone("Iphone 14 Pro Max", "Ios 16")
	androidPhone := electronic.NewAndroidPhone("Samsung", "M32", "Android 14")
	radioPhone := electronic.NewRadioPhone("Panasonic", "PSC-45", 15)

	printCharacteristics(applePhone)
	printCharacteristics(androidPhone)
	printCharacteristics(radioPhone)
}

func printCharacteristics(p electronic.Phone) {
	switch p.(type) {
	case electronic.Smartphone:
		fmt.Printf("Brand: %s\nModel: %s\nType: %s\nOS: %s\n\n", p.Brand(), p.Model(), p.Type(), p.(electronic.Smartphone).OS())
	case electronic.StationPhone:
		fmt.Printf("Brand: %s\nModel: %s\nType: %s\nOS: %d\n\n", p.Brand(), p.Model(), p.Type(), p.(electronic.StationPhone).ButtonsCount())
	default:
		fmt.Println("phoneType mismatch")
	}
}

/*
func printCharacteristics(p electronic.Phone) {
	switch p.(type) {
	case electronic.Smartphone, electronic.StationPhone:
		fmt.Printf("%s\n", p)
	default:
		fmt.Println("phoneType mismatch")
	}
}*/
