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
	ar := make([]int, 51)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}
	fmt.Println(ar)
	selectionSortByMinMax(ar)
	fmt.Println(ar)
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func selectionSortByMinMax(ar []int) {
	var min, max, minIndex, maxIndex, iteration int
	for i := 0; i < len(ar); i++ {
		min, minIndex = ar[i], i
		max, maxIndex = ar[len(ar)-1-i], len(ar)-1-i
		for j := i; j < len(ar)-i; j++ {
			if ar[j] < min {
				min, minIndex = ar[j], j
			}
			if ar[j] > max {
				max, maxIndex = ar[j], j
			}
			iteration++
		}
		if i != minIndex {
			ar[i], ar[minIndex] = ar[minIndex], ar[i]
		}
		if len(ar)-1-i != maxIndex {
			ar[len(ar)-1-i], ar[maxIndex] = ar[maxIndex], ar[len(ar)-1-i]
		}
	}
	fmt.Println(iteration)
}
