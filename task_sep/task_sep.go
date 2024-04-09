package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("task.data")
	rd := bufio.NewReader(file)
	reader := csv.NewReader(rd)
	reader.Comma = ';' // Ставим разделитель ';'Читаем данные в цикле:
	record, _ := reader.Read()
	for num, item := range record {
		if item == "0" {
			fmt.Println(num + 1)
			break
		}
	}
	/*for {
		record, err := reader.Read()
	}

	// record - теперь это просто массив строк (array []string).  Проходим его посредством for и range:

	for num, item := range record {
		if item == "0" {
			fmt.Println(num + 1)
			break
		}
	}*/
}
