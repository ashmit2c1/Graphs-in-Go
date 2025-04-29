package graph_introduction

// find the centre of a star graph
func findCentre(edges [][]int) int {
	n := len(edges)
	indegree := make([]int, n)
	outdegree := make([]int, n)
	for i := 0; i < n; i++ {
		indegree[i] = 0
		outdegree[i] = 0
	}
	for i := 0; i < len(edges); i++ {
		source := edges[i][0]
		destination := edges[i][1]
		indegree[source]++
		indegree[destination]++
		outdegree[source]++
		outdegree[destination]++
	}
	for i := 0; i < n; i++ {
		if indegree[i] == n-1 && outdegree[i] == n-1 {
			return i
		}
	}
	return -1
}
