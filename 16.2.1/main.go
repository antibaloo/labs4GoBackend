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
	fmt.Print("Введите количество горутин: ")
	nStr, _ := reader.ReadString('\n')
	n, err := strconv.Atoi(strings.Replace(nStr, "\r\n", "", -1))
	if err != nil {
		fmt.Println("необходимо ввести целое положительное число")
		return
	}
	for i := 1; i <= n; i++ {
		go routine(i)
	}
	fmt.Print("Нажмите Enter для завершения работы.")
	_, _ = reader.ReadString('\n')
}

func routine(i int) {
	for {
		time.Sleep(time.Second)
		fmt.Printf("%d ", i)
	}
}
