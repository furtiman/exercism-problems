// Package twofer contains solution and tests for twofer exercise
package twofer

// ShareWith returns a string of form "One for X, one for me.", where X is passed as a parameter.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}
