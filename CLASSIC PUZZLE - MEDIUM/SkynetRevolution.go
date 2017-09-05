package main

import "fmt"
import "os"

// описываем граф
type Graph map[int][]int


func appendLink(G Graph, fromKey int, toKey int) {

	G[fromKey] = append(G[fromKey], toKey)
	G[toKey] = append(G[toKey], fromKey)
	
}

func search(G Graph, node int) []int {

	queue := []int{node}
	marked := map[int]bool{node:true}
	
	visited :=  []int{}

	for len(queue) > 0 {

		r := queue[0]
		queue = queue[1:]

        visited = append(visited, r)


		for _, n := range G[r] {
			if !marked[n] {
				marked[n] = true
				queue = append(queue, n)
			}
		}
	}

    return visited
}



func main() {

    // N: the total number of nodes in the level, including the gateways
    // L: the number of links
    // E: the number of exit gateways
    var N, L, E int
    fmt.Scan(&N, &L, &E)
    
    G := make(Graph)
    
    for i := 0; i < L; i++ {
        // N1: N1 and N2 defines a link between these nodes
        var N1, N2 int
        fmt.Scan(&N1, &N2)
        
        appendLink(G, N1, N2)
        
    }
    for i := 0; i < E; i++ {
        // EI: the index of a gateway node
        var EI int
        fmt.Scan(&EI)
        fmt.Fprintln(os.Stderr, EI)
        
    }
    
        fmt.Fprintln(os.Stderr, "> ",G)
    
    for {
        // SI: The index of the node on which the Skynet agent is positioned this turn
        var SI int
        fmt.Scan(&SI)
        
        
        fmt.Fprintln(os.Stderr, "> ",SI)
        visited := search(G, SI)
        
        // Example: 0 1 are the indices of the nodes you wish to sever the link between
        fmt.Println(visited[0],visited[1])
    }
}
