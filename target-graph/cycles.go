package targetgraph

func detectCyclesIn(graph map[string]*Target) cycleInfo {
	isCyclic, path := dfs(graph, START_TARGET_ID, []string{})
	return cycleInfo{
		hasCycle: isCyclic,
		path:     path,
	}
}

// dfs traverses the graph in a depth first manner tracking the nodes visited
// and checking if the graph is cyclic
func dfs(graph map[string]*Target, currNode string, visited []string) (bool, []string) {

	// If the currNode had been visited before then the graph is cyclic
	if contains(visited, currNode) {
		visited = append(visited, currNode)
		return true, visited
	}

	// Ignore the START_TARGET_ID from the visited path because it is a placeholder node
	if currNode != START_TARGET_ID {
		visited = append(visited, currNode)
	}

	for _, neighbour := range graph[currNode].Dependencies {
		return dfs(graph, neighbour, visited)
	}
	visited = pop(visited)

	return false, nil
}

func pop[T any](items []T) []T {
	return items[:len(items)-1]
}

func contains(items []string, item string) bool {
	for _, element := range items {
		if element == item {
			return true
		}
	}

	return false
}

type cycleInfo struct {
	path     []string
	hasCycle bool
}
