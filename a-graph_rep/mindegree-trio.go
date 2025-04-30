package graph_introduction

import "math"

func minDegreeTrio(n int, edges [][]int) int {
	ans := math.MaxInt32
	degree := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		degree[i] = 0
	}
	adj := make([]map[int]bool, n+1)
	for i := 0; i <= n; i++ {
		adj[i] = make(map[int]bool)
	}
	for i := 0; i < len(edges); i++ {
		source := edges[i][0]
		destination := edges[i][1]
		degree[source]++
		degree[destination]++
		adj[source][destination] = true
		adj[destination][source] = true
	}
	for i := 1; i <= n; i++ {
		for j := range adj[i] {
			if j <= i {
				continue
			}
			for k := range adj[j] {
				if k <= j {
					continue
				}
				if adj[i][k] == true && adj[j][k] == true {
					deg := degree[i] + degree[j] + degree[k] - 6
					if deg < ans {
						ans = deg
					}
				}
			}
		}
	}
	if ans == math.MaxInt32 {
		return -1
	} else {
		return ans
	}
}
