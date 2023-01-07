package targetgraph

import "fmt"

func NewTargetGraph() *TargetGraph {
	return &TargetGraph{
		targets: make(map[string]*Target),
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

func NewStartNode() *Target {
	return &Target{
		Id:               START_TARGET_ID,
		Cwd:              "",
		Task:             START_TARGET_ID,
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

	// Ensure target graph has no cycles
	info := detectCyclesIn(t.targets)
	if info.hasCycle == true {
		return fmt.Errorf("Cycle detected: %v", info.path), nil
	}

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

	// TaskDependencies is a list of task dependencies like "^build", "build"
	TaskDependencies []string

	// Dependecies  are the targets that must be complete before the target can be complete
	Dependencies []string

	// Dependents are targets that depend on this target
	Dependents []string
}
