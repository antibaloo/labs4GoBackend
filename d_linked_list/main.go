package main

import "fmt"

//Ошибка при передаче несуществующего в списке индекса
var ErrorWrongListIndex = fmt.Errorf("Неверный индекс списка")

//Описание типа - узел списка
type IntNode struct {
	Value int
	Prev  *IntNode
	Next  *IntNode
}

//Описание метода создания узла списка
func NewIntNode(value int) *IntNode {
	return &IntNode{value, nil, nil}
}

//Описание типа - двусвязный список целых чисел
type DoubleLinkedList struct {
	size int
	Head *IntNode
	Tail *IntNode
}

//Описание метода создания списка целых чисел
func NewDoubleLInkedList() *DoubleLinkedList {
	return &DoubleLinkedList{0, nil, nil}
}

//Опсиание метода возвращающего размер  списка
func (d *DoubleLinkedList) Size() int {
	return d.size
}

//Описание метода получения элемента списка по индексу
func (d *DoubleLinkedList) Get(index int) (*IntNode, error) {
	if index < 0 || index >= d.Size() {
		return nil, ErrorWrongListIndex
	}
	//Если индекс строго меньше результата целочисленного деления d.Size()/2, то ищем с головы списка, иначе с хвоста
	if index < d.Size()/2 {
		node := d.Head
		for i := 1; i <= index; i++ {
			node = node.Next
		}
		return node, nil
	} else {
		node := d.Tail
		for i := d.Size() - 2; i >= index; i-- {
			node = node.Prev
		}
		return node, nil
	}
}

//Описание метода присвоения значения элементу списка
func (d *DoubleLinkedList) Set(el int, index int) error {
	if index < 0 || index >= d.Size() {
		return ErrorWrongListIndex
	}
	node, err := d.Get(index)
	if err != nil {
		return err
	}
	node.Value = el
	return nil
}

//Описание метода добавления нового элемента в начало списка
func (d *DoubleLinkedList) Add(el int) {
	newNode := NewIntNode(el)
	if d.Head == nil {
		d.Head = newNode
		d.Tail = newNode
	} else {
		d.Head.Prev = newNode
		newNode.Next = d.Head
		d.Head = newNode
	}
	d.size++
}

//Описание метода добавления нового элемента в конец списка
func (d *DoubleLinkedList) App(el int) {
	newNode := NewIntNode(el)
	if d.Head == nil {
		d.Head = newNode
		d.Tail = newNode
	} else {
		d.Tail.Next = newNode
		newNode.Prev = d.Tail
		d.Tail = newNode
	}
	d.size++
}

//Описание метода вставки нового элемента в произвольную позиция
func (d *DoubleLinkedList) Insert(el int, index int) error {
	switch {
	case index < 0 || index > d.Size():
		return ErrorWrongListIndex
	case index == 0:
		d.Add(el)
		return nil
	case index == d.Size():
		d.App(el)
		return nil
	default:
		newNode := NewIntNode(el)
		prevNode, err := d.Get(index - 1)
		if err != nil {
			return err
		}
		nexNode, err := d.Get(index)
		if err != nil {
			return err
		}
		prevNode.Next = newNode
		newNode.Prev = prevNode
		newNode.Next = nexNode
		nexNode.Prev = newNode
		d.size++
		return nil
	}
}

//Описание метода удаления элемента из произвольной позиции
func (d *DoubleLinkedList) Remove(index int) error {
	switch {
	case index < 0 || index > d.Size():
		return ErrorWrongListIndex
	case index == 0:
		nextNode, err := d.Get(index + 1)
		if err != nil {
			return err
		}
		d.Head = nextNode
		nextNode.Prev = nil
		d.size--
	case index == d.Size()-1:
		prevNode, err := d.Get(index - 1)
		if err != nil {
			return err
		}
		d.Tail = prevNode
		prevNode.Next = nil
		d.size--
	default:
		prevNode, err := d.Get(index - 1)
		if err != nil {
			return err
		}
		nextNode, err := d.Get(index + 1)
		if err != nil {
			return err
		}
		prevNode.Next = nextNode
		nextNode.Prev = prevNode
		d.size--
	}
	return nil
}

//Описание метода вывода на печать двусвязного списка челых чисел
func (d DoubleLinkedList) Print() {
	node := d.Head
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
	dlist := NewDoubleLInkedList()
	dlist.Print()
	dlist.Add(0)
	dlist.App(1)
	dlist.Print()
	dlist.Insert(33, 2)
	dlist.Print()
	dlist.Insert(44, 2)
	dlist.Print()
	dlist.Remove(3)
	dlist.Print()
	fmt.Println(dlist.Size())
}
