package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Шаг наращивания счётчика
const step int64 = 1

// Конечное значение счетчика
const endCounterValue int64 = 1000

// Количество вывымаемы горутин
const amountOfGoroutines int64 = 10

// Количество приращений на горутину
const steps4goroutine int64 = endCounterValue / amountOfGoroutines

func main() {

	var counter int64 = 0
	var wg sync.WaitGroup

	increment := func() {
		defer wg.Done()
		for i := 1; i <= int(steps4goroutine); i++ {
			atomic.AddInt64(&counter, step)
		}
	}

	for i := 1; i <= int(amountOfGoroutines); i++ {
		wg.Add(1)
		go increment()
	}
	// Ожидаем поступления сигнала
	wg.Wait()
	// Печатаем результат, надеясь, что будет 1000
	fmt.Println(counter)
}
