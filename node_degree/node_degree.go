package node_degree

import (
	"errors"
	"strconv"
)

// Degree func
func Degree(nodes int, graph [][2]int, node int) (int, error) {
	if node > nodes {
		return 0, errors.New("node " + strconv.Itoa(node) + " not found in the graph")
	}

	degree := 0

	for i := 0; i < len(graph); i++ {
		if graph[i][0] == node || graph[i][1] == node {
			degree++
		}
	}

	return degree, nil
}
