package simplemath

import "fmt"

type SemanticVersion struct {
	major, minor, patch int
}

func NewSemanticVersion(major, minor, patch int) SemanticVersion {
	return SemanticVersion{
		major: major,
		minor: minor,
		patch: patch,
	}
}

// Method declartion
func (sv SemanticVersion) String() string{
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)
}

//modifying state with pointer
func (sv *SemanticVersion) IncrementMajor(){
	sv.major += 1
}

//modifying state with pointer
func (sv *SemanticVersion) IncrementMinor(){
	sv.minor += 1
}

//modifying state with pointer
func (sv *SemanticVersion) IncrementPatch(){
	sv.patch += 1
}