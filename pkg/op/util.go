package op

import (
	"fmt"
	"strings"

	"github.com/coreos/go-semver/semver"
)

var OP_CLI_VERSION = "1.3.0"

func checkOpCliVersion(executor IExecutor) (bool, error) {

	out, err := executor.Run("op", "--version")

	if err != nil {
		return false, fmt.Errorf("Error getting op version: %s", err)
	}

	currentVersion := semver.New(strings.TrimSuffix(string(out), "\n"))

	requiredVersion := semver.New(OP_CLI_VERSION)

	if currentVersion.LessThan(*requiredVersion) {
		return false, nil
	}

	return true, nil

	// if c.Check(v) {
	// 	return "op", nil
	// }

	// return "", fmt.Errorf("op version needs to be equal or greater than: %s", version)
}
