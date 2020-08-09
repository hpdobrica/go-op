package op

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

var masterPassword, domain, email, secretKey string

var someAuthToken = "some-auth-token"

type FakeExecutor struct {
	// Run(string) (string, error)
	// RunOp(string, map[string]string) (string, error)
}

func (f FakeExecutor) Run(command string) (string, error) {
	if command == fmt.Sprintf("echo %s | op signin %s %s %s --raw", masterPassword, domain, email, secretKey) {
		return someAuthToken, nil
	}

	return "", errors.New("Error while running command")
}

func (f FakeExecutor) RunOp(command string, flags map[string]string) (string, error) {
	if command == "op" {
		if os.Getenv("TEST_BAD_OP_VERSION") == "true" {
			return "1.2.0", nil
		}
		return "1.3.0", nil

	}
	return "", errors.New("Error while running command")
}

func TestSignin(t *testing.T) {

	t.Run("happypath login", func(t *testing.T) {
		masterPassword = "some-master-password"
		domain = "my.1password.com"
		email = "some@email.com"
		secretKey = "A3-SOME-SECRET-KEY"

		defer os.Unsetenv("OP_SESSION")

		fakeExecutor := FakeExecutor{}

		opcli := newFake(fakeExecutor)

		err := opcli.Signin(domain, email, secretKey, masterPassword)

		if err != nil {
			t.Errorf("There was an error during TestSignin, %v", err)
		}

		tokenResult := os.Getenv("OP_SESSION")

		if tokenResult != someAuthToken {
			t.Errorf("should have set token to %q, got %q instead", someAuthToken, tokenResult)
		}
	})

	t.Run("wrong password login", func(t *testing.T) {
		masterPassword = "some-master-password"
		domain = "my.1password.com"
		email = "some@email.com"
		secretKey = "A3-SOME-SECRET-KEY"

		defer os.Unsetenv("OP_SESSION")

		fakeExecutor := FakeExecutor{}

		opcli := newFake(fakeExecutor)

		err := opcli.Signin(domain, email, secretKey, "invalid-master-password")

		if err == nil {
			t.Error("Should have failed with an error but didnt")
		}

		tokenResult := os.Getenv("OP_SESSION")

		if tokenResult == someAuthToken {
			t.Error("should not have successfully set the token, but it did")
		}

	})

	t.Run("bad OP version", func(t *testing.T) {
		masterPassword = "some-master-password"
		domain = "my.1password.com"
		email = "some@email.com"
		secretKey = "A3-SOME-SECRET-KEY"

		os.Setenv("TEST_BAD_OP_VERSION", "true")

		defer os.Unsetenv("TEST_BAD_OP_VERSION")

		fakeExecutor := FakeExecutor{}

		opcli := newFake(fakeExecutor)

		err := opcli.Signin(domain, email, secretKey, "invalid-master-password")

		if err == nil {
			fmt.Println(err)
			t.Error("Should have failed because of version but didnt")
		}

		tokenResult := os.Getenv("OP_SESSION")

		if tokenResult == someAuthToken {
			t.Error("should not have successfully set the token, but it did")
		}

	})

}
