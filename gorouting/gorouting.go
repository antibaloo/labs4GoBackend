package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func MyFunc(i int) error {
	defer wg.Done()
	<-time.After(time.Millisecond * 100)
	fmt.Println(time.Now(), " Hello from ", i)
	return nil
}
func main() {
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go MyFunc(i)
	}
	wg.Wait()
	return
}
