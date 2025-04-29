package graph_introduction

import (
	"fmt"
)

// graph struct
type Graph struct {
	V    int
	List [][]int
}

// construction ( kinda of ) to create the graph
func NewGraph(v int) *Graph {
	graph := &Graph{
		V:    v,
		List: make([][]int, v),
	}
	return graph
}

// adding edge between nodes in the graph
func (g *Graph) addEdge(source int, destination int, undir bool) {
	g.List[source] = append(g.List[source], destination)
	if undir == true {
		g.List[destination] = append(g.List[destination], source)
	}
}

// printing the list
func (g *Graph) printList() {
	for i := 0; i < len(g.List); i++ {
		fmt.Print(g.List[i])
		fmt.Print("-->")
		for j := 0; j < len(g.List[i]); j++ {
			fmt.Print(g.List[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

// creating adjacency list in the graph
func createList() [][]int {
	var n int
	var m int
	var undir bool
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&undir)
	adj := make([][]int, n)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		adj[u] = append(adj[u], v)
		if undir {
			adj[v] = append(adj[v], u)
		}
	}
	return adj
}

// creating adj matrix
func createMatix() [][]bool {
	var n int
	var m int
	var undir bool
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&undir)
	adj := make([][]bool, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			adj[i][j] = false
		}
	}
	for i := 0; i < m; i++ {
		var u int
		var v int
		fmt.Scan(&u, &v)
		adj[u][v] = true
		if undir == true {
			adj[v][u] = true
		}
	}
	return adj
}
