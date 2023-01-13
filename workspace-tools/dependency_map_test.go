package workspace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDependencyMap(t *testing.T) {
	pkgA := PackageInfo{
		Name: "pkgA",
		Dependencies: Dependency{
			"pkgB": "14.0.0",
			"pkgC": "15.0.0",
		},
	}
	pkgB := PackageInfo{
		Name: "pkgA",
	}
	pkgC := PackageInfo{
		Name: "pkgA",
	}

	pkgInfos := PackageInfos{
		"pkgA": pkgA,
		"pkgB": pkgB,
		"pkgC": pkgC,
	}
	pkgOptions := PackageDepsOptions{
		WithDevDependencies: true,
		WithPeerDependecies: false,
	}

	dm := NewDependencyMap()
	dm.createDependencyMap(pkgInfos, pkgOptions)

	assert.Contains(t, dm.Dependencies["pkgA"], "pkgB")
	assert.Contains(t, dm.Dependencies["pkgA"], "pkgC")
}
