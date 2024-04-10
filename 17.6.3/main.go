package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	c := make(chan int)
	wg.Add(1)
	go func() {
		defer close(c)
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			time.Sleep(100 * time.Millisecond)
			c <- i
		}
	}()
	wg.Add(1)
	go func(c <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := range c {
			fmt.Printf("%v ", i)
		}
	}(c, &wg)
	wg.Wait()
}
