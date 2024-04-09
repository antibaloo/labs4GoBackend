package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// Структура очереди
type queue struct {
	data []int
}

// Метод возвращающий длину очереди
func (q *queue) Len() int {
	return len(q.data)
}

// Добавление элемента в очередь
func (q *queue) Push(i int) {
	q.data = append(q.data, i)
}

// Извлечение элемента из очереди
func (q *queue) Pop() (int, error) {
	if len(q.data) > 0 {
		res := q.data[0]
		q.data = q.data[1:]
		return res, nil
	}
	return 0, errors.New("очередь пуста")
}

// Структура неориентированного графа
type uGraph struct {
	size int
	data map[int]map[int]int
}

// Создание графа заданного размера s
func NewUGraph(s int) (*uGraph, error) {
	if s <= 0 {
		return &uGraph{}, errors.New("нельзя создать граф размером меньше 1")
	}
	d := make(map[int]map[int]int)
	for i := 1; i <= s; i++ {
		d[i] = make(map[int]int)
		for j := 1; j <= s; j++ {
			d[i][j] = 0
		}
	}
	return &uGraph{
		size: s,
		data: d,
	}, nil
}

// Добавление ребра между вершинами i и j, с проверкой существования вершин
func (u *uGraph) addEdge(i, j int) error {
	if i < 1 && i > u.size || j < 1 && j > u.size {
		return errors.New("вершин с таким индексом нет в графе")
	}
	//в неориентурованном графе ребра симметричны, по ним можно пройти в обе стороны
	u.data[i][j] = 1
	u.data[j][i] = 1
	return nil
}

// Вывод графа в виде матрицы связности
func (u *uGraph) Matrix() {
	for i := 1; i <= u.size; i++ {
		for j := 1; j <= u.size; j++ {
			fmt.Printf("%v ", u.data[i][j])
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

// Обход неориентированного графав ширину (Breadth-First Search)
func (u *uGraph) BFS(v int) {
	if v < 1 || v > u.size {
		fmt.Println("такой вершины нет в графе")
		return
	}
	var q queue
	visitCheck := make(map[int]string)
	q.Push(v)         //Ставим начальную вершину в очередь
	for q.Len() > 0 { //Пока очередь вершин не пуста
		vertex, _ := q.Pop()                           //Извлекаем вершину
		if _, vizited := visitCheck[vertex]; vizited { //Если уже посещали ее, извлекаем следующую
			continue
		}
		visitCheck[vertex] = ""        //Помечаем, что посетили
		fmt.Printf("%v ", vertex)      //Выводим посещенную вершину
		for i := 1; i <= u.size; i++ { //Добавляем смежные вершины в очередь
			if u.data[vertex][i] == 1 {
				q.Push(i)
			}
		}
	}
	fmt.Print("\n")
}

// Генерация случайных ребер в графе
func (u *uGraph) randomEdge() {
	for i := 1; i <= u.size; i++ {
		for j := 1; j <= u.size; j++ {
			if rand.Intn(2) == 1 && i != j { //исключаем ребра с самим собой
				u.addEdge(i, j)
			}
		}
	}
}

func main() {
	const size = 6
	uG, _ := NewUGraph(size) //Не проверяем ошибку, т.к. уверены в передаваемых параметрах
	uG.randomEdge()
	uG.Matrix()
	uG.BFS(1)
}
