# AtlasGo: Go SDK for MongoDB Atlas API

This project is work in progress. Eventually it will become a Go SDK for the [MongoDB Atlas API](https://docs.atlas.mongodb.com/api/). You are welcome for contribution.


## Package Structure

*  services: API services of Atlas Go SDK
*  common: common library of Atlas Go SDK


## Quick Start

```go
package main

import (
    "fmt"
    "os"
    "github.com/visit1985/atlasgo/services/group"
)

func main() {
    os.Setenv("ATLAS_USERNAME", "username")
    os.Setenv("ATLAS_ACCESS_KEY", "secret")

    output, err := group.New("group_id").GetWhitelist()
    if err == nil {
        fmt.Printf("%+v\n", output)
    } else {
        fmt.Println(err)
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
