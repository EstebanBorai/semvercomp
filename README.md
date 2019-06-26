# semvercomp
🆕 Tiny library to compare and process version numbers based on [semver](https://semver.org/) conventions.

### Installation
```go
go get github.com/estebanborai/semantic-version-comparison
```

### Usage
```go
package main

import (
	"fmt"
	semver "github.com/estebanborai/semantic-version-comparison"
)

func main() {
	fmt.Println(semver.IsStringGreater("1.10.0", "1.9.1")) // true
}
```
