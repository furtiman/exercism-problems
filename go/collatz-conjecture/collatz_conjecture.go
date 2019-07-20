// Package collatzconjecture contains solution and tests for 'collatzconjecture' exercise of exercism.io
package collatzconjecture

import (
	"errors"
)

// CollatzConjecture implements a solution for Collatz Conjecture (or 3x+1) problem {Given a number n, return the number of steps required to reach 1.}
func CollatzConjecture(input int) (steps int, err error) {
	steps = 0
	if input <= 0 {
		return steps, errors.New("Input should be greater or equal to zero")
	}
	if input == 1 {
		return 0, nil
	}
	if input % 2 == 0 {
		input = input / 2 
	} else {
		input = (input * 3) + 1
	}
	steps, err = CollatzConjecture(input)
	return steps + 1, err
}
