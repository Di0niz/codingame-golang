package main

import "fmt"

// описываем граф
type Graph map[int][]int


func appendLink(G Graph, fromKey int, toKey int) {

	G[fromKey] = append(G[fromKey], toKey)
	G[toKey] = append(G[toKey], fromKey)
	
}

func search(G Graph, node int) {

	queue := []int{node}
	marked := map[int]bool{node:true}

	for len(queue) > 0 {

		r := queue[0]
		queue = queue[1:]

		fmt.Println(r)

		for _, n := range G[r] {
			if !marked[n] {
				marked[n] = true
				queue = append(queue, n)
			}
		}
	}


}

func main() {
	
	G := make(Graph)


	appendLink(G, 1, 2)
	appendLink(G, 3, 1)
	appendLink(G, 3, 4)
	appendLink(G, 5, 4)
	appendLink(G, 5, 9)
	appendLink(G, 7, 9)
	
	fmt.Println(G)

	search(G, 5)
	

}
