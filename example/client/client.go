package main

import (
	"fmt"
	"os"

	atlas "github.com/visit1985/atlasgo/common/client"
)

func main() {
	os.Setenv("ATLAS_USERNAME", "username")
	os.Setenv("ATLAS_ACCESS_KEY", "secret")

	client := atlas.NewClient()
	if client.Error == nil {
		fmt.Printf("Endpoint: %s\n", client.Endpoint)
	}
}
