package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	start := time.Now()
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}

	selectionSort(ar)
	fmt.Println(ar)
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func selectionSort(ar []int) {
	var min, minIndex int
	for i := 0; i <= len(ar)-2; i++ {
		min, minIndex = ar[i], i
		for j := i; j < len(ar); j++ {
			if ar[j] < min {
				min = ar[j]
				minIndex = j
			}
		}
		if i != minIndex {
			ar[i], ar[minIndex] = ar[minIndex], ar[i]
		}
	}
}
