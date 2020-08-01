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

	os.Setenv("OP_SESSION_my", token)

	return err
}

func ListItems(vault string) (string, error) {
	token := os.Getenv("OP_SESSION_my")
	out, err := executor.Run(fmt.Sprintf("op list items --vault %s --session %s", vault, token))

	if err != nil {
		return "", err
	}

	var items []Item
	errJson := json.Unmarshal([]byte(out), &items)

	fmt.Println(errJson)

	fmt.Println(fmt.Sprintf("json parsed %s", PrettyPrint(items)))

	return out, nil
}
