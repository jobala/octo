package targetgraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackageId_CreateIdWithPkg(t *testing.T) {
	id := createTargetId("hello", "lint")
	assert.Equal(t, "hello#lint", id)
}

func TestPackageId_CreateIdWithoutPkg(t *testing.T) {
	id := createTargetId("", "lint")
	assert.Equal(t, "#lint", id)
}

func TestPackageId_GetPackageAndTask(t *testing.T) {
	targetId := "hello#build"

	err, pkg, task := getPackageAndTask(targetId)
	assert.NoError(t, err)

	assert.Equal(t, pkg, "hello")
	assert.Equal(t, task, "build")
}

func TestPackageId_InvalidTargetId(t *testing.T) {
	targetId := "hello,build"

	err, _, _ := getPackageAndTask(targetId)
	assert.Error(t, err)
}
