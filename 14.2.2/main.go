package main

import "fmt"

func hashstr(val string) uint64 {
	var sum int
	for i, s := range val {
		sum += (i + 1) * int(s)
	}
	return uint64(sum)
}
func main() {
	fmt.Println(hashstr("abc"))
	fmt.Println(hashstr("cba"))
}
