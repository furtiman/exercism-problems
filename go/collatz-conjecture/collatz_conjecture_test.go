// Package collatzconjecture contains solution and tests for 'collatzconjecture' exercise of exercism.io
package collatzconjecture

import (
	"testing"
)

// TestCollatzConjecture() is a test function for CollatzConjecture.
func TestCollatzConjecture(t *testing.T) {
	for _, testCase := range testCases {
		steps, err := CollatzConjecture(testCase.input)
		if testCase.expectError {
			if err == nil {
				t.Fatalf("FAIL: %s\n\tCollatzConjecture(%v) expected an error, got %v and err %v",
					testCase.description, testCase.input, steps, err)
			}
		} else {
			if err != nil {
				t.Fatalf("FAIL: %s\n\tCollatzConjecture(%v) returns unexpected error %s",
					testCase.description, testCase.input, err.Error())
			}
			if steps != testCase.expected {
				t.Fatalf("FAIL: %s\n\tCollatzConjecture(%v) expected %v, got %v",
					testCase.description, testCase.input, testCase.expected, steps)
			}
		}
		t.Logf("PASS: %s", testCase.description)
	}
}

// BenchmarkCollatzConjecture() is a benchmarking function for CollatzConjecture.
func BenchmarkCollatzConjecture(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			CollatzConjecture(testCase.input)
		}
	}
}
