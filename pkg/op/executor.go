package op

import (
	"bufio"
	"errors"
	"os/exec"
	"strings"
)

type IExecutor interface {
	Run(string, ...string) (string, error)
}

type Executor struct{}

func (r Executor) Run(command string, args ...string) (string, error) {
	out, err := exec.Command(command, args...).CombinedOutput()

	if err != nil {
		var errorMessage strings.Builder
		scanner := bufio.NewScanner(strings.NewReader(string(out)))
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "[ERROR]") {
				errorMessage.WriteString(line)
			}
		}

		return "", errors.New(errorMessage.String())
	}

	return string(out), nil
}

var executor Executor
