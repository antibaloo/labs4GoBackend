package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите Время выполнения программы в секундах: ")
	nStr, _ := reader.ReadString('\n')
	n, err := strconv.Atoi(strings.Replace(nStr, "\r\n", "", -1))
	if err != nil {
		fmt.Println("необходимо ввести целое положительное число")
		return
	}
	counter := 0
	for i := 0; i < n; i++ {
		for j := 0; j < 5; j++ {
			time.Sleep(time.Millisecond * time.Duration(200))
			fmt.Printf("%v - %v\n", time.Now().Format("15:04:05.000"), counter)
		}
		counter++
	}
}

/*
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите Время выполнения программы в секундах: ")
	nStr, _ := reader.ReadString('\n')
	n, err := strconv.Atoi(strings.Replace(nStr, "\r\n", "", -1))
	if err != nil {
		fmt.Println("необходимо ввести целое положительное число")
		return
	}
	counter := 0
	go func() {
		for {
			time.Sleep(time.Millisecond * time.Duration(200))
			fmt.Printf("%v - %v\n", time.Now().Format("15:04:05.000"), counter)
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second)
			counter++
		}
	}()
	time.Sleep(time.Second * time.Duration(n))
}
*/
