package main

import (
	"fmt"
	"github.com/rickard-von-essen/go-parallels/prlapi"
)

func main() {

	server, err := prlapi.LoginLocal()
	if err == nil {
		server.PrintVmList()
		server.Disconnect()
	} else {
		fmt.Printf("Error: %s", err)
	}
}
