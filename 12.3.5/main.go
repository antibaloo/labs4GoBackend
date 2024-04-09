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

	bubbleSortRecursive(ar)

	fmt.Println(ar)
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func bubbleSortRecursive(ar []int) {
	bubbleSortStep(ar, 1)
}

func bubbleSortStep(ar []int, s int) {
	localCount := 0
	for i := 1; i <= len(ar)-s; i++ {
		if ar[i-1] > ar[i] {
			ar[i-1], ar[i] = ar[i], ar[i-1]
			localCount++
		}
	}
	if localCount == 0 {
		return
	} else {
		bubbleSortStep(ar, s+1)
	}
}
