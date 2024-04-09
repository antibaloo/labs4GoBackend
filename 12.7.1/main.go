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
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}
	fmt.Println(ar)
	quickSort(ar)

	fmt.Println(ar)
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func quickSort(ar []int) {
	if len(ar) <= 1 {
		return
	}
	if len(ar) == 2 {
		if ar[1] < ar[0] {
			ar[0], ar[1] = ar[1], ar[0]
		}
		return
	}
	index := rand.Intn(len(ar) - 1)
	swap(ar, index, len(ar)-1)
	left := 0
	for i := 0; i < len(ar)-1; i++ {
		if ar[i] < ar[len(ar)-1] {
			swap(ar, i, left)
			left++
		}
	}
	swap(ar, left, len(ar)-1)

	quickSort(ar[:left])
	quickSort(ar[left+1:])
}

func swap(ar []int, i, j int) {
	ar[i], ar[j] = ar[j], ar[i]
}
