package workspace

func createDependencyMap(packages PackageInfos, options PackageDepsOptions) {

}

type DependencyMap struct {
	Dependencies map[string]string
	Dependents   map[string]string
}

type PackageDepsOptions struct {
	WithDevDependencies bool
	WithPeerDependecies bool
}
