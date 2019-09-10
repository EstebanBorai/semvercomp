package semvercomp

import (
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	var ver = Version{
		Major: 1,
		Minor: 0,
		Patch: 0,
	}

	want := "1.0.0"
	if got := ver.String(); got != want {
		t.Errorf("String() = %s, want %s", got, want)
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

func TestRelationshipWithMajor(t *testing.T) {
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

	if got := versionA.Relationship(versionB); got != want {
		t.Errorf("[Test Major Version] - Relationship(%s, %s) = %s, want %s",
			versionA.String(), versionB.String(), got, want)
	}
}

func TestRelationshipWithMinor(t *testing.T) {
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

	if got := versionA.Relationship(versionB); got != want {
		t.Errorf("[Test Minor Version] - Relationship(%s, %s) = %s, want %s",
			versionA.String(), versionB.String(), got, want)
	}
}

func TestRelationshipWithPatch(t *testing.T) {
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

	if got := versionA.Relationship(versionB); got != want {
		t.Errorf("[Test Patch Version] - Relationship(%s, %s) = %s, want %s",
			versionA.String(), versionB.String(), got, want)
	}
}

func TestStrRelationshipWithMajor(t *testing.T) {
	var versionA = "2.0.0"
	var versionB = "1.1.1"

	want := Greater

	if got := StrRelationship(versionA, versionB); got != want {
		t.Errorf("[String] [Test Major Version] - StrRelationship(%s, %s) = %s, want %s",
			versionA, versionB, got, want)
	}
}

func TestStrRelationshipWithMinor(t *testing.T) {
	var versionA = "1.3.1"
	var versionB = "1.1.1"

	want := Greater

	if got := StrRelationship(versionA, versionB); got != want {
		t.Errorf("[String] [Test Minor Version] - StrRelationship(%s, %s) = %s, want %s",
			versionA, versionB, got, want)
	}
}

func TestStrRelationshipWithPatch(t *testing.T) {
	var versionA = "1.1.3"
	var versionB = "1.1.1"

	want := Greater

	if got := StrRelationship(versionA, versionB); got != want {
		t.Errorf("[String] [Test Patch Version] - StrRelationship(%s, %s) = %s, want %s",
			versionA, versionB, got, want)
	}
}

func TestStrRelationship(t *testing.T) {
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
		result := StrRelationship(current.versionA, current.versionB)

		if result != current.expects {
			t.Errorf("StrRelationship(%s, %s) = %s, want %s", current.versionA, current.versionB, result, current.expects)
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

func TestGreaterVersion(t *testing.T) {
	versions := []string{
		"4.4.3",
		"v8.12.4",
		"0.1.0",
		"7.3.3",
		"4.67.31",
	}

	greaterVersion := GreaterVersion(versions)

	var expect = "v8.12.4"

	if expect != greaterVersion {
		t.Errorf("GreaterVersion(versions) = %s, want %s", greaterVersion, expect)
	}
}
