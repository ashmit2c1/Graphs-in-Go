package graph_introduction

type Graph struct {
	V    int
	List [][]int
}

func NewGraph(v int) *Graph {
	graph := &Graph{
		V: v,
		L: make([][]int, v),
	}
	return graph
}
