package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err // Если по какой-то причине мы получили ошибку, проигнорируем эту итерацию
	}

	if info.IsDir() {
		return nil // Проигнорируем директории
	}

	buf, _ := os.Open(path /*info.Name()*/)

	defer buf.Close()

	r := csv.NewReader(buf)

	s, err = r.ReadAll()

	if err != nil {

		fmt.Println(info.Name(), "пусто!")

	} else {
		if len(s) == 10 {
			fmt.Println("нашёл: ", s[4][2])
		}

	}
	return nil
}

var s [][]string

func main() {
	const root = "." // Файлы моей программы находятся в другой директории

	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}

	// Name: file1     Size: 6 byte    Path: test/dir1/file1
	// Name: file2     Size: 6 byte    Path: test/dir1/file2
	// Name: file3     Size: 6 byte    Path: test/dir2/file3
	// Name: file4     Size: 6 byte    Path: test/dir3/file4
	// Name: file5     Size: 6 byte    Path: test/dir3/file5
	// Name: file6     Size: 6 byte    Path: test/dir3/file6
}
