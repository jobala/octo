package targetgraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCycles_HasCycleIsTrue(t *testing.T) {
	cyclicGraph := generateGraphFrom(cyclicEdges)
	cycleInfo := detectCyclesIn(cyclicGraph)

	assert.Equal(t, cycleInfo.hasCycle, true)
	assert.Equal(t, cycleInfo.path, []string{"a#build", "b#build", "c#build", "d#build", "a#build"})
}

func TestCycles_HasCycleIsFalse(t *testing.T) {
	acyclicGraph := generateGraphFrom(acyclicEdges)
	cycleInfo := detectCyclesIn(acyclicGraph)

	assert.Equal(t, cycleInfo.hasCycle, false)
	assert.Equal(t, len(cycleInfo.path), 0)
}

func generateGraphFrom(edges [][]string) map[string]*Target {
	targets := make(map[string]*Target)

	for _, edge := range edges {
		from, to := edge[0], edge[1]

		if _, ok := targets[from]; !ok {
			_, pkg, task := getPackageAndTask(from)
			targets[from] = NewTarget(pkg, task)
		}

		if _, ok := targets[to]; !ok {
			_, pkg, task := getPackageAndTask(to)
			targets[to] = NewTarget(pkg, task)
		}

		targets[from].Dependencies = append(targets[from].Dependencies, to)
		targets[to].Dependents = append(targets[to].Dependents, from)
	}

	return targets
}

var cyclicEdges = [][]string{
	{"a#build", "b#build"},
	{"b#build", "c#build"},
	{"c#build", "d#build"},
	{"d#build", "a#build"},
	{"__start", "a#build"},
	{"__start", "b#build"},
	{"__start", "c#build"},
	{"__start", "d#build"},
}

var acyclicEdges = [][]string{
	{"a#build", "b#build"},
	{"b#build", "c#build"},
	{"c#build", "d#build"},
	{"__start", "a#build"},
	{"__start", "b#build"},
	{"__start", "c#build"},
	{"__start", "d#build"},
}
