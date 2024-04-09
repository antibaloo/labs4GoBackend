package main

import (
	//"bufio"

	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type item struct {
		Global_id int64
	}
	var testData []item
	var sum int64
	file, err := os.Open("data-20190514T0100.json")
	if err != nil {
		fmt.Print("Error reading file")
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&testData)
	for _, item := range testData {
		fmt.Println(item.Global_id)
		sum += item.Global_id
	}
	fmt.Println(sum)
}
