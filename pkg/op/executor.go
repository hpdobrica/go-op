package op

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type IExecutor interface {
	Run(string) (string, error)
	RunOp(string, map[string]string) (string, error)
}

type Executor struct{}

func (r Executor) Run(command string) (string, error) {
	out, err := exec.Command("/bin/sh", "-c", command).CombinedOutput()

	result := strings.TrimSuffix(string(out), "\n")

	if err != nil {
		return "", errors.New(result)
	}

	return result, nil
}

func (r Executor) RunOp(command string, flags map[string]string) (string, error) {
	if flags == nil {
		flags = make(map[string]string)
	}

	token := os.Getenv("OP_SESSION")

	if token != "" {
		flags["session"] = token
	}

	out, err := r.Run(buildOpCommand(command, flags))

	if err != nil {
		fmt.Println(err)
		err = getOpError(out, err)
	}

	return out, err
}

func buildOpCommand(command string, flags map[string]string) string {
	var result strings.Builder

	result.WriteString(command)

	for k, v := range flags {
		result.WriteString(fmt.Sprintf(" --%s %s", k, v))
	}

	return result.String()
}

func getOpError(out string, err error) error {
	var errorMessage strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(out))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "[ERROR]") {
			errorMessage.WriteString(line)
		}
	}

	return errors.New(errorMessage.String())
}

var executor Executor
