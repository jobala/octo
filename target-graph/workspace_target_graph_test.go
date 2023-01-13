package targetgraph

import (
	"fmt"
	"testing"

	workspacetools "github.com/jobala/octo/workspace-tools"
	"github.com/stretchr/testify/assert"
)

func TestWorkspaceTargetGraph_BuildFromPackageAndTaskGraphs(t *testing.T) {
	root := "repos/a"
	packageInfos := createPackageInfo(map[string][]string{
		"a": {"b"},
		"b": {},
	})

	workspaceGraph := NewWorkspaceTargetGraph(root, packageInfos)
	workspaceGraph.addTargetConfig("build", map[string][]string{
		"dependsOn": {"^build"},
	})

	targetGraph := workspaceGraph.build([]string{"build"})

	assert.Contains(t, targetGraph["__start"].dependencies, "a#build")
	assert.Contains(t, targetGraph["__start"].dependencies, "b#build")
	assert.Contains(t, targetGraph["a#build"].dependencies, "b#build")
	assert.Equal(t, len(targetGraph["b#build"].dependencies), 0)
}

func createPackageInfo(packages map[string][]string) workspacetools.PackageInfos {
	var res workspacetools.PackageInfos

	for pkgId, dependencies := range packages {
		res[pkgId] = workspacetools.PackageInfo{
			PackageJsonPath: fmt.Sprintf("/path/to/%s/package.json", pkgId),
			Name:            pkgId,
			Version:         "1.0.0",
			Dependencies:    dependencies,
		}
	}

	return res
}
