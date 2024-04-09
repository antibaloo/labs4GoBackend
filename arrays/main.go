package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	// var intArr = []int{-15, 1, 3, -4, -1, 1, 3, 3}
	// var sorted = []int{1, 1, 2, 3, 4, 5, 6, 7, 7, 8, 9, 9, 10}
	var unsorted []int
	/*maxNeg, err := findMaxNegative(intArr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(maxNeg)
	fmt.Println(findMostOftenRepeated(intArr))
	fmt.Println(findMostOftenRepeatedWM(intArr))
	fmt.Println(sorted, checkSliceIsSorted(sorted))
	*/
	for i := 0; i < 1000; i++ {
		unsorted = append(unsorted, rand.Intn(100))
	}
	fmt.Println(unsorted)
	start := time.Now()
	sorted := sort(unsorted)
	fmt.Println(sorted, checkSliceIsSorted(sorted))
	end := time.Now()
	fmt.Println(end.Sub(start).Nanoseconds())
	start = time.Now()
	sortedM := mergeSort(unsorted)
	fmt.Println(sortedM, checkSliceIsSorted(sortedM))
	end = time.Now()
	fmt.Println(end.Sub(start).Nanoseconds())
}
func findMaxNegative(array []int) (int, error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found maxNegative in empty slice")
	}

	maxN := 0
	for _, val := range array {
		if val < 0 && val < maxN {
			maxN = val
		}
	}
	if maxN == 0 {
		return 0, fmt.Errorf("where is no negative numbers in slice")
	}
	return maxN, nil
}

func findMostOftenRepeated(array []int) (mostOften int, err error) {
	iter := 0
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found repeated numbers in empty slice")
	}

	var maxIndex, maxCount = 0, 0
	for i, number := range array {
		currentCount := 0
		for _, numberToCompare := range array[i:] {
			iter++
			if number == numberToCompare {
				currentCount++
			}
		}

		if currentCount > maxCount {
			maxIndex = i
			maxCount = currentCount
		}
	}
	fmt.Println("Iter count:", iter)
	return array[maxIndex], nil
}
func findMostOftenRepeatedWM(array []int) (int, error) {
	iter := 0
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found repeated numbers in empty slice")
	}
	mo := make(map[int]int, len(array))
	for _, number := range array {
		iter++
		if _, ok := mo[number]; ok {
			mo[number]++
		} else {
			mo[number] = 1
		}
	}
	//fmt.Println(mo)
	maxNumber := 0
	maxCount := 0
	for number, count := range mo {
		iter++
		if count > maxCount {
			maxCount = count
			maxNumber = number
		}
	}
	fmt.Println("Iter count:", iter)
	return maxNumber, nil
}
func checkSliceIsSorted(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}
func sort(a []int) []int {
	res := make([]int, len(a))
	copy(res, a)
	wasChanged := true
	for wasChanged {
		wasChanged = false
		for i := 1; i < len(res); i++ {
			if res[i] < res[i-1] {
				after := res[i+1:]
				res = append(res[:i-1], res[i], res[i-1])
				res = append(res, after...)
				wasChanged = true
			}
		}
	}
	return res
}
func mergeSort(a []int) []int {
	res := make([]int, len(a))
	copy(res, a)
	doMergeSort(res)
	return res
}

func doMergeSort(items []int) {
	length := len(items)
	if length == 1 {
		return
	}
	ILeft := length / 2
	left := make([]int, ILeft)
	copy(left, items[:ILeft])
	IRight := length - ILeft
	right := make([]int, IRight)
	copy(right, items[ILeft:])
	doMergeSort(left)
	doMergeSort(right)
	merge(left, right, items)
}

func merge(left []int, right []int, items []int) {
	var l, r, i int
	for (l < len(left)) && (r < len(right)) {
		if left[l] < right[r] {
			items[i] = left[l]
			l++
		} else {
			items[i] = right[r]
			r++
		}
		i++
	}
	lenght := len(left) - l
	copy(items[i:i+lenght], left[l:])
	i += lenght
	lenght = len(right) - r
	copy(items[i:i+lenght], right[r:])
}
