package main

import (
	"fmt"
	"sync"
)

// Структура счетчика
type Counter struct {
	wg                          sync.WaitGroup
	counter, max, numGoroutines int
	send, balance               chan int // Каналы для взаимодействия
}

// Конструктор экземпляря счетчика
func NewCounter(max, num int) *Counter {
	return &Counter{
		wg:            sync.WaitGroup{},
		counter:       0,
		max:           max,
		numGoroutines: num,
		send:          make(chan int),
		balance:       make(chan int),
	}
}

// Возвращает значение счетчика
func (c *Counter) Balance() int {
	return c.counter
}

// Запускает основную горутину, которая слушает канал send и увеличивает счетчик,
// а потом отправляет в канал обратной связи его значение
func (c *Counter) StartCounter() {
	c.wg.Add(1)
	go func() {
		c.balance <- c.counter //Отправка нулевого значения для запуская взаимодействия
		defer c.wg.Done()
		defer close(c.balance) // Отложенное закрытие канала обратной связи, канал закроется при завершении горутины
		for step := range c.send {
			c.counter += step
			c.balance <- c.counter
		}
	}()
}

// Метод стартует блок горутин в заданном количестве, в горутине читает канал обратной связи, если предельное значение
// не достигнуто, то канал send, что приведет к завершению основной горутины
func (c *Counter) StartSenders() {
	for i := 1; i <= c.numGoroutines; i++ {
		go func() {
			for counter := range c.balance {
				if counter < c.max {
					c.send <- 1
				} else {
					close(c.send)
				}
			}
		}()
	}
}

func (c *Counter) Wait() {
	c.wg.Wait()
}

const (
	maxCounter    = 123
	numGoroutines = 13
)

func main() {
	counter := NewCounter(maxCounter, numGoroutines)
	counter.StartCounter()
	counter.StartSenders()
	counter.Wait()
	fmt.Println(counter.Balance())
}
