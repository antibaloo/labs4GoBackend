package main

import (
	"fmt"
	"time"
)

func main() {
	var c1count, c2count int
	c1 := make(chan int)
	c2 := make(chan int)
	numbers2send := 50
	go func() {
		defer close(c1)
		defer close(c2)
		for i := 1; i <= numbers2send; i++ {
			if i%2 <= 0 {
				time.Sleep(10 * time.Millisecond)
				c1 <- i
			} else {
				time.Sleep(20 * time.Millisecond)
				c2 <- i
			}
		}
	}()

	for {
		select {
		case num, ok := <-c1:
			if ok {
				fmt.Println("from 1 channel", num)
				c1count++
			} else {
				fmt.Printf("Recieved from 1 channel: %d. Recieved from 2 channel: %d.", c1count, c2count)
				return
			}
		case num, ok := <-c2:
			if ok {
				fmt.Println("from 2 channel", num)
				c2count++
			} else {
				fmt.Printf("Recieved from 1 channel: %d. Recieved from 2 channel: %d.", c1count, c2count)
				return
			}
		default:
			time.Sleep(30 * time.Millisecond)
			fmt.Println("No data in channels, current time is", time.Now().Format(time.RFC3339))
		}
	}

}
