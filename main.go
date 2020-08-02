package main

import (
	"fmt"
	"os"

	"github.com/hpdobrica/op/pkg/op"
)

func main() {
	signinAddr := os.Getenv("OP_SIGNIN_ADDRESS")
	email := os.Getenv("OP_EMAIL")
	secretKey := os.Getenv("OP_SECRET_KEY")
	masterPass := os.Getenv("OP_MASTER_PASSWORD")
	signinErr := op.Signin(signinAddr, email, secretKey, masterPass)

	if signinErr != nil {
		fmt.Println(signinErr)
		os.Exit(1)
	}

	fmt.Println("successfully initialized op")

	items, err := op.ListItems("Private")

	if err != nil {
		fmt.Println(err)
	}

	for _, item := range items {
		fmt.Println(item.Overview.Title)
	}

}
