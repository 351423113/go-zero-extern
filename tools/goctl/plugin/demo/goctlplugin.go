package main

import (
	"fmt"

	"github.com/351423113/go-zero-extern/tools/goctl/plugin"
)

func main() {
	plugin, err := plugin.NewPlugin()
	if err != nil {
		panic(err)
	}

	if plugin.Api != nil {
		fmt.Printf("api: %+v \n", plugin.Api)
	}
	fmt.Println("Enjoy anything you want.")
}
