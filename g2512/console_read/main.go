package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		lines = append(lines, line)
	}
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}
