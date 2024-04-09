package main

import "fmt"

//Ошибка при передаче несуществующего в списке индекса
var ErrorWrongListIndex = fmt.Errorf("Неверный индекс списка")

//Описание типа - узел списка
type IntNode struct {
	Value int
	Next  *IntNode
}

//Описание метода создания узла списка
func NewIntNode(value int) *IntNode {
	return &IntNode{value, nil}
}

//Описание типа - односвязный список целых чисел
type IntList struct {
	size int
	Head *IntNode
}

//Описание метода создания списка целых чисел
func NewIntList() *IntList {
	return &IntList{0, nil}
}

//Опсиание метода возвращающего размер односвязного списка
func (l IntList) Size() int {
	return l.size
}

//Описание метода получения элемента списка по индексу
func (l IntList) Get(index int) (*IntNode, error) {
	if index < 0 || index >= l.Size() {
		return nil, ErrorWrongListIndex
	}
	node := l.Head
	for i := 1; i <= index; i++ {
		node = node.Next
	}
	return node, nil
}

//Описание метода присвоения значения элементу списка
func (l *IntList) Set(el int, index int) error {
	if index < 0 || index >= l.Size() {
		return ErrorWrongListIndex
	}
	node, err := l.Get(index)
	if err != nil {
		return err
	}
	node.Value = el
	return nil
}

//Описание метода добавления дового элемента в начало списка
func (l *IntList) Add(el int) {
	newNode := NewIntNode(el)
	if l.Head == nil {
		l.Head = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
	l.size++
}

//Описание метода вставки нового элемента в произвольную позицию списка
func (l *IntList) Insert(el int, index int) error {
	if index < 0 || index > l.Size() {
		return ErrorWrongListIndex
	}
	if index == 0 {
		l.Add(el)
		return nil
	}
	newNode := NewIntNode(el)
	node, err := l.Get(index - 1)
	if err != nil {
		return err
	}
	newNode.Next = node.Next
	node.Next = newNode
	l.size++
	return nil
}

//Описание метода удаления элемента списка из произвольной позиции
func (l *IntList) Remove(index int) error {
	if index < 0 || index >= l.Size() {
		return ErrorWrongListIndex
	}
	if index == 0 {
		l.Head = l.Head.Next
	} else {
		node, err := l.Get(index - 1)
		if err != nil {
			return err
		}
		node.Next = node.Next.Next
	}
	l.size--
	return nil
}

//Описание метода вывода на печать односвязного списка целых чисел
func (l IntList) Print() {
	node := l.Head
	if node != nil {
		for node != nil {
			fmt.Printf("%d\t", node.Value)
			node = node.Next
		}
		fmt.Printf("\n")
	} else {
		fmt.Println("Список пуст!")
	}
}

func main() {
	list := NewIntList()
	list.Print()
	list.Add(5)
	list.Add(4)
	list.Add(3)
	list.Add(2)
	list.Add(1)
	list.Add(0)
	list.Print()
	err := list.Insert(20, 6)
	if err != nil {
		fmt.Println(err)
	}
	list.Print()
	list.Remove(6)
	list.Print()
	list.Set(13, 5)
	list.Print()
}
