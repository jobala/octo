package targetgraph

import (
	"fmt"
	"testing"

	"github.com/jobala/octo/workspace-tools"
	workspacetools "github.com/jobala/octo/workspace-tools"
)

func TestWorkspaceTargetGraph_BuildFromPackageAndTaskGraphs(t *testing.T) {
	root := "repos/a"
	packageInfos := createPackageInfo(map[string]workspace.Dependency{
		"a": {"b": "10"},
		"b": {},
	})

	workspaceGraph := NewWorkspaceTargetGraph(root, packageInfos)
	workspaceGraph.AddTargetConfig("build", TargetConfig{DependsOn: []string{"^build"}})

	workspaceGraph.Build([]string{"build"}, nil)

	// assert.Contains(t, targetGraph["__start"].dependencies, "a#build")
	// assert.Contains(t, targetGraph["__start"].dependencies, "b#build")
	// assert.Contains(t, targetGraph["a#build"].dependencies, "b#build")
	// assert.Equal(t, len(targetGraph["b#build"].dependencies), 0)
}

func createPackageInfo(packages map[string]workspacetools.Dependency) workspacetools.PackageInfos {
	res := make(workspacetools.PackageInfos)

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
