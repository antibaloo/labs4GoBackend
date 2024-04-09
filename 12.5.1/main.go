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
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}
	fmt.Println(ar)
	insertionSort(ar)

	fmt.Println(ar)
}

func insertionSort(ar []int) {
	// ваш код здесь
	for i := 1; i < len(ar); i++ {
		for j := i - 1; j >= 0; j-- {
			if ar[i] < ar[j] {
				if j == 0 {
					insert(ar, i, j)
				}
				continue
			}
			if ar[i] >= ar[j] {
				if i-j > 1 {
					insert(ar, i, j+1)
				}
				break
			}
		}
	}
}

func insert(ar []int, from, to int) {
	temp := ar[from]
	for i := from; i > to; i-- {
		ar[i] = ar[i-1]
	}
	ar[to] = temp
}
