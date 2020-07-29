package main

import (
	"fmt"
	"os"

	"github.com/hpdobrica/op/pkg/op"
)

func main() {
	err := op.Init()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("successfully initialized op")

}
