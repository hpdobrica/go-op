package op

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/coreos/go-semver/semver"
)

var minOpCliVersion = "1.3.0"

func (o Op) checkOpCliVersion() (bool, error) {

	flags := make(map[string]string)
	flags["version"] = ""
	out, err := o.executor.RunOp("op", flags)

	if err != nil {
		return false, fmt.Errorf("Error getting op version: %s", err)
	}

	currentVersion := semver.New(strings.TrimSuffix(string(out), "\n"))

	requiredVersion := semver.New(minOpCliVersion)

	if currentVersion.LessThan(*requiredVersion) {
		return false, nil
	}

	return true, nil
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
