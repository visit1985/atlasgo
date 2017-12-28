# WIP: AtlasGo: Go SDK for MongoDB Atlas API

This project is work in progress. Eventually it will become a Go SDK for the [MongoDB Atlas API](https://docs.atlas.mongodb.com/api/). You are welcome for contribution.


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
	os.Setenv("ATLAS_USERNAME", "username")
	os.Setenv("ATLAS_ACCESS_KEY", "secret")

	client := atlas.NewClient()
	if client.Error == nil {
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
