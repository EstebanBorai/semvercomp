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

- Structs and Constants
	- [Version](https://github.com/estebanborai/semantic-version-comparison#version-struct)
	- [Relation](https://github.com/estebanborai/semantic-version-comparison#relation-enumerable)
- Parsing and Conversion
	- 

### Structs and Constants

#### `Version` Struct
The `Version` struct represents the version number following `X.Y.Z` nomenclature
Version Number | Name | Description
------------ | -------------
`X` | `Major` | Version when you make incompatible API changes
`Y` | `Minor` | Version when you add functionality in a backwards-compatible manner
`Z` | `Patch` | Version when you make backwards-compatible bug fixes

Source: [Semantic Versioning 2.0.0](https://semver.org/)

```go
type Version struct {
	Major int64
	Minor int64
	Patch int64
}
```

#### `Relation` Enumerable
Represents the different relationships between version numbers.

Let `verA` (stands for version A) and `verB` (stands for version B), be two version numbers,
where `verA` is `v1.1.0` version number, and `verB` is `0.1.0`.

Enum Key | Description | Sample
--- | --- | ---
`Greater` | Case where a version is greater than the another. | `verA` is `Greater` than `verB`
`Lower` | Case when a version is lower than the other. | `verB` is `Lower` than `verA`
`Equal` | Case when two versions are the same. | `v1.0.1` is `Equal` to `v1.0.1`

### Parsing and Conversion

#### `ParseStringToVersion(version string) Version`
ParseStringToVersion parses a semantic version string into a Version struct.

```go
var ver string = "v1.4.11"
	
version := semvercomp.ParseStringToVersion(ver)
	
fmt.Println(version.Major) // 1
fmt.Println(version.Minor) // 4
fmt.Println(version.Patch) // 11
```

#### `String(version Version) string`
String returns the string from a Version struct.

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
`Relationship` returns the `Relation` between two versions based in `versionA` as point of comparison.

```go
vera := semvercomp.Version{
	Major: 1,
	Minor: 2,
	Patch: 8,
}
	
verb := semvercomp.Version{
	Major: 2,
	Minor: 4,
	Patch: 6,
}
		
fmt.Println(string(semvercomp.Relationship(verb, vera))) // Greater
```

It's possible to get the `Relation` of two versions as strings using `StrRelationship` function.

```go
var vera, verb string = "v1.2.8", "2.4.6"
		
fmt.Println(string(semvercomp.StrRelationship(verb, vera))) // Greater
```

#### `IsSameVersion(versionA Version, versionB Version) bool`
`IsSameVersion` evaluates if two versions are equal

```go
vera := semvercomp.Version{
	Major: 1,
	Minor: 1,
	Patch: 1,
}
	
verb := semvercomp.Version{
	Major: 1,
	Minor: 1,
	Patch: 1,
}
		
fmt.Println(semvercomp.IsSameVersion(verb, vera)) // true
```

#### `GreaterVersion(versions []string) string`
`GreaterVersion` receives a slice of versions and returns the `Greater` version.

```go
versions := []string{
	"4.4.3",
	"v8.12.4",
	"0.1.0",
	"7.3.3",
	"4.67.31",
}

greaterVersion := GreaterVersion(versions)

fmt.Println(greaterVersion) // "v8.12.4"
```
