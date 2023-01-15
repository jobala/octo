package targetgraph

import "path/filepath"

func NewTargetFactory(root string, resolve func(string) string) TargetFactory {
	return TargetFactory{
		options: Options{
			root:    root,
			resolve: resolve,
		},
	}
}

func (tf TargetFactory) createPackageTarget(pkgName, task string, config TargetConfig) *Target {
	// TODO: make Target a value type
	return &Target{
		Id:               createTargetId(pkgName, task),
		PkgName:          pkgName,
		Cwd:              filepath.Dir(tf.options.resolve(pkgName)),
		Task:             task,
		Type:             config.Type,
		Inputs:           config.Inputs,
		Outputs:          config.Outputs,
		Cache:            IF(config.Cache == true),
		TaskDependencies: config.DependsOn,
		Dependencies:     make([]string, 0),
		Dependents:       make([]string, 0),
	}
}

func IF(predicate bool) bool {
	if predicate {
		return true
	}
	return false
}

type TargetFactory struct {
	options Options
}

type Options struct {
	root    string
	resolve func(string) string
}

type TargetConfig struct {
	// Type can be one of npmScript, worker etc
	Type string

	// DependsOn are the target's dependencies of the form task or ^task
	DependsOn []string

	// Inputs are used to determine the target's cache key
	Inputs []string

	// Outputs contains list of files to be stored for caching
	Outputs []string

	// Cache used to determine whether we should cache this targets output or not, defaults to true
	Cache bool
}
