// Package greeting contains HelloWorld and tests.
package greeting

import "testing"

// TestHelloWorld() is a test function for HelloWorld.
func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	if observed := HelloWorld(); observed != expected {
		t.Fatalf("HelloWorld() = %v, want %v", observed, expected)
	}
}

// BenchmarkHelloWorld() is a benchmarking function for HelloWorld.
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld()
	}
}
