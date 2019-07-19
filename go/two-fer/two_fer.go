// Package twofer contains solution and tests for twofer exercise
package twofer

// ShareWith returns a string of form "One for X, one for me.", where X is passed as a parameter.
func ShareWith(name string) string {
	s1 := "One for "
	s2 := ", one for me."
	if name == ""{
		name = "you"
	}

	return s1 + name + s2
}