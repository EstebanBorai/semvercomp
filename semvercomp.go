package semvercomp

import (
	"fmt"
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

func parseTo64BitInteger(numStr string) int64 {
	if number, err := strconv.ParseInt(numStr, 10, 32); err == nil {
		return number
	}

	panic(fmt.Sprintf("Unable to parse %s to int64.", numStr))
}

// ParseStringToVersion parses a semantic version string into a Version struct
func ParseStringToVersion(version string) Version {
	versionArray := strings.Split(version, ".")

	return Version{
		Major: parseTo64BitInteger(versionArray[0]),
		Minor: parseTo64BitInteger(versionArray[1]),
		Patch: parseTo64BitInteger(versionArray[2]),
	}
}

// GetVersionString returns the string version of a Version struct
func GetVersionString(version Version) string {
	return fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Patch)
}

// IsSameVersion returns true if two versions are equal
func IsSameVersion(versionA Version, versionB Version) bool {
	if versionA.Major == versionB.Minor {
		if versionA.Minor == versionB.Minor {
			if versionA.Patch == versionB.Patch {
				return true
			}

			return false
		}

		return false
	}

	return false
}

// IsGreater returns true if versionA is newer/greater than versionB
func IsGreater(versionA Version, versionB Version) bool {
	if IsSameVersion(versionA, versionB) {
		return false
	}

	if versionA.Major == versionB.Major {
		if versionA.Minor == versionB.Minor {
			if versionA.Patch > versionB.Patch {
				return true
			}

			return false
		}

		if versionA.Minor > versionB.Minor {
			return true
		}

		return false
	}

	if versionA.Major > versionB.Major {
		return true
	}

	return false
}

// IsStringGreater returns true if versionA is newer/greater than versionB
func IsStringGreater(versionA string, versionB string) bool {
	verA := ParseStringToVersion(versionA)
	verB := ParseStringToVersion(versionB)

	return IsGreater(verA, verB)
}

// GreaterVersion receives an slice of versions and returns the greater version
func GreaterVersion(versions []string) string {
	var greaterVersion = "0.0.0"

	for _, version := range versions {
		if IsStringGreater(version, greaterVersion) {
			greaterVersion = version
		}
	}

	return greaterVersion
}
