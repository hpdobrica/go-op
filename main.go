package op

import (
	"errors"
	"fmt"
	"os"
)

func Signin(domain, email, secretKey, masterPassword string) error {

	versionOk, versionErr := checkOpCliVersion(executor)

	if versionErr != nil {
		return versionErr
	}

	if !versionOk {
		return errors.New("op-cli version must be 1.3.0 or higher")
	}

	token, err := executor.Run(fmt.Sprintf("echo %s | op signin %s %s %s --raw", masterPassword, domain, email, secretKey))

	if err != nil {
		return err
	}

	os.Setenv("OP_SESSION", token)

	return nil
}
