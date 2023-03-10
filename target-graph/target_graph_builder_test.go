package targetgraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTargetGraph_BuildFullGraph(t *testing.T) {
	targetA := NewTarget("a", "build")
	targetB := NewTarget("b", "build")

	targetGraph := NewTargetGraph()
	targetGraph.addTarget(targetA)
	targetGraph.addTarget(targetB)

	// targetA is a dependency of targetB
	targetGraph.addDependency(targetB.Id, targetA.Id)

	err, graph := targetGraph.build()
	assert.NoError(t, err)

	assert.Contains(t, graph["b#build"].Dependencies, "a#build")
	assert.Contains(t, graph["a#build"].Dependents, "b#build")
}

func TestTargetGraph_BuildSubgraph(t *testing.T) {
	targetA := NewTarget("a", "build")
	targetB := NewTarget("b", "build")
	targetC := NewTarget("c", "build")
	targetD := NewTarget("d", "build")

	targetGraph := NewTargetGraph()
	targetGraph.addTarget(targetA)
	targetGraph.addTarget(targetB)
	targetGraph.addTarget(targetC)
	targetGraph.addTarget(targetD)

	targetGraph.addDependency(targetB.Id, targetA.Id)
	targetGraph.addDependency(targetC.Id, targetA.Id)
	targetGraph.addDependency(targetA.Id, targetD.Id)

	_, graph := targetGraph.subgraph([]string{"a#build"})

	assert.Contains(t, graph.targets[START_TARGET_ID].Dependencies, "a#build")
	assert.Contains(t, graph.targets["a#build"].Dependencies, "d#build")
	assert.Equal(t, len(graph.targets["d#build"].Dependencies), 0)
}
