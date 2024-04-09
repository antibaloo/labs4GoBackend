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
	ar := make([]int, 500)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}
	fmt.Println(ar)
	ar = mergeSort(ar)

	fmt.Println(ar)
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func mergeSort(ar []int) []int {
	if len(ar) == 1 {
		return ar
	}
	res := make([]int, 0)
	left := mergeSort(ar[:len(ar)/2])
	right := mergeSort(ar[len(ar)/2:])
	for {
		if left[0] < right[0] {
			res = append(res, left[0])
			left = left[1:]
			if len(left) == 0 {
				res = append(res, right...)
				break
			}
		} else {
			res = append(res, right[0])
			right = right[1:]
			if len(right) == 0 {
				res = append(res, left...)
				break
			}
		}
	}
	return res
}
