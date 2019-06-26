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

	want := true

	if got := IsGreater(versionA, versionB); got != want {
		t.Errorf("[Test Major Version] - IsGreater(%s, %s) = %t, want %t",
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

	want := true

	if got := IsGreater(versionA, versionB); got != want {
		t.Errorf("[Test Minor Version] - IsGreater(%s, %s) = %t, want %t",
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

	want := true

	if got := IsGreater(versionA, versionB); got != want {
		t.Errorf("[Test Patch Version] - IsGreater(%s, %s) = %t, want %t",
			GetVersionString(versionA), GetVersionString(versionB), got, want)
	}
}

func TestIsStringGreaterWithMajor(t *testing.T) {
	var versionA = "2.0.0"
	var versionB = "1.1.1"

	want := true

	if got := IsStringGreater(versionA, versionB); got != want {
		t.Errorf("[String] [Test Major Version] - IsStringGreater(%s, %s) = %t, want %t",
			versionA, versionB, got, want)
	}
}

func TestIsStringGreaterWithMinor(t *testing.T) {
	var versionA = "1.3.1"
	var versionB = "1.1.1"

	want := true

	if got := IsStringGreater(versionA, versionB); got != want {
		t.Errorf("[String] [Test Minor Version] - IsStringGreater(%s, %s) = %t, want %t",
			versionA, versionB, got, want)
	}
}

func TestIsStringGreaterWithPatch(t *testing.T) {
	var versionA = "1.1.3"
	var versionB = "1.1.1"

	want := true

	if got := IsStringGreater(versionA, versionB); got != want {
		t.Errorf("[String] [Test Patch Version] - IsStringGreater(%s, %s) = %t, want %t",
			versionA, versionB, got, want)
	}
}
