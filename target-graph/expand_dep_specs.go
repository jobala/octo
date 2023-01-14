package targetgraph

import "github.com/jobala/octo/workspace-tools"

func expandDepSpecs(targets map[string]*Target, depMap *workspace.DependencyMap) [][]string {
	dependencies := make([][]string, 0)

	for _, target := range targets {
		depSpecs := target.TaskDependencies
		pkgName := target.PkgName
		parentNodeId := target.Id

		// all parent nodes are children of the START_TARGET_ID node
		dependencies = append(dependencies, []string{parentNodeId, START_TARGET_ID})

		if len(depSpecs) == 0 {
			continue
		}

		for _, taskName := range depSpecs {
			task := string(taskName[1:])
			targetDependencies := depMap.Dependencies[target.PkgName]

			if string(taskName[0]) == "^" {
				dependencyTargetIds := findDependenciesByTask(task, targetDependencies, targets)
				for _, dep := range dependencyTargetIds {
					dependencies = append(dependencies, []string{dep, parentNodeId})
				}
			} else if pkgName != "" {
				if _, ok := targets[createTargetId(pkgName, task)]; ok {
					dependencies = append(dependencies, []string{createTargetId(pkgName, task), parentNodeId})
				}
			}
		}

	}

	return dependencies
}

func findDependenciesByTask(task string, dependencies []string, targets map[string]*Target) []string {
	res := make([]string, 0)

	if len(dependencies) > 0 {
		for _, target := range targets {
			if contains(dependencies, target.PkgName) && target.Task == task {
				res = append(res, target.Id)
			}
		}
	}

	return res
}
