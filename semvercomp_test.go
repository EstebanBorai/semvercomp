package semvercomp

import (
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	t.Run("Expect version struct to be printed as a string in major.minor.patch format", func(t *testing.T) {
		var ver = Version{
			Major: 1,
			Minor: 0,
			Patch: 0,
		}

		want := "1.0.0"
		if got := ver.String(); got != want {
			t.Errorf("String() = %s, want %s", got, want)
		}
	})
	t.Run("Expect version struct to be printed as a string in major.minor.patch-prerelease format", func(t *testing.T) {
		var ver = Version{
			Major:      1,
			Minor:      0,
			Patch:      0,
			PreRelease: "alpha",
		}

		want := "1.0.0-alpha"
		if got := ver.String(); got != want {
			t.Errorf("String() = %s, want %s", got, want)
		}
	})
}

func TestNewVersionFromString(t *testing.T) {
	t.Run("Expect version struct to be constructed when the version string is valid", func(t *testing.T) {

		type ExpectedVersion struct {
			versionString string
			want          Version
		}

		versions := []ExpectedVersion{
			{
				versionString: "0.1.12",
				want: Version{
					Major: 0,
					Minor: 1,
					Patch: 12,
				},
			},
			{
				versionString: "0.1.12-alpha",
				want: Version{
					Major:      0,
					Minor:      1,
					Patch:      12,
					PreRelease: "alpha",
				},
			},
		}
		for index := range versions {
			current := versions[index]
			got, _ := NewVersionFromString(current.versionString)
			if !reflect.DeepEqual(got, current.want) {
				t.Errorf("NewVersionFromString(%s) = Version(Major: %d, Minor: %d, Patch: %d), want Version(Major: %d, Minor: %d, Patch: %d)",
					current.versionString, got.Major, got.Minor, got.Patch, got.Major, got.Minor, got.Patch)
			}
		}
	})

	t.Run("Expect error when version string is invalid", func(t *testing.T) {
		var versionString = "1.0_alpha.1"
		_, err := NewVersionFromString(versionString)
		if err == nil {
			t.Errorf("Expected validation error")
		}
	})

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

	if got, _ := StrRelationship(versionA, versionB); got != want {
		t.Errorf("[String] [Test Major Version] - StrRelationship(%s, %s) = %s, want %s",
			versionA, versionB, got, want)
	}
}

func TestStrRelationshipWithMinor(t *testing.T) {
	var versionA = "1.3.1"
	var versionB = "1.1.1"

	want := Greater

	if got, _ := StrRelationship(versionA, versionB); got != want {
		t.Errorf("[String] [Test Minor Version] - StrRelationship(%s, %s) = %s, want %s",
			versionA, versionB, got, want)
	}
}

func TestStrRelationshipWithPatch(t *testing.T) {
	var versionA = "1.1.3"
	var versionB = "1.1.1"

	want := Greater

	if got, _ := StrRelationship(versionA, versionB); got != want {
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
		{
			versionA: "2.0.0",
			versionB: "v1.0.0",
			expects:  Greater,
		},
		{
			versionA: "v2.0.0",
			versionB: "1.0.0",
			expects:  Greater,
		},
		{
			versionA: "0.0.0",
			versionB: "0.0.0",
			expects:  Equal,
		},
		{
			versionA: "0.0.2",
			versionB: "0.0.1",
			expects:  Greater,
		},
		{
			versionA: "0.1.0",
			versionB: "0.0.9",
			expects:  Greater,
		},
		{
			versionA: "1.0.0",
			versionB: "v1.0.0",
			expects:  Equal,
		},
		{
			versionA: "1.1.0",
			versionB: "1.1.0",
			expects:  Equal,
		},
		{
			versionA: "1.1.2",
			versionB: "1.1.2",
			expects:  Equal,
		},
		{
			versionA: "2.0.0",
			versionB: "1.1.1",
			expects:  Greater,
		},
		{
			versionA: "1.1.2",
			versionB: "1.2.1",
			expects:  Lower,
		},
		{
			versionA: "1.1.1",
			versionB: "2.1.1",
			expects:  Lower,
		},
		{
			versionA: "0.37.1",
			versionB: "0.37.1",
			expects:  Equal,
		},
		{
			versionA: "1.2.2-alpha",
			versionB: "1.2.2-alpha",
			expects:  Equal,
		},
	}

	for index := range versions {
		current := versions[index]
		result, _ := StrRelationship(current.versionA, current.versionB)

		if result != current.expects {
			t.Errorf("StrRelationship(%s, %s) = %s, want %s", current.versionA, current.versionB, result, current.expects)
		}
	}
}

func TestCleanVersionString(t *testing.T) {
	versions := [3]string{
		"v1.0.0",
		"0.37.1",
		"v1.2.0-alpha",
	}

	expected := [3]map[string]string{
		{"major": "1", "minor": "0", "patch": "0", "buildmetadata": "", "prerelease": ""},
		{"major": "0", "minor": "37", "patch": "1", "buildmetadata": "", "prerelease": ""},
		{"major": "1", "minor": "2", "patch": "0", "prerelease": "alpha", "buildmetadata": ""},
	}

	for index, version := range versions {
		exp := cleanVersionString(version)

		if !reflect.DeepEqual(exp, expected[index]) {
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

	greaterVersion, _ := GreaterVersion(versions)

	var expect = "v8.12.4"

	if expect != greaterVersion {
		t.Errorf("GreaterVersion(versions) = %s, want %s", greaterVersion, expect)
	}
}

func TestValidVersionString(t *testing.T) {
	versions := []string{
		"4.4.3",
		"8",
		"A1.2.3",
		"v2.3.0",
		"V2.3.0",
		"0.1.0-alpha.0",
		"1.0.0-alpha+001",
	}
	expected := [7]bool{
		true,
		false,
		false,
		true,
		false,
		true,
		true,
	}
	for index, version := range versions {
		isValid := isValid(version)
		if expected[index] != isValid {
			t.Errorf("isValid(%s)=%v,want %v", version, isValid, expected[index])
		}
	}
}
