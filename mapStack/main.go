package main

import (
	"fmt"
)

var emptyStack = fmt.Errorf("Стэк пуст")

type mapStack struct {
	size int
	data map[int]int
}

// Создание структуры стэка
func newStack() mapStack {
	return mapStack{0, make(map[int]int)}
}

// Размер стэка
func (m mapStack) Size() int {
	return m.size
}

// Добавить элемент в стэк
func (m *mapStack) Push(el int) {
	m.size++
	m.data[m.Size()] = el
}

// Извлечь элемент из стека
func (m *mapStack) Pop() (int, error) {
	if m.Size() == 0 {
		return 0, emptyStack
	}
	temp := m.data[m.Size()]
	delete(m.data, m.Size())
	m.size--
	return temp, nil
}

// Выести содержимое стэка на экран
func (m mapStack) Print() {
	if m.Size() == 0 {
		fmt.Println(emptyStack)
	}
	fmt.Println("Содежжимое стэка:")
	for i := m.Size(); i >= 1; i-- {
		fmt.Println(m.data[i])

	}
}
func main() {
	stack := newStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.Print()
	fmt.Println("-----------------------------------------")
	el, err := stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(el)
	}
	el, err = stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(el)
	}
	el, err = stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(el)
	}
	el, err = stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(el)
	}
	el, err = stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(el)
	}
	el, err = stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(el)
	}
	el, err = stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(el)
	}

}
