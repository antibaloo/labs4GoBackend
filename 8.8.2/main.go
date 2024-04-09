package main

import (
	"cars"
	"fmt"
)

func main() {
	dodge := cars.NewAuto("USA", "Dodge", "Viper", 445.8, 167.5, 120.9, 314, 507)
	bmw := cars.NewAuto("Europe", "BMW", "X7", 515.1, 200, 180.5, 250, 530)
	mercedes := cars.NewAuto("Europe", "Mercedes", "GLE 500", 481.9, 193.5, 179.6, 250, 435)
	printAutoCharacteristics(dodge)
	printAutoCharacteristics(bmw)
	printAutoCharacteristics(mercedes)
}
func printAutoCharacteristics(a cars.Auto) {
	fmt.Println("Brand:", a.Brand())
	fmt.Println("Model:", a.Model())
	fmt.Printf("Габариты (д/ш/в): %s / %s / %s\n", a.Dimensions().Length(), a.Dimensions().Width(), a.Dimensions().Height())
	fmt.Println("Max speed:", a.MaxSpeed())
	fmt.Println("Engine power:", a.EnginePower())
	fmt.Println()
}
