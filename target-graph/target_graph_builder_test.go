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

	graph := targetGraph.build()
	assert.Equal(t, graph, "")
}
