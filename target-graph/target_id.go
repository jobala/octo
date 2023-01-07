package targetgraph

import (
	"errors"
	"fmt"
	"strings"
)

const START_TARGET_ID = "__start"

// createTargetId creates an id from the npm package's name and the task the target runs
//
// For example
//
//	createTargetId("hello", "build")
//
// Will create a hello#build id
func createTargetId(pkg, task string) string {
	if pkg == "" {
		return fmt.Sprintf("#%s", task)
	}
	return fmt.Sprintf("%s#%s", pkg, task)
}

// getPackageAndTask returns the package and task from a targetId.
//
// For example
//
//	getPackageAndTask("hello#build")
//
// Will return hello and build
func getPackageAndTask(targetId string) (error, string, string) {
	if !strings.Contains(targetId, "#") {
		return errors.New("Invalid targetId"), "", ""
	}

	id := strings.Split(targetId, "#")
	pkg, task := id[0], id[1]

	return nil, pkg, task
}
