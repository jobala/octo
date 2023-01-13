package workspace

import "golang.org/x/exp/maps"

func NewDependencyMap() *DependencyMap {
	return &DependencyMap{
		Dependencies: make(map[string][]string),
		Dependents:   make(map[string][]string),
	}
}

func (dm DependencyMap) CreateDependencyMap(packages PackageInfos, options PackageDepsOptions) {
	for pkg, info := range packages {
		deps := getPackageDependencies(info, packages, options)

		for _, dep := range deps {
			if !contains(dm.Dependencies[pkg], dep) {
				dm.Dependencies[pkg] = append(dm.Dependencies[pkg], dep)
			}

			if !contains(dm.Dependents[dep], pkg) {
				dm.Dependents[dep] = append(dm.Dependents[dep], pkg)
			}
		}
	}
}

// getPackageDependencies takes a `PackageInfo` of one of the workspace packages and returns the list of
// other workspace packages it depends on.
func getPackageDependencies(info PackageInfo, packages PackageInfos, options PackageDepsOptions) []string {
	deps := make(map[string]string)

	maps.Copy(deps, info.Dependencies)

	if options.WithDevDependencies {
		maps.Copy(deps, info.DevDependencies)
	}

	if options.WithPeerDependecies {
		maps.Copy(deps, info.PeerDependencies)
	}

	// Collates a list of other workspace packages that the passed package info depends on
	res := make([]string, 0)
	for pkg := range packages {
		if _, ok := deps[pkg]; ok {
			res = append(res, pkg)
		}
	}

	return res
}

type DependencyMap struct {
	Dependencies map[string][]string
	Dependents   map[string][]string
}

type PackageDepsOptions struct {
	WithDevDependencies bool
	WithPeerDependecies bool
}

func contains(items []string, item string) bool {
	for _, k := range items {
		if k == item {
			return true
		}
	}

	return false
}
