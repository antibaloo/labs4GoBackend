package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// Метод всатвки нового элемента в бинарное дерево
func (t *TreeNode) Insert(value int) error {
	if t == nil {
		return errors.New("дерево пустое")
	}
	if t.val == value {
		return errors.New("узел с таким значение существует")
	}

	if t.val > value {
		if t.left == nil {
			t.left = &TreeNode{val: value}
			return nil
		}
		return t.left.Insert(value)
	}

	if t.val < value {
		if t.right == nil {
			t.right = &TreeNode{val: value}
			return nil
		}
		return t.right.Insert(value)
	}
	return nil
}

// Вывод на печать элементов дерева (обход в глубину)
func (t *TreeNode) printInDepth() {
	if t == nil {
		return
	}
	t.left.printInDepth()
	fmt.Printf("%v ", t.val)
	t.right.printInDepth()
}
func (t *TreeNode) printInWidth() {
	if t == nil {
		return
	}
	v := make([]*TreeNode, 0)
	v = append(v, t)
	level := 0
	for len(v) != 0 {
		vn := make([]*TreeNode, 0)
		fmt.Printf("уровень %v: ", level)
		for _, node := range v {

			fmt.Printf("%v ", node.val)
			if node.left != nil {
				vn = append(vn, node.left)
			}
			if node.right != nil {
				vn = append(vn, node.right)
			}
		}
		level++
		fmt.Printf("\n")
		v = vn
	}
}

// Поиск узла дерева по значению
func (t *TreeNode) Find(value int) (*TreeNode, error) {
	if t == nil {
		return nil, fmt.Errorf("узел с заданным значением: %v не найден", value)
	}
	switch {
	case value == t.val:
		return t, nil
	case value < t.val:
		return t.left.Find(value)
	default:
		return t.right.Find(value)
	}
}

// Удаление элемента по значению
func (t *TreeNode) Delete(value int) {
	t.remove(value)
}

// Вспомагательный рекурсивный метод для удаления элемента по значению
func (t *TreeNode) remove(value int) *TreeNode {

	if t == nil {
		return nil
	}

	if value < t.val {
		t.left = t.left.remove(value)
		return t
	}
	if value > t.val {
		t.right = t.right.remove(value)
		return t
	}
	//если у удаляемого элемента нет наследников, просто обнуляем указательна него
	if t.left == nil && t.right == nil {
		t = nil
		return nil
	}
	//если у удаляемого элемента нет левого наследника, то присваемаем указателю на удаляемы элемент указатель на правого наследника
	if t.left == nil {
		t = t.right
		return t
	}
	//если у удаляемого элемента нет правого наследника, то присваемаем указателю на удаляемы элемент указатель на левого наследника
	if t.right == nil {
		t = t.left
		return t
	}

	smallestValOnRight := t.right
	for {
		//поиск минимального по значению элемента в правой ветке
		if smallestValOnRight != nil && smallestValOnRight.left != nil {
			smallestValOnRight = smallestValOnRight.left
		} else {
			break
		}
	}
	t.val = smallestValOnRight.val  //значение минимума в val удаляемого элемента
	t.right = t.right.remove(t.val) //сам минимальный элемент удаляем.
	return t
}

// Генерация среза заданной длины и диапазова значений
func generateSlice(max, size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = rand.Intn(max*2) - max
	}

	return ar
}

func main() {
	const max = 100
	sourceSlice := generateSlice(max, 10)
	tree := &TreeNode{val: sourceSlice[0]}
	fmt.Println(sourceSlice)
	for i, el := range sourceSlice {
		if i != 0 {
			tree.Insert(el)
		}
	}
	tree.printInDepth()
	fmt.Printf("\n")
	tree.printInWidth()
	node, err := tree.Find(max + 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(node.val)
	}
	node, err = tree.Find(sourceSlice[rand.Intn(len(sourceSlice)-1)])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("найден узел со значением %v \n", node.val)
	}
	toDelete := sourceSlice[rand.Intn(len(sourceSlice)-1)]
	fmt.Println("Удаляем элемент со значением", toDelete)
	tree.Delete(toDelete)
	tree.printInDepth()
	fmt.Printf("\n")
	tree.printInWidth()
}
