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
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}

	bubbleSort(ar)

	fmt.Println(ar)
}

func bubbleSort(ar []int) {
	allSwaps := 0
	iteration := 0
	for i := 1; i < len(ar); i++ {
		swapCount := 0
		for j := 1; j <= len(ar)-i; j++ {
			iteration++
			if ar[j-1] > ar[j] {
				ar[j-1], ar[j] = ar[j], ar[j-1]
				swapCount++
			}
		}
		allSwaps += swapCount
	}
	fmt.Println("Всего Итераций:", iteration)
	fmt.Println("Всего замен:", allSwaps)
}
