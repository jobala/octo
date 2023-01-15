package targetgraph

import "fmt"

func NewTargetGraph() *TargetGraph {
	return &TargetGraph{
		targets: map[string]*Target{
			START_TARGET_ID: {
				Id: START_TARGET_ID,
			},
		},
	}
}

func NewTarget(pkg, task string) *Target {
	return &Target{
		Id:               createTargetId(pkg, task),
		Cwd:              "",
		Task:             task,
		Type:             "",
		TaskDependencies: []string{},
		Dependencies:     []string{},
		Dependents:       []string{},
		Inputs:           []string{},
		Outputs:          []string{},
		Cache:            false,
	}
}

// addTarget adds a node to the target graph
func (t *TargetGraph) addTarget(target *Target) {
	t.targets[target.Id] = target
	t.addDependency(target.Id, START_TARGET_ID)
}

// addDependency creates an edge between two  nodes in the target graph
func (t *TargetGraph) addDependency(child, parent string) {
	parentNode := t.targets[parent]
	childNode := t.targets[child]

	parentNode.Dependencies = append(parentNode.Dependencies, childNode.Id)
	childNode.Dependents = append(childNode.Dependents, parentNode.Id)
}

func (t *TargetGraph) build() (error, map[string]*Target) {

	// Ensure target graph has no cycles
	info := detectCyclesIn(t.targets)
	if info.hasCycle == true {
		return fmt.Errorf("Cycle detected: %v", info.path), nil
	}

	return nil, t.targets
}

// subgraph creates a smaller target graph from the passed ids
func (t *TargetGraph) subgraph(ids []string) (error, *TargetGraph) {
	subGraph := NewTargetGraph()

	for _, targetId := range ids {
		if _, presentInSubgraph := subGraph.targets[targetId]; !presentInSubgraph {

			// Create a copy of a target to avoid unintentional modification of targets in the main graph
			target := *t.targets[targetId]

			target.Dependencies = make([]string, 0)
			target.Dependents = make([]string, 0)

			subGraph.addTarget(&target)
		}
	}

	for _, targetId := range ids {
		t.populateSubgraph(subGraph, targetId, []string{})
	}

	return nil, subGraph
}

// populateSubgraph recursively adds dependencies for subgraph nodes
func (t *TargetGraph) populateSubgraph(subGraph *TargetGraph, targetId string, path []string) {
	if contains(path, targetId) {
		return
	}

	path = append(path, targetId)

	for _, neighbour := range t.targets[targetId].Dependencies {
		if _, presentInSubgraph := subGraph.targets[neighbour]; !presentInSubgraph {
			// Create a copy of a target to avoid unintentional modification of targets in the main graph
			target := *t.targets[neighbour]

			target.Dependencies = []string{}
			target.Dependents = []string{}

			subGraph.addTarget(&target)
		}
		subGraph.addDependency(neighbour, targetId)

		t.populateSubgraph(subGraph, neighbour, path)
	}
}

type TargetGraph struct {
	targets map[string]*Target
}

type Target struct {
	// Id is the target's identifier of the form package#task
	Id      string
	Cwd     string
	Task    string
	Type    string
	PkgName string

	// TaskDependencies is a list of task dependencies like "^build", "build"
	TaskDependencies []string

	// Dependecies  are the targets that must be complete before the target can be complete
	Dependencies []string

	// Dependents are targets that depend on this target
	Dependents []string

	Inputs []string

	Outputs []string

	Cache bool
}
