package main

import (
	"fmt"
	"os"

	common_client "github.com/visit1985/atlasgo/common/client"
)

func main() {
	os.Setenv("ATLAS_GROUP_ID", "groupid")
	os.Setenv("ATLAS_USERNAME", "username")
	os.Setenv("ATLAS_ACCESS_KEY", "secret")

	client, err := common_client.New().Init()
	if err == nil {
		fmt.Printf("Endpoint: %s\n", client.Endpoint)
	}
}
