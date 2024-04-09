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

// Структура ориентированного графа
type dGraph struct {
	size int
	data map[int]map[int]int
}

// Создание графа заданного размера s
func NewUGraph(s int) (*dGraph, error) {
	if s <= 0 {
		return &dGraph{}, errors.New("нельзя создать граф размером меньше 1")
	}
	d := make(map[int]map[int]int)
	for i := 1; i <= s; i++ {
		d[i] = make(map[int]int)
		for j := 1; j <= s; j++ {
			d[i][j] = 0
		}
	}
	return &dGraph{
		size: s,
		data: d,
	}, nil
}

// Добавление ребра между вершинами i и j, с проверкой существования вершин
func (d *dGraph) addEdge(i, j int) error {
	if i < 1 && i > d.size || j < 1 && j > d.size {
		return errors.New("вершин с таким индексом нет в графе")
	}
	//Рембра несимметричны и имеют четкое направление
	d.data[i][j] = 1
	return nil
}

// Вывод графа в виде матрицы связности
func (d *dGraph) Matrix() {
	for i := 1; i <= d.size; i++ {
		for j := 1; j <= d.size; j++ {
			fmt.Printf("%v ", d.data[i][j])
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

// Обход ориентированного графав ширину (Breadth-First Search) с восстановлением пути, возвращает строку
// последовательности вершин типа start- ... -n[i]- ... goal, или сообщение об ошибке
func (d *dGraph) sRoute(start, goal int) string {
	route := make(map[int]int)
	routeString := fmt.Sprint(goal)
	if start < 1 || start > d.size {
		return "стратовой вершины нет в графе"
	}
	if goal < 1 || goal > d.size {
		return "финишной вершины нет в графе"
	}
	var q queue
	visitCheck := make(map[int]string)
	q.Push(start)     //Ставим начальную вершину в очередь
	for q.Len() > 0 { //Пока очередь вершин не пуста
		vertex, _ := q.Pop() //Берем вершину из очереди
		if vertex == goal {  //если извлеченная вершина равно целевой, восстанавливаем маршрут и возвращаем результат
			last := goal
			for route[last] != start {
				step := route[last]
				routeString = fmt.Sprint(step, "-") + routeString
				last = step
			}
			routeString = fmt.Sprint(route[last], "-") + routeString
			return routeString
		}
		visitCheck[vertex] = ""        //Помечаем, что посетили
		for i := 1; i <= d.size; i++ { //Добавляем смежные вершины в очередь
			if d.data[vertex][i] == 1 {
				if _, vizited := visitCheck[i]; !vizited { //Если не посещали ее, помечаем
					route[i] = vertex // фиксируем шаг маршрута для последующего восстановления
					q.Push(i)
					visitCheck[i] = "" //Помечаем, что посетили
				}

			}
		}
	}
	return "вершина недостижима"
}

// Генерация случайных ребер в графе
func (d *dGraph) randomEdge() {
	for i := 1; i <= d.size; i++ {
		for j := 1; j <= d.size; j++ {
			if rand.Intn(2) == 1 && i != j { //исключаем ребра с самим собой
				d.addEdge(i, j)
			}
		}
	}
}

func main() {
	const size = 6
	dG, _ := NewUGraph(size)
	dG.randomEdge()
	dG.Matrix()
	fmt.Print(dG.sRoute(1, 6))
}
