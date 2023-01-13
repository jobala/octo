package workspace

type PackageInfo struct {
	Name             string
	PackageJsonPath  string
	Version          string
	Dependencies     []string
	DevDependencies  []string
	PeerDependencies []string
	Private          bool
	Group            string
	Scripts          map[string]string
	Repository       string
}

type PackageInfos = map[string]PackageInfo
