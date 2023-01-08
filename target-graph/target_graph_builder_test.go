package targetgraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTargetGraph_BuildsAFullGraph(t *testing.T) {
	targetA := NewTarget("a", "build")
	targetB := NewTarget("b", "build")

	targetGraph := NewTargetGraph()
	targetGraph.addTarget(targetA)
	targetGraph.addTarget(targetB)

	// targetA is a dependency of targetB
	targetGraph.addDependency(targetA.Id, targetB.Id)

	err, graph := targetGraph.build()
	assert.NoError(t, err)

	assert.Contains(t, graph["b#build"].Dependencies, "a#build")
	assert.Contains(t, graph["a#build"].Dependents, "b#build")
}

func TestTargetGraph_BuildsASubgraph(t *testing.T) {
	targetA := NewTarget("a", "build")
	targetB := NewTarget("b", "build")
	targetC := NewTarget("c", "build")
	targetD := NewTarget("d", "build")

	targetGraph := NewTargetGraph()
	targetGraph.addTarget(targetA)
	targetGraph.addTarget(targetB)
	targetGraph.addTarget(targetC)
	targetGraph.addTarget(targetD)

	targetGraph.addDependency(targetA.Id, targetB.Id)
	targetGraph.addDependency(targetA.Id, targetC.Id)
	targetGraph.addDependency(targetD.Id, targetA.Id)

	_, graph := targetGraph.subgraph([]string{"a#build"})

	assert.Contains(t, graph.targets[START_TARGET_ID].Dependencies, "a#build")
	assert.Contains(t, graph.targets["a#build"].Dependencies, "d#build")
}
