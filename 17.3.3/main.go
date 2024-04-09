package main

import (
	"fmt"
	"sync"
)

const interationAmount int64 = 1000

func main() {
	var counter int64 = 0
	var c = sync.NewCond(&sync.Mutex{})
	increment := func() {
		c.L.Lock()
		defer c.L.Unlock()
		counter++
		if counter == interationAmount {
			c.Signal()
		}
	}
	for i := 1; i <= int(interationAmount); i++ {
		go increment()
	}
	c.L.Lock()
	c.Wait()
	c.L.Unlock()
	fmt.Println(counter)
}
