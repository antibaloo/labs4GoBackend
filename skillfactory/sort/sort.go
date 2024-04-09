package sort

import "math/rand"

//Сортировка пузырьком без проверки количества перестановок
func bubbleSortwoCheck(ar []int) {
	for i := 1; i < len(ar); i++ {
		for j := 1; j <= len(ar)-i; j++ {
			if ar[j-1] > ar[j] {
				ar[j-1], ar[j] = ar[j], ar[j-1]
			}
		}
	}
}

//Сортировка пузырьком с проверкой количества перестановок
func bubbleSortwCheck(ar []int) {
	for i := 1; i < len(ar); i++ {
		swapCount := 0
		for j := 1; j <= len(ar)-i; j++ {
			if ar[j-1] > ar[j] {
				ar[j-1], ar[j] = ar[j], ar[j-1]
				swapCount++
			}
		}
		if swapCount == 0 {
			break
		}
	}
}

//Сортировка пузырьком с использованием рекурсии
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

func selectionSort(ar []int) {
	var min, minIndex int
	for i := 0; i <= len(ar)-2; i++ {
		min, minIndex = ar[i], i
		for j := i; j < len(ar); j++ {
			if ar[j] < min {
				min = ar[j]
				minIndex = j
			}
		}
		if i != minIndex {
			ar[i], ar[minIndex] = ar[minIndex], ar[i]
		}
	}
}

//Двунапрвленная сортировкавыбором
func selectionSortByMinMax(ar []int) {
	var min, max, minIndex, maxIndex int
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
		}
		if i != minIndex {
			ar[i], ar[minIndex] = ar[minIndex], ar[i]
		}
		if len(ar)-1-i != maxIndex {
			ar[len(ar)-1-i], ar[maxIndex] = ar[maxIndex], ar[len(ar)-1-i]
		}
	}
}

func insertionSort(ar []int) {
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
