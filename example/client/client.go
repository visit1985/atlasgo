package main

import (
    "fmt"
    "os"

    "github.com/visit1985/atlasgo/services/group"
)

func main() {
    os.Setenv("ATLAS_USERNAME", "username")
    os.Setenv("ATLAS_ACCESS_KEY", "access_key")

    output, err := group.New("group_id").GetIpWhitelist()
    if err == nil {
        fmt.Printf("%s\n", output)
    } else {
        fmt.Printf("%s\n", err)
    }
}
