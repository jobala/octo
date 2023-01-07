package targetgraph

func detectCyclesIn(graph map[string]*Target) cycleInfo {
	isCyclic, path := dfs(graph, START_TARGET_ID, []string{})
	return cycleInfo{
		hasCycle: isCyclic,
		path:     path,
	}
}

func dfs(graph map[string]*Target, currNode string, visited []string) (bool, []string) {

	// If we have visited this node before then we are in a cyclic graph
	if contains(visited, currNode) {
		visited = append(visited, currNode)
		return true, visited
	}

	// avoid adding the starting node to the visited nodes path. The starting node is a placeholder node
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
	for _, k := range items {
		if k == item {
			return true
		}
	}

	return false
}

type cycleInfo struct {
	path     []string
	hasCycle bool
}
