package op

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/coreos/go-semver/semver"
)

var OP_CLI_VERSION = "1.3.0"

func checkOpCliVersion(executor IExecutor) (bool, error) {

	flags := make(map[string]string)
	flags["version"] = ""
	out, err := executor.RunOp("op", flags)

	if err != nil {
		return false, fmt.Errorf("Error getting op version: %s", err)
	}

	currentVersion := semver.New(strings.TrimSuffix(string(out), "\n"))

	requiredVersion := semver.New(OP_CLI_VERSION)

	if currentVersion.LessThan(*requiredVersion) {
		return false, nil
	}

	return true, nil
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
