package main

import (
	"fmt"
	"os"

	"github.com/hpdobrica/op/pkg/op"
)

func main() {
	secretKey := os.Getenv("OP_SECRET_KEY")
	masterPass := os.Getenv("OP_MASTER_PASSWORD")
	signinErr := op.Signin("my.1password.com", "hpdobrica@gmail.com", secretKey, masterPass)

	if signinErr != nil {
		fmt.Println(signinErr)
		os.Exit(1)
	}

	fmt.Println("successfully initialized op")

	_, err := op.ListItems("Private")

	// fmt.Println(out)
	fmt.Println(err)

}
