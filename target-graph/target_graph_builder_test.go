package targetgraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTargetGraph_BuildsAFullGraph(t *testing.T) {
	startNode := NewStartNode()
	targetA := NewTarget("a", "build")
	targetB := NewTarget("b", "build")

	targetGraph := NewTargetGraph()
	targetGraph.addTarget(startNode)
	targetGraph.addTarget(targetA)
	targetGraph.addTarget(targetB)

	// targetA is a dependency of targetB
	targetGraph.addDependency(targetA.Id, targetB.Id)
	targetGraph.addDependency(targetA.Id, startNode.Id)
	targetGraph.addDependency(targetB.Id, startNode.Id)

	err, graph := targetGraph.build()
	assert.NoError(t, err)

	assert.Contains(t, graph["b#build"].Dependencies, "a#build")
	assert.Contains(t, graph["a#build"].Dependents, "b#build")
}
