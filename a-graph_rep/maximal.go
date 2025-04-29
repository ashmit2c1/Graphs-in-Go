package graph_introduction

import "math"

func maximalNetworkRank(n int, roads [][]int) int {
	connected := make([][]bool, n)
	for i := 0; i < n; i++ {
		connected[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			connected[i][j] = false
		}
	}
	ans := math.MinInt32
	degree := make([]int, n)
	for i := 0; i < n; i++ {
		degree[i] = 0
	}
	for i := 0; i < len(roads); i++ {
		source := roads[i][0]
		destination := roads[i][1]
		degree[source]++
		degree[destination]++
		connected[source][destination] = true
		connected[destination][source] = true
	}
	for i := 0; i < len(degree); i++ {
		for j := i + 1; j < len(degree); j++ {
			rank := degree[i] + degree[j]
			if connected[i][j] == true {
				rank--
			}
			if rank > ans {
				ans = rank
			}
		}
	}
	return ans
}
