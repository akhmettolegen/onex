package main

import (
	"fmt"
	"github.com/akhmettolegen/onex/internal/api"
	"os"
)

func main() {
	//cmd.Execute()

	server, err := api.NewServer()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	server.Start()
}
