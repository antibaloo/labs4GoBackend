package main

import "fmt"

func main() {
	str := "abcccccc"

	fmt.Printf("'%s' -> '%s", str, strArc(str))
	fmt.Print("\n")
	fmt.Printf("'%s' -> '%s", str, Archive(str))
}
func strArc(str string) string {
	var last rune
	var count int
	var result string
	for _, s := range str {
		if s != last {
			if count > 0 {
				result += fmt.Sprintf("%c%d", last, count)

			}
			last = s
			count = 1
		} else {
			count++
		}
	}
	result += fmt.Sprintf("%c%d", last, count)
	return result
}
func Archive(s string) string {
	var curr rune
	currLen := 0
	res := ""
	for _, v := range s {
		if v != curr {
			if currLen > 0 {
				res += fmt.Sprintf("%c%d", curr, currLen)
			}
			curr = v
			currLen = 1
		} else {
			currLen++
		}
	}
	res += fmt.Sprintf("%c%d", curr, currLen)
	return res
}
