package main

import "fmt"

type IntRing struct {
	size  int
	data  []*int
	Start int
	End   int
}

//Конструктор кольцевого массива
func NewIntRing(size int, start int) (*IntRing, error) {
	if start >= size {
		return nil, fmt.Errorf("Стартовая позиция не соответствует размеру кольцевого массива")
	}
	return &IntRing{size, make([]*int, size, size), start, start}, nil
}

//Size - vетод, возвращающий размер кольцевого массива
func (r *IntRing) Size() int {
	return r.size
}

//IsEmpty - Метод, проверяющий кольцо на пустоту
func (r *IntRing) IsEmpty() bool {
	return r.Start == r.End
}

//IsFull - Проверка на полноту, достигнут ли конец буфера
func (r *IntRing) IsFull() bool {
	return (r.End < r.Start && r.Start-r.End == 1) || (r.Start == 0 && r.End == r.Size()-1)
}

//Read - метод чтения элемента
func (r *IntRing) Read() (int, error) {
	if !r.IsEmpty() {
		el := r.data[r.Start]
		for el == nil && r.Start < r.End {
			el = r.data[r.Start]
			r.Start++
		}
		if el == nil {
			return 0, fmt.Errorf("Нет новых данных в буфере")
		}
		return *el, nil
	}
	return 0, fmt.Errorf("Нет новых данных в буфере")
}

//Write - забпись нового элемента в буфер
func (r *IntRing) Write(v int) error {
	if r.IsEmpty() || !r.IsFull() {
		r.data[r.End] = &v
		r.End++
		return nil
	}
	return fmt.Errorf("Буфер полон")
}

//Remove - удаление произвольного элемента
func (r *IntRing) Remove(index int) error {
	if index < 0 || index >= r.Size() { //В скринкасет почему-то используется условие index > r.Size(), строго больше размера
		return fmt.Errorf("Неверный индекс списка")
	}
	r.data[index] = nil
	return nil
}

//Print - метод печати содержимого буфера
func (r *IntRing) Print() {
	if r.Start < r.End {
		for _, el := range r.data[r.Start:r.End] {
			if el != nil {
				fmt.Printf("%d\t", *el)
			}
		}
		fmt.Print("\n")
	} else if r.Start > r.End {
		tempData := append(r.data[r.Start:], r.data[:r.End]...)
		for _, el := range tempData {
			if el != nil {
				fmt.Printf("%d\t", *el)
			}
		}
		fmt.Print("\n")
	} else {
		fmt.Println("Список пуст")
	}
}

func main() {
	ring, err := NewIntRing(5, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	ring.Print()

	err = ring.Write(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ring.Write(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ring.Write(3)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ring.Write(4)
	if err != nil {
		fmt.Println(err)
		return
	}
	ring.Print()
	err = ring.Write(5)
	if err != nil {
		fmt.Println(err)
		return
	}
	ring.Print()
	err = ring.Remove(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	ring.Print()
}
