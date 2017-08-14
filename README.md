# AtlasGo: Go SDK for MongoDB Atlas API

This is an unofficial Go SDK for the MongoDB Atlas API. You are welcome for contribution.


## Package Structure

*  common: common library of the Atlas Go SDK
*  util: utilities and helpers


## Quick Start

```go
package main

import (
	"fmt"
	"os"

	atlas "github.com/visit1985/atlasgo/common/client"
)

func main() {
	os.Setenv("ATLAS_GROUP_ID", "groupid")
	os.Setenv("ATLAS_USERNAME", "username")
	os.Setenv("ATLAS_ACCESS_KEY", "secret")

	client, err := atlas.NewClient()
	if err == nil {
		fmt.Printf("Endpoint: %s\n", client.Endpoint)
	}
}

```


## Documentation


## Build and Install


## Contributors

  * visit1985


## License

This project is licensed under the Apache License, Version 2.0. See [LICENSE](https://github.com/visit1985/atlasgo/blob/master/LICENSE.txt) for the full license text.


## Related projects
