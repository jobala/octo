package targetgraph

import (
	"fmt"
	"path/filepath"

	"github.com/jobala/octo/workspace-tools"
)

type WorkspaceTargetGraph struct {
	dependencyMap workspace.DependencyMap
	targetGraph   *TargetGraph
	targetFactory TargetFactory
}

func NewWorkspaceTargetGraph(root string, pkgInfos workspace.PackageInfos) *WorkspaceTargetGraph {
	return &WorkspaceTargetGraph{
		dependencyMap: createDependencyMap(root, pkgInfos),
		targetGraph:   NewTargetGraph(),
		targetFactory: NewTargetFactory(root, func(pkgName string) string {
			return filepath.Dir(fmt.Sprint("%", pkgInfos[pkgName].PackageJsonPath))
		}),
	}
}
