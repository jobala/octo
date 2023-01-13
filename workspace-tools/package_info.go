package workspace

type PackageInfo struct {
	Name             string
	PackageJsonPath  string
	Version          string
	Dependencies     Dependency
	DevDependencies  Dependency
	PeerDependencies Dependency
	Private          bool
	Group            string
	Scripts          map[string]string
	Repository       string
}

type Dependency map[string]string
type PackageInfos = map[string]PackageInfo
