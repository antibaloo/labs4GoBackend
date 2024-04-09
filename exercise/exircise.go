package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const (
		january = iota + 1
		february
		march
		april
		may
	)
	fmt.Println(january, february, march, april, may)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input text: ")
	text, _ := reader.ReadString('\n')

	fmt.Print("Your text is: ", text)
	var number int

	fmt.Scanln(&number)

	fmt.Println(number + 30)
}
