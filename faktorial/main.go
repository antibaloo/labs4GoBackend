package main

import (
	"fmt"
	"time"
)

const n = 62

func main() {
	start := time.Now()
	fmt.Println(itFaktorial(n))
	end := time.Now()
	fmt.Println("Calculated in: ", end.Sub(start))
	start = time.Now()
	fmt.Println(recFaktorial(n))
	end = time.Now()
	fmt.Println("Calculated in: ", end.Sub(start))
}

func recFaktorial(n uint64) uint64 {
	if n < 3 {
		return n
	} else {
		return n * recFaktorial(n-1)
	}
}

func itFaktorial(n uint64) uint64 {
	if n < 3 {
		return n
	}
	var res uint64 = 1
	for i := n; i > 1; i-- {
		res *= i
	}
	return res
}
