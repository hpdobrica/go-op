package op

import (
	"errors"
	"fmt"
	"os"
)

// Op : 1Password package
type Op struct {
	executor IExecutor
}

// New : returns a new instance of Op
func New() Op {
	return Op{executor}
}

func newFake(e IExecutor) Op {
	return Op{e}
}

// Signin : creates 1password session which will be used for cli calls
func (o Op) Signin(domain, email, secretKey, masterPassword string) error {
	versionOk, versionErr := o.checkOpCliVersion()

	if versionErr != nil {
		return versionErr
	}

	if !versionOk {
		return errors.New("op-cli version must be 1.3.0 or higher")
	}

	token, err := o.executor.Run(fmt.Sprintf("echo %s | op signin %s %s %s --raw", masterPassword, domain, email, secretKey))

	if err != nil {
		return err
	}

	os.Setenv("OP_SESSION", token)

	return nil
}
