package semvercomp

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

// Version represents the version number following X.Y.Z nomenclature
// X (Major): Version when you make incompatible API changes
// Y (Minor): Version when you add functionality in a backwards-compatible manner
// Z (Patch): Version when you make backwards-compatible bug fixes
// Prerelease: Version is unstable and might not satisfy the intended compatibility requirements
// Source: Semantic Versioning 2.0.0 https://semver.org/
type Version struct {
	Major      int64
	Minor      int64
	Patch      int64
	PreRelease string
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

	SemverRegexp string = `^v?(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`
)

// cleanVersionString checks for extra characters in a version string
// and removes them in order to parse the string to Version struct
func cleanVersionString(versionString string) map[string]string {
	var re *regexp.Regexp

	re = regexp.MustCompile(SemverRegexp)

	result := re.FindStringSubmatch(versionString)
	versionFields := make(map[string]string)

	for index, member := range re.SubexpNames() {
		if index != 0 && member != "" {
			versionFields[member] = result[index]
		}
	}
	return versionFields
}

// parseTo64BitInteger shorthand for "strconv.ParseInt"
func parseTo64BitInteger(numStr string) int64 {
	if number, err := strconv.ParseInt(numStr, 10, 32); err == nil {
		return number
	}

	panic(fmt.Sprintf("Unable to parse %s to int64.", numStr))
}

// NewVersionFromString parses a semantic version string into a Version struct
func NewVersionFromString(version string) (Version, error) {
	if !isValid(version) {
		return Version{}, fmt.Errorf("provided tag (%s) is invalid", version)
	}

	versionMap := cleanVersionString(version)

	return Version{
		Major:      parseTo64BitInteger(versionMap["major"]),
		Minor:      parseTo64BitInteger(versionMap["minor"]),
		Patch:      parseTo64BitInteger(versionMap["patch"]),
		PreRelease: versionMap["prerelease"],
	}, nil
}

// String returns the string from a Version struct
func (version Version) String() string {
	versionString := fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Patch)
	if len(version.PreRelease) != 0 {
		versionString += fmt.Sprintf("-%s", version.PreRelease)
	}
	return versionString
}

// IsSame evaluates if two versions are equal
func (version Version) IsSame(otherVersion Version) bool {
	return reflect.DeepEqual(version, otherVersion)
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
func StrRelationship(versionA string, versionB string) (Relation, error) {
	verA, err := NewVersionFromString(versionA)
	if err != nil {
		return "", err
	}
	verB, err := NewVersionFromString(versionB)
	if err != nil {
		return "", err
	}
	return verA.Relationship(verB), nil
}

// GreaterVersion receives an slice of versions and returns the greater version
func GreaterVersion(versions []string) (string, error) {
	var greaterVersion = "0.0.0"

	for _, version := range versions {
		relation, err := StrRelationship(version, greaterVersion)
		if err != nil {
			return "", err
		}
		if relation == Greater {
			greaterVersion = version
		}
	}

	return greaterVersion, nil
}

//isValid() validates the version string
func isValid(version string) bool {
	pattern, _ := regexp.Compile(SemverRegexp)
	return pattern.MatchString(version)
}
