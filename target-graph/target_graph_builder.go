package targetgraph

import (
	"errors"
	"fmt"
)

func NewTargetGraph() *TargetGraph {
	return &TargetGraph{
		targets: map[string]*Target{},
	}
}

func NewTarget(pkg, task string) *Target {
	return &Target{
		Id:               createTargetId(pkg, task),
		Cwd:              "",
		Task:             task,
		TaskDependencies: []string{},
		Dependencies:     []string{},
		Dependents:       []string{},
	}
}

func (t *TargetGraph) addTarget(target *Target) {
	t.targets[target.Id] = target
}

func (t *TargetGraph) addDependency(dependency, dependent string) {
	parent := t.targets[dependent]
	child := t.targets[dependency]

	parent.Dependencies = append(parent.Dependencies, child.Id)
	child.Dependents = append(child.Dependents, parent.Id)
}

func (t *TargetGraph) build() (error, map[string]*Target) {

	// ensure target graph has no cycles
	info := detectCyclesIn(t.targets)
	if info.hasCycle == true {
		return fmt.Errorf("Cycle detected: %v", info.path), nil
	}

	// prioritize(t.targets)

	return nil, t.targets
}

func (t *TargetGraph) subgraph() {

}

type TargetGraph struct {
	targets map[string]*Target
}

type Target struct {
	// Id is the target's identifier of the form package#task
	Id   string
	Cwd  string
	Task string

	// taskDeps is a list of task dependencies like "^build", "build"
	TaskDependencies []string

	// dependecies  are the targets that must be complete before the target can be complete
	Dependencies []string

	// dependents are targets that depend on this target
	Dependents []string
}
