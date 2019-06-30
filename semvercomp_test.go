package semvercomp

import (
	"reflect"
	"testing"
)

func TestPrintVersion(t *testing.T) {
	var ver = Version{
		Major: 1,
		Minor: 0,
		Patch: 0,
	}

	want := "1.0.0"
	if got := GetVersionString(ver); got != want {
		t.Errorf("GetVersionString() = %s, want %s", got, want)
	}
}

func TestParseStringToVersion(t *testing.T) {
	var stringVersion = "0.1.12"

	want := Version{
		Major: 0,
		Minor: 1,
		Patch: 12,
	}

	var got = ParseStringToVersion(stringVersion)

	var isEqual = reflect.DeepEqual(got, want)

	if !isEqual {
		t.Errorf("ParseStringToVersion(%s) = Version(Major: %d, Minor: %d, Patch: %d), want Version(Major: %d, Minor: %d, Patch: %d)",
			stringVersion, got.Major, got.Minor, got.Patch, got.Major, got.Minor, got.Patch)
	}
}

func TestIsGreaterWithMajor(t *testing.T) {
	var versionA = Version{
		Major: 2,
		Minor: 0,
		Patch: 0,
	}

	var versionB = Version{
		Major: 1,
		Minor: 1,
		Patch: 1,
	}

	want := Greater

	if got := IsGreater(versionA, versionB); got != want {
		t.Errorf("[Test Major Version] - IsGreater(%s, %s) = %s, want %s",
			GetVersionString(versionA), GetVersionString(versionB), got, want)
	}
}

func TestIsGreaterWithMinor(t *testing.T) {
	var versionA = Version{
		Major: 1,
		Minor: 1,
		Patch: 0,
	}

	var versionB = Version{
		Major: 1,
		Minor: 0,
		Patch: 1,
	}

	want := Greater

	if got := IsGreater(versionA, versionB); got != want {
		t.Errorf("[Test Minor Version] - IsGreater(%s, %s) = %s, want %s",
			GetVersionString(versionA), GetVersionString(versionB), got, want)
	}
}

func TestIsGreaterWithPatch(t *testing.T) {
	var versionA = Version{
		Major: 1,
		Minor: 0,
		Patch: 1,
	}

	var versionB = Version{
		Major: 1,
		Minor: 0,
		Patch: 0,
	}

	want := Greater

	if got := IsGreater(versionA, versionB); got != want {
		t.Errorf("[Test Patch Version] - IsGreater(%s, %s) = %s, want %s",
			GetVersionString(versionA), GetVersionString(versionB), got, want)
	}
}

func TestIsStringGreaterWithMajor(t *testing.T) {
	var versionA = "2.0.0"
	var versionB = "1.1.1"

	want := Greater

	if got := IsStringGreater(versionA, versionB); got != want {
		t.Errorf("[String] [Test Major Version] - IsStringGreater(%s, %s) = %s, want %s",
			versionA, versionB, got, want)
	}
}

func TestIsStringGreaterWithMinor(t *testing.T) {
	var versionA = "1.3.1"
	var versionB = "1.1.1"

	want := Greater

	if got := IsStringGreater(versionA, versionB); got != want {
		t.Errorf("[String] [Test Minor Version] - IsStringGreater(%s, %s) = %s, want %s",
			versionA, versionB, got, want)
	}
}

func TestIsStringGreaterWithPatch(t *testing.T) {
	var versionA = "1.1.3"
	var versionB = "1.1.1"

	want := Greater

	if got := IsStringGreater(versionA, versionB); got != want {
		t.Errorf("[String] [Test Patch Version] - IsStringGreater(%s, %s) = %s, want %s",
			versionA, versionB, got, want)
	}
}

func TestIsStringGreater(t *testing.T) {
	type ComparisonExpects struct {
		versionA string
		versionB string
		expects  Relation
	}

	versions := []ComparisonExpects{
		ComparisonExpects{
			versionA: "2.0.0",
			versionB: "v1.0.0",
			expects:  Greater,
		},
		ComparisonExpects{
			versionA: "V2.0.0",
			versionB: "1.0.0",
			expects:  Greater,
		},
		ComparisonExpects{
			versionA: "0.0.0",
			versionB: "0.0.0",
			expects:  Equal,
		},
		ComparisonExpects{
			versionA: "0.0.2",
			versionB: "0.0.1",
			expects:  Greater,
		},
		ComparisonExpects{
			versionA: "0.1.0",
			versionB: "0.0.9",
			expects:  Greater,
		},
		ComparisonExpects{
			versionA: "1.0.0",
			versionB: "v1.0.0",
			expects:  Equal,
		},
		ComparisonExpects{
			versionA: "1.1.0",
			versionB: "1.1.0",
			expects:  Equal,
		},
		ComparisonExpects{
			versionA: "1.1.2",
			versionB: "1.1.2",
			expects:  Equal,
		},
		ComparisonExpects{
			versionA: "2.0.0",
			versionB: "1.1.1",
			expects:  Greater,
		},
		ComparisonExpects{
			versionA: "1.1.2",
			versionB: "1.2.1",
			expects:  Lower,
		},
		ComparisonExpects{
			versionA: "1.1.1",
			versionB: "2.1.1",
			expects:  Lower,
		},
		ComparisonExpects{
			versionA: "0.37.1",
			versionB: "0.37.1",
			expects:  Equal,
		},
	}

	for index := range versions {
		current := versions[index]
		result := IsStringGreater(current.versionA, current.versionB)

		if result != current.expects {
			t.Errorf("IsStringGreater(%s, %s) = %s, want %s", current.versionA, current.versionB, result, current.expects)
		}
	}
}

func TestCleanVersionString(t *testing.T) {
	versions := [4]string{
		"v1.0.0",
		"V2.0.0",
		"0.37.1",
		"0.2.9999999999999999",
	}

	expected := [4]string{
		"1.0.0",
		"2.0.0",
		"0.37.1",
		"0.2.9999999999999999",
	}

	for index, version := range versions {
		exp := cleanVersionString(version)

		if exp != expected[index] {
			t.Errorf("cleanVersionString(%s) = %s, want %s", version, exp, expected[index])
		}
	}
}
