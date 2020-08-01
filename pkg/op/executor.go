package op

import (
	"bufio"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

type IExecutor interface {
	Run(string, ...string) (string, error)
}

type Executor struct{}

func (r Executor) Run(command string, args ...string) (string, error) {
	out, err := exec.Command("/bin/sh", "-c", command).CombinedOutput()

	result := strings.TrimSuffix(string(out), "\n")

	fmt.Println(fmt.Sprintf("command: %q", command))
	fmt.Println(fmt.Sprintf("result: %q", result))
	fmt.Println(fmt.Sprintf("err: %q", err))

	if err != nil {
		fmt.Println(err)
		var errorMessage strings.Builder
		scanner := bufio.NewScanner(strings.NewReader(result))
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "[ERROR]") {
				errorMessage.WriteString(line)
			}
		}

		return "", errors.New(errorMessage.String())
	}

	return result, nil
}

var executor Executor
