// Package twofer contains solution and tests for twofer exercise
package twofer

import "testing"
var tests = []struct {
	name, expected string
}{
	{"", "One for you, one for me."},
	{"Alice", "One for Alice, one for me."},
	{"Bob", "One for Bob, one for me."},
}

// TestShareWith is a benchmarking function for ShareWith().
func TestShareWith(t *testing.T) {
	for _, test := range tests {
		if observed := ShareWith(test.name); observed != test.expected {
			t.Fatalf("ShareWith(%s) = \"%v\", want \"%v\"", test.name, observed, test.expected)
		}
	}
}

// BenchmarkShareWith is a benchmarking function for ShareWith().
func BenchmarkShareWith(b *testing.B) {
	for i := 0; i < b.N; i++ {

		for _, test := range tests {
			ShareWith(test.name)
		}

	}
}
