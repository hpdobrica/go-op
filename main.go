package main

import (
	"fmt"
	"os"

	"github.com/hpdobrica/op/pkg/op"
)

var testVault string = "Sandbox"

func main() {
	signinAddr := os.Getenv("OP_SIGNIN_ADDRESS")
	email := os.Getenv("OP_EMAIL")
	secretKey := os.Getenv("OP_SECRET_KEY")
	masterPass := os.Getenv("OP_MASTER_PASSWORD")
	signinErr := op.Signin(signinAddr, email, secretKey, masterPass)

	if signinErr != nil {
		fmt.Println("error while signing in")
		fmt.Println(signinErr)
		os.Exit(1)
	}

	fmt.Println("successfully initialized op")

	// ExampleListItems()
	// ExampleGetItem()

	// ExampleListTemplates()

	ExampleListVaults()

}

func ExampleListItems() {
	items, err := op.ListItems(testVault)

	if err != nil {
		fmt.Println(err)
	}

	for _, item := range items {
		fmt.Println(item.Overview.Title)
	}
}

func ExampleListTemplates() {
	templates, err := op.ListTemplates()

	if err != nil {
		fmt.Println(err)
	}

	for _, tpl := range templates {
		fmt.Println(tpl.Name)
	}

}

func ExampleGetItem() {
	items, _ := op.ListItems(testVault)

	item, err := op.GetItem(items[0].Uuid)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(item.Details.Fields[0].Value)

}

func ExampleListVaults() {
	vaults, err := op.ListVaults()

	if err != nil {
		fmt.Println(err)
	}

	for _, vault := range vaults {
		fmt.Println(vault.Name)
	}

}
