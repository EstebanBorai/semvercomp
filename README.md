# semvercomp
🆕 Tiny library to compare and process version numbers based on [semver](https://semver.org/) conventions.

## Installation
```go
go get github.com/estebanborai/semantic-version-comparison
```

## Usage

`semvercomp` helps validating and comparing versions that follows Semantic Versioning convention.

```go
package main

import (
	"fmt"
	semvercomp "github.com/estebanborai/semantic-version-comparison"
)

func main() {
	var vera, verb string = "v1.0.1", "1.0.1"

	if semvercomp.IsStringGreater(vera, verb) == semvercomp.Equal {
		fmt.Printf("Versions %s and %s are equal\n", vera, verb)
	}
}
```

## API

### Structs and Constants

#### Version Struct
The `Version` struct represents the version number following X.Y.Z nomenclature
- X (Major): Version when you make incompatible API changes
- Y (Minor): Version when you add functionality in a backwards-compatible manner
- Z (Patch): Version when you make backwards-compatible bug fixes

Source: [Semantic Versioning 2.0.0](https://semver.org/)

```go
type Version struct {
	Major int64
	Minor int64
	Patch int64
}
```

#### Relation Enumerable
Represents the different relationships between version numbers.

Let `verA` (stands for version A) and `verB` (stands for version B), be two version numbers,
where `verA` is `v1.1.0` version number, and `verB` is `0.1.0`.

- `Greater`: Case where a version is greater than the another.
	`verA` is `Greater` than `verB`.

- `Lower`: Case when a version is lower than the other.
	`verB` is `Lower` than `verA`

- `Equal`: Case when two versions are the same:
	`v1.0.1` is `Equal` to `v1.0.1`

### Parsing and Conversion

#### Parse a string to Version struct
Parsing an string to a `Version` struct is possible using the `ParseStringToVersion` function, as follows:

```go
var ver string = "v1.4.11"
	
version := semvercomp.ParseStringToVersion(ver)
	
fmt.Println(version.Major) // 1
fmt.Println(version.Minor) // 4
fmt.Println(version.Patch) // 11
```

#### String representation of a `Version` struct
Its possible to get the string of a `Version` struct using the `String` function.

```go
ver := semvercomp.Version{
	Major: 3,
	Minor: 9,
	Patch: 0,
}

var versionString string = semvercomp.String(ver)

fmt.Println(versionString) // "3.9.0"
```

### Version Evaluation
`semvercomp` comes with a couple functions to help evaluate versions, either as `Version` structs or versions as strings.

#### `Relationship(versionA Version, versionB Version) Relation`
Evaluates the relationship between two versions based on `versionA`
