package targetgraph

import (
	"fmt"
	"path/filepath"

	"github.com/jobala/octo/workspace-tools"
)

type WorkspaceTargetGraph struct {
	DependencyMap *workspace.DependencyMap
	Graph         *TargetGraph
	Factory       TargetFactory
	PkgInfos      workspace.PackageInfos
}

func NewWorkspaceTargetGraph(root string, pkgInfos workspace.PackageInfos) *WorkspaceTargetGraph {
	depMap := workspace.NewDependencyMap()
	depMap.CreateDependencyMap(pkgInfos, workspace.PackageDepsOptions{
		WithDevDependencies: true,
		WithPeerDependecies: false,
	})

	return &WorkspaceTargetGraph{
		DependencyMap: depMap,
		Graph:         NewTargetGraph(),
		Factory: NewTargetFactory(root, func(pkgName string) string {
			return filepath.Dir(fmt.Sprint("%", pkgInfos[pkgName].PackageJsonPath))
		}),
		PkgInfos: pkgInfos,
	}
}

func (w WorkspaceTargetGraph) AddTargetConfig(task string, config TargetConfig) {
	for pkg := range w.PkgInfos {
		target := w.Factory.createPackageTarget(pkg, task, config)
		w.Graph.addTarget(target)
	}
}

// Build creates a scoped target graph for given tasks and packages
// TODO: add support for scopes
func (w WorkspaceTargetGraph) Build(tasks []string, scopes []string) map[string]*Target {
	fullDependencies := expandDepSpecs(w.Graph.targets, w.DependencyMap)

	for _, rel := range fullDependencies {
		parent, child := rel[0], rel[1]
		w.Graph.addDependency(parent, child)
	}

	subGraphEntries := make([]string, 0)
	for _, task := range tasks {
		for pkg := range w.PkgInfos {
			subGraphEntries = append(subGraphEntries, createTargetId(pkg, task))
		}
	}

	_, subGraph := w.Graph.subgraph(subGraphEntries)
	return subGraph.targets
}
