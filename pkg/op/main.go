package op

import (
	"errors"
)

func Init() error {
	versionOk, err := checkOpCliVersion(executor)

	if err != nil {
		return err
	}

	if !versionOk {
		return errors.New("op-cli version not ok")
	}

	return nil
}
