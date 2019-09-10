package semvercomp

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Version represents the version number following X.Y.Z nomenclature
// X (Major): Version when you make incompatible API changes
// Y (Minor): Version when you add functionality in a backwards-compatible manner
// Z (Patch): Version when you make backwards-compatible bug fixes
// Source: Semantic Versioning 2.0.0 https://semver.org/
type Version struct {
	Major int64
	Minor int64
	Patch int64
}

// Relation enumerates the different relationships between version numbers
type Relation string

const (
	// Greater stands for a greater/newer version
	Greater Relation = "Greater"
	// Lower stands for a lower/older version
	Lower Relation = "Lower"

	// Equal describes the case when two versions are the same
	Equal Relation = "Equal"
)

// cleanVersionString checks for extra characters in a version string
// and removes them in order to parse the string to Version struct
func cleanVersionString(versionString string) string {
	re := regexp.MustCompile("(\\d|\\.\\d+)*$")
	result := re.FindAllString(versionString, -1)[0]
	return result
}

// parseTo64BitInteger shorthand for "strconv.ParseInt"
func parseTo64BitInteger(numStr string) int64 {
	if number, err := strconv.ParseInt(numStr, 10, 32); err == nil {
		return number
	}

	panic(fmt.Sprintf("Unable to parse %s to int64.", numStr))
}

// ParseStringToVersion parses a semantic version string into a Version struct
func ParseStringToVersion(version string) Version {
	versionArray := strings.Split(cleanVersionString(version), ".")

	return Version{
		Major: parseTo64BitInteger(versionArray[0]),
		Minor: parseTo64BitInteger(versionArray[1]),
		Patch: parseTo64BitInteger(versionArray[2]),
	}
}

// String returns the string from a Version struct
func (version Version) String() string {
	return fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Patch)
}

// IsSameVersion evaluates if two versions are equal
func (version Version) IsSame(otherVersion Version) bool {
	if version.Major == otherVersion.Major {
		if version.Minor == otherVersion.Minor {
			if version.Patch == otherVersion.Patch {
				return true
			}

			return false
		}

		return false
	}

	return false
}

// Relationship returns the Relation between two versions based in versionA as point of comparison
func (version Version) Relationship(otherVersion Version) Relation {
	if version.IsSame(otherVersion) {
		return Equal
	}

	if version.Major == otherVersion.Major {
		if version.Minor == otherVersion.Minor {
			if version.Patch > otherVersion.Patch {
				return Greater
			}

			return Lower
		}

		if version.Minor > otherVersion.Minor {
			return Greater
		}

		return Lower
	}

	if version.Major > otherVersion.Major {
		return Greater
	}

	return Lower
}

// StrRelationship returns the Relation between two versions as strings
func StrRelationship(versionA string, versionB string) Relation {
	verA := ParseStringToVersion(versionA)
	verB := ParseStringToVersion(versionB)

	return verA.Relationship(verB)
}

// GreaterVersion receives an slice of versions and returns the greater version
func GreaterVersion(versions []string) string {
	var greaterVersion = "0.0.0"

	for _, version := range versions {
		if StrRelationship(version, greaterVersion) == Greater {
			greaterVersion = version
		}
	}

	return greaterVersion
}
