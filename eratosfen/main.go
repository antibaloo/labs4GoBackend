package main

import (
	"fmt"
	"time"
)

const max = 1500000

func main() {
	start := time.Now()
	var sMap = make(map[int]string)
	var sVector = make([]int, 0, max-1)
	var result = make([]int, 0, max-1)
	for i := 2; i <= max; i++ {
		sMap[i] = ""
		sVector = append(sVector, i)
	}
	//fmt.Println(sMap)
	//fmt.Println(sVector)

	for i := 0; i < len(sVector); i++ {
		if checkSimple(sVector[i], sMap) {
			result = append(result, sVector[i])
		}
	}
	fmt.Printf("Ряд простых чисел от 2 до %d\n", max)
	fmt.Println(result)
	end := time.Now()
	fmt.Println(len(result))
	fmt.Printf("Выполнено за %s", end.Sub(start))
}
func checkSimple(i int, m map[int]string) bool {
	//fmt.Println("serching", i)
	if _, ok := m[i]; ok {
		//fmt.Println("found", i)
		for j := 2; j*i <= max; j++ {
			//fmt.Println("searching", j, "*", i)
			if _, ok := m[j*i]; ok {
				delete(m, j*i)
				//fmt.Println("deleted", j*i)
			} else {
				//fmt.Println("not found", j*i)
			}
		}
	} else {
		//fmt.Println("not found", i)
		return false
	}
	return true
}
