package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	SIZE_MATRIX = 100
)

// Graph описываем граф
type Graph [SIZE_MATRIX][SIZE_MATRIX]int

// NewGraph определяет новый граф в виде разряженного массива map
func NewGraph() *Graph {
	g := &Graph{}
	return g
}

// AppendLink добавляет в Graph новые индексы
func (g *Graph) AppendLink(fromKey int, toKey int) {

	(*g)[fromKey][toKey] = 1
}

// Neighbors формируем список соседей для элемента
func (g *Graph) Neighbors(node int) <-chan int {
	ch := make(chan int)
	go func(node int) {
		nodes := (*g)[node]
		for val, _ := range nodes {
			if nodes[val] > 0 {
				ch <- val
			}
		}
		close(ch)
	}(node)

	return ch
}

// ToString возвращает описание графа
func (g *Graph) String() string {
	// создаем массив размеров с граф
	buf := ""
	for i, _ := range *g {

		for j, _ := range (*g)[i] {
			if (*g)[i][j] > 0 {
				buf += fmt.Sprintf("%d -> %d\n", i, j)
			}
		}

	}

	return buf
}

// Search ищет кратчайший путь до элемента графа
// возвращает массив элементов
func (g *Graph) Search(from int, goal int) ([]int, error) {

	queue := []int{from}

	// список пройденных вершин
	visited := make([]int, SIZE_MATRIX)

	// список пройденных вершин
	// определяем расстояние до вершины
	marked := make([]int, SIZE_MATRIX)

	findPath := false

	dist := 0
	for len(queue) > 0 {
		r := queue[0]

		queue = queue[1:]

		dist = marked[r]

		visited[dist] = r

		if r == goal {
			findPath = true
			break
		}

		for neighbor := range g.Neighbors(r) {

			if marked[neighbor] == 0 {
				marked[neighbor] = dist + 1
				queue = append(queue, neighbor)
			}

		}

	}

	var err error

	if !findPath {
		err = errors.New("Path not found")
	}

	return visited[:dist+1], err

}

func main() {

	// N: the total number of nodes in the level, including the gateways
	// L: the number of links
	// E: the number of exit gateways
	var N, L, E int
	fmt.Scan(&N, &L, &E)

	G := NewGraph()

	Exits := make([]int, E)

	for i := 0; i < L; i++ {
		// N1: N1 and N2 defines a link between these nodes
		var N1, N2 int
		fmt.Scan(&N1, &N2)

		G.AppendLink(N1, N2)
		G.AppendLink(N2, N1)

	}
	for i := 0; i < E; i++ {
		// EI: the index of a gateway node
		var EI int
		fmt.Scan(&EI)
		Exits[i] = EI
		fmt.Fprintln(os.Stderr, EI)

	}

	for {
		// SI: The index of the node on which the Skynet agent is positioned this turn
		var SI int
		fmt.Scan(&SI)

		fmt.Fprintln(os.Stderr, "SI: ", SI)

		nearExits := 100
		nextGateway := 0
		for _, exit := range Exits {

			visited, _ := G.Search(SI, exit)

			if len(visited) < nearExits {
				nextGateway = visited[1]
				nearExits = len(visited)

			}

			fmt.Fprintln(os.Stderr, "visited: ", visited)
		}

		// Example: 0 1 are the indices of the nodes you wish to sever the link between
		fmt.Println(SI, nextGateway)
	}
}
