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

	selectionSortByMax(ar)
	fmt.Println(ar)
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func selectionSortByMax(ar []int) {
	var max, maxIndex, iteration int
	for i := len(ar) - 1; i > 1; i-- {
		max, maxIndex = ar[i], i
		for j := 0; j <= i; j++ {
			if ar[j] > max {
				max, maxIndex = ar[j], j
			}
			iteration++
		}
		if i != maxIndex {
			ar[i], ar[maxIndex] = ar[maxIndex], ar[i]
		}
	}
	fmt.Println(iteration)
}
