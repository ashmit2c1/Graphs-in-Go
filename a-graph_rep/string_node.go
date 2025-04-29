package graph_introduction

import "fmt"

type Node struct {
	Name      string
	Neighbors []string
}

type Graph2 struct {
	mp map[string]*Node
}

func NewGraph2(cities []string) *Graph2 {
	graph := Graph2{mp: make(map[string]*Node)}
	for i := 0; i < len(cities); i++ {
		city := cities[i]
		graph.mp[city] = &Node{Name: city}
	}
	return &graph
}

func (g *Graph2) AddEdge2(source string, destination string, undir bool) {
	g.mp[source].Neighbors = append(g.mp[source].Neighbors, destination)
	if undir == true {
		g.mp[destination].Neighbors = append(g.mp[destination].Neighbors, source)
	}
}

func (g *Graph2) PrintList() {
	for city, node := range g.mp {
		fmt.Printf("%s --> ", city)
		for _, neighbor := range node.Neighbors {
			fmt.Printf("%s  ", neighbor)
		}
		fmt.Println()
	}
}
