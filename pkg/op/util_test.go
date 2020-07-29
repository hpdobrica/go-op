package op

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestFindExistingOPClient(t *testing.T) {
	got, err := FindExistingOPClient(TestRunner{})

	if err != nil {
		fmt.Println()
		t.Fatalf("Whatever.")
	}

	if got != "1.3.0" {
		t.Fatalf("Version aint 1.3.0")
	}
}

type TestRunner struct{}

func (r TestRunner) Run(command string, args ...string) ([]byte, error) {
	cs := []string{"-test.run=TestHelperProcess", "--"}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	out, err := cmd.CombinedOutput()
	return out, err
}

func TestHelperProcess(*testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	defer os.Exit(0)

	if os.Args[3] == "--version" {
		fmt.Println("1.3.0")
	}

}
