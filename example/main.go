package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hpdobrica/go-op"
	"github.com/joho/godotenv"
)

var testVault string = "Sandbox"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	opcli := op.New()

	signinAddr := os.Getenv("OP_SIGNIN_ADDRESS")
	email := os.Getenv("OP_EMAIL")
	secretKey := os.Getenv("OP_SECRET_KEY")
	masterPass := os.Getenv("OP_MASTER_PASSWORD")
	signinErr := opcli.Signin(signinAddr, email, secretKey, masterPass)

	if signinErr != nil {
		fmt.Println("error while signing in")
		fmt.Println(signinErr)
		os.Exit(1)
	}

	fmt.Println("successfully initialized op")

	// exampleListItems(opcli)
	// exampleGetItem(opcli)

	// exampleListTemplates(opcli)

	exampleListVaults(opcli)

}

func exampleListItems(op op.Op) {
	items, err := op.ListItems(testVault)

	if err != nil {
		fmt.Println(err)
	}

	for _, item := range items {
		fmt.Println(item.Overview.Title)
	}
}

func exampleListTemplates(op op.Op) {
	templates, err := op.ListTemplates()

	if err != nil {
		fmt.Println(err)
	}

	for _, tpl := range templates {
		fmt.Println(tpl.Name)
	}

}

func exampleGetItem(op op.Op) {
	items, _ := op.ListItems(testVault)

	item, err := op.GetItem(items[0].Uuid)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(item.Details.Fields[0].Value)

}

func exampleListVaults(op op.Op) {
	vaults, err := op.ListVaults()

	if err != nil {
		fmt.Println(err)
	}

	for _, vault := range vaults {
		fmt.Println(vault.Name)
	}

}
