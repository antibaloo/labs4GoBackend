package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type message struct {
	cond *sync.Cond
	msg  string
}

func main() {
	msg := message{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	// 1
	for i := 1; i <= 3; i++ {
		go func(num string) {
			for {
				msg.cond.L.Lock()
				msg.cond.Wait()
				if msg.msg == num {
					fmt.Printf("hello, i am worker %s. text:%s\n", num, msg.msg)
				} else {
					fmt.Printf("hello, i am worker %s. its not my message\n", num)
				}

				msg.cond.L.Unlock()
			}
		}(fmt.Sprint(i))
	}

	// 2
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter text: ")
	for scanner.Scan() {
		msg.cond.L.Lock()
		msg.msg = scanner.Text()
		msg.cond.L.Unlock()

		msg.cond.Broadcast()
	}

}
