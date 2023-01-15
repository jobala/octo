package targetgraph

import "github.com/jobala/octo/workspace-tools"

func expandDepSpecs(targets map[string]*Target, depMap *workspace.DependencyMap) [][]string {
	dependencies := make([][]string, 0)

	for _, target := range targets {
		tasks := target.TaskDependencies
		pkgName := target.PkgName
		parentNodeId := target.Id

		// All nodes are children of the START_TARGET_ID node
		dependencies = append(dependencies, []string{parentNodeId, START_TARGET_ID})

		if len(tasks) == 0 {
			continue
		}

		for _, taskName := range tasks {

			// topoligical dependencies
			targetDependencies := depMap.Dependencies[target.PkgName]
			isTopologicalDependency := string(taskName[0]) == "^"

			if isTopologicalDependency {
				task := string(taskName[1:])

				// Find all targets that are a dependency of the current target node and executes the current task
				dependencyTargetIds := findDependenciesByTask(task, targetDependencies, targets)
				for _, dep := range dependencyTargetIds {
					dependencies = append(dependencies, []string{dep, parentNodeId})
				}
			} else {
				targetNodeId := createTargetId(pkgName, taskName)

				if _, nodeExists := targets[targetNodeId]; nodeExists {
					dependencies = append(dependencies, []string{createTargetId(pkgName, taskName), parentNodeId})
				}
			}
		}

	}

	return dependencies
}

// findDependeciesByTask gets the list of a targets dependency that execute the passed task
func findDependenciesByTask(task string, dependencies []string, targets map[string]*Target) []string {
	res := make([]string, 0)
	dependenciesExist := len(dependencies) > 0

	if dependenciesExist {
		for _, target := range targets {
			isDependency := contains(dependencies, target.PkgName)
			executesTask := target.Task == task

			if isDependency && executesTask {
				res = append(res, target.Id)
			}
		}
	}

	return res
}
