package op

import (
	"encoding/json"
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

	os.Setenv("OP_SESSION", token)

	return err
}

func ListItems(vault string) ([]Item, error) {
	flags := make(map[string]string)

	if vault != "" {
		flags["vault"] = vault
	}

	out, err := executor.RunOp("op list items", flags)

	if err != nil {
		return nil, err
	}

	var items []Item
	errJson := json.Unmarshal([]byte(out), &items)

	if errJson != nil {
		return nil, errJson
	}

	return items, nil
}
