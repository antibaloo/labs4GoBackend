package main

import (
	"fmt"
)

func hashint64(val int64) uint64 {
	return uint64(val) % 1000
}

func main() {
	fmt.Println(hashint64(-1000), hashint64(1000))
}
