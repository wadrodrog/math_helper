/*
	Math Helper - implementation of common algorithms

# Copyright (C) 2024  wadrodrog

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package permutations

import (
	"fmt"
	"reflect"
	"testing"
)

// Should return correct error
func TestNewPermutation(t *testing.T) {
	tests := []struct {
		size      int
		arguments []int
		values    []int
		want      error
	}{
		// Permutation length does not equal n
		{3, []int{2, 1}, []int{2}, InvalidLengthError(2, 3)},
		{3, []int{2}, []int{2, 1}, InvalidLengthError(1, 3)},
		// Permutation has elements not in range 1..n
		{3, []int{1, 2, 4}, []int{2, 1, 3}, InvalidElementError(4, 3)},
		{3, []int{1, 2, 3}, []int{2, 0, 3}, InvalidElementError(0, 3)},
		// Permutation has repeating elements
		{3, []int{1, 2, 1}, []int{2, 1, 3}, RepeatingElementError(1)},
		{3, []int{3, 2, 1}, []int{2, 2, 3}, RepeatingElementError(2)},
		// No error
		{3, []int{3, 2, 1}, []int{2, 1, 3}, nil},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%v,%v", tt.size, tt.arguments, tt.values)
		t.Run(testname, func(t *testing.T) {
			_, err := NewPermutation(tt.size, tt.arguments, tt.values)
			if err == nil && tt.want != nil {
				t.Fatalf("no error %q", tt.want)
			}
			if err != nil && err.Error() != tt.want.Error() {
				t.Errorf("got %q, want %q", err, tt.want)
			}
		})
	}
}

// Should return correct error
func TestNewSequencePermutation(t *testing.T) {
	tests := []struct {
		n           int
		permutation []int
		want        error
	}{
		// Permutation length does not equal n
		{3, []int{2, 1}, InvalidLengthError(2, 3)},
		// Permutation has elements not in range 1..n
		{3, []int{1, 3, 4}, InvalidElementError(4, 3)},
		// Permutation has repeating elements
		{3, []int{1, 2, 1}, RepeatingElementError(1)},
		// No error
		{3, []int{3, 2, 1}, nil},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%v", tt.n, tt.permutation)
		t.Run(testname, func(t *testing.T) {
			_, err := NewSequencePermutation(tt.n, tt.permutation)
			if err == nil && tt.want != nil {
				t.Fatalf("no error %q", tt.want)
			}
			if err != nil && err.Error() != tt.want.Error() {
				t.Errorf("got %q, want %q", err, tt.want)
			}
		})
	}
}

// Should correctly transpose sequence permutation from transposition cycles
func TestSequencePermutationFromTranspositions(t *testing.T) {
	tests := []struct {
		transpositions  [][]int
		wantPermutation []int
		wantError       error
	}{
		{[][]int{{1, 4}, {4, 3, 1}}, []int{}, InvalidTranspositionError(1)},
		{[][]int{{1}, {1}}, []int{}, InvalidTranspositionError(0)},
		{[][]int{{1, 0}, {1, 2}}, []int{}, InvalidTranspositionError(0)},
		{[][]int{}, []int{}, InvalidTranspositionError(0)},
		{[][]int{{1, 4}, {4, 3}, {2, 5}}, []int{4, 5, 1, 3, 2}, nil},
		{[][]int{{1, 2}, {1, 6}, {1, 3}, {4, 7}}, []int{3, 1, 6, 7, 5, 2, 4}, nil},
		{[][]int{{1, 3}, {3, 6}, {6, 2}, {4, 7}}, []int{3, 1, 6, 7, 5, 2, 4}, nil},
		{[][]int{{1, 3}, {1, 5}, {2, 7}, {2, 4}}, []int{5, 4, 1, 7, 3, 6, 2}, nil},
		{[][]int{{1, 5}, {5, 3}, {2, 4}, {4, 7}}, []int{5, 4, 1, 7, 3, 6, 2}, nil},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.transpositions)
		t.Run(testname, func(t *testing.T) {
			permutation, err := NewSequencePermutationFromTranspositions(tt.transpositions)

			// Error
			if err == nil && tt.wantError != nil {
				t.Fatalf("no error %q", tt.wantError)
			}
			if err != nil && err.Error() != tt.wantError.Error() {
				t.Fatalf("got %q, want %q", err, tt.wantError)
			}

			// Success
			if err != nil && tt.wantError == nil {
				t.Fatalf("got an error while initializing Permutation: %v", err)
			}
			if err == nil {
				got := permutation.values
				if !reflect.DeepEqual(got, tt.wantPermutation) {
					t.Errorf("got %v, want %v", got, tt.wantPermutation)
				}
			}
		})
	}
}

// Should count inversions for sequence permutation
func TestPermutationInversions(t *testing.T) {
	tests := []struct {
		n           int
		permutation []int
		want        int
	}{
		{5, []int{2, 3, 5, 1, 4}, 4},
		{7, []int{6, 5, 1, 2, 7, 4, 3}, 12},
		{7, []int{5, 6, 4, 7, 2, 1, 3}, 15},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.permutation)
		t.Run(testname, func(t *testing.T) {
			permutation, err := NewSequencePermutation(tt.n, tt.permutation)
			if err != nil {
				t.Fatalf("got an error while initializing Permutation: %v", err)
			}

			got := permutation.Inversions()
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

// Should be even or not
func TestPermutationParity(t *testing.T) {
	tests := []struct {
		n         int
		arguments []int
		values    []int
		want      bool
	}{
		{7, []int{}, []int{5, 6, 4, 7, 2, 1, 3}, false},
		{8, []int{}, []int{3, 5, 2, 1, 6, 4, 8, 7}, true},
		{7, []int{3, 5, 6, 4, 2, 1, 7}, []int{2, 4, 1, 7, 6, 5, 3}, true},
		{8, []int{2, 7, 5, 4, 8, 3, 6, 1}, []int{3, 5, 8, 7, 2, 6, 1, 4}, false},
		{5, []int{3, 1, 2, 5, 4}, []int{5, 4, 3, 2, 1}, false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v:%v", tt.arguments, tt.values)
		t.Run(testname, func(t *testing.T) {
			var permutation *Permutation
			var err error
			if len(tt.arguments) == 0 {
				permutation, err = NewSequencePermutation(tt.n, tt.values)
			} else {
				permutation, err = NewPermutation(tt.n, tt.arguments, tt.values)
			}

			if err != nil {
				t.Fatalf("got an error while initializing Permutation: %v", err)
			}

			got := permutation.IsEven()
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

// Permutations count or factorial of n
func TestPermutationsCount(t *testing.T) {
	tests := []struct {
		n           int
		permutation []int
		want        uint
	}{
		{1, []int{1}, 1},
		{2, []int{1, 2}, 2},
		{3, []int{1, 2, 3}, 6},
		{7, []int{1, 2, 3, 4, 5, 6, 7}, 5040},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.n)
		t.Run(testname, func(t *testing.T) {
			permutation, err := NewSequencePermutation(tt.n, tt.permutation)
			if err != nil {
				t.Fatalf("got an error while initializing Permutation: %v", err)
			}

			got := permutation.Count()
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

// Should decompose permutation into cycles
func TestPermutationCycles(t *testing.T) {
	tests := []struct {
		n         int
		arguments []int
		values    []int
		want      [][]int
	}{
		{9, []int{}, []int{3, 9, 8, 6, 1, 4, 7, 5, 2}, [][]int{{1, 3, 8, 5}, {2, 9}, {4, 6}}},
		{7, []int{}, []int{6, 5, 1, 2, 7, 4, 3}, [][]int{{1, 6, 4, 2, 5, 7, 3}}},
		{5, []int{}, []int{4, 5, 1, 3, 2}, [][]int{{1, 4, 3}, {2, 5}}},
		{7, []int{}, []int{3, 1, 6, 7, 5, 2, 4}, [][]int{{1, 3, 6, 2}, {4, 7}}},
		{5, []int{3, 1, 2, 5, 4}, []int{5, 4, 3, 2, 1}, [][]int{{3, 5, 2}, {1, 4}}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v:%v", tt.arguments, tt.values)
		t.Run(testname, func(t *testing.T) {
			var permutation *Permutation
			var err error
			if len(tt.arguments) == 0 {
				permutation, err = NewSequencePermutation(tt.n, tt.values)
			} else {
				permutation, err = NewPermutation(tt.n, tt.arguments, tt.values)
			}

			if err != nil {
				t.Fatalf("got an error while initializing Permutation: %v", err)
			}

			got := permutation.Cycles()
			if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

// Should decompose permutation cycles into transpositions
func TestPermutationTranspositions(t *testing.T) {
	tests := []struct {
		n         int
		arguments []int
		values    []int
		want      [][]int
	}{
		{5, []int{}, []int{4, 5, 1, 3, 2}, [][]int{{1, 4}, {4, 3}, {2, 5}}},
		{7, []int{}, []int{5, 4, 1, 7, 3, 6, 2}, [][]int{{1, 5}, {5, 3}, {2, 4}, {4, 7}}},
		{7, []int{}, []int{3, 1, 6, 7, 5, 2, 4}, [][]int{{1, 3}, {3, 6}, {6, 2}, {4, 7}}},
		{5, []int{3, 1, 2, 5, 4}, []int{5, 4, 3, 2, 1}, [][]int{{3, 5}, {5, 2}, {1, 4}}},
		// Another way to decompose
		//{7, []int{}, []int{5, 4, 1, 7, 3, 6, 2}, [][]int{{1, 3}, {1, 5}, {2, 7}, {2, 4}}},
		//{7, []int{}, []int{3, 1, 6, 7, 5, 2, 4}, [][]int{{1, 2}, {1, 6}, {1, 3}, {4, 7}}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v:%v", tt.arguments, tt.values)
		t.Run(testname, func(t *testing.T) {
			var permutation *Permutation
			var err error
			if len(tt.arguments) == 0 {
				permutation, err = NewSequencePermutation(tt.n, tt.values)
			} else {
				permutation, err = NewPermutation(tt.n, tt.arguments, tt.values)
			}

			if err != nil {
				t.Fatalf("got an error while initializing Permutation: %v", err)
			}

			got := permutation.Transpositions()
			if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

// Should not mutliply permutations and return correct error.
func TestPermutationsMultiplyErrors(t *testing.T) {
	tests := []struct {
		n1      int
		values1 []int
		n2      int
		values2 []int
		want    error
	}{
		{3, []int{1, 2, 3}, 2, []int{2, 1}, InvalidLengthError(2, 3)},
		{2, []int{2, 1}, 3, []int{1, 2, 3}, InvalidLengthError(3, 2)},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%vx%v", tt.values1, tt.values2)
		t.Run(testname, func(t *testing.T) {
			permutation1, err1 := NewSequencePermutation(tt.n1, tt.values1)
			if err1 != nil {
				t.Fatalf("got an error while initializing Permutation 1: %v", err1)
			}

			permutation2, err2 := NewSequencePermutation(tt.n2, tt.values2)
			if err2 != nil {
				t.Fatalf("got an error while initializing Permutation 2: %v", err2)
			}

			_, err := permutation1.Multiply(*permutation2)
			if err == nil {
				t.Fatalf("no error: %v", tt.want)
			}

			if err.Error() != tt.want.Error() {
				t.Errorf("got %v, want %v", err.Error(), tt.want.Error())
			}
		})
	}
}

// Should mutliply permutations. Only sequence permutations.
func TestPermutationsMultiply(t *testing.T) {
	tests := []struct {
		n       int
		values1 []int
		values2 []int
		want    []int
	}{
		{5, []int{3, 4, 1, 5, 2}, []int{5, 3, 1, 2, 4}, []int{2, 1, 3, 4, 5}},
		{5, []int{5, 3, 1, 2, 4}, []int{3, 4, 1, 5, 2}, []int{1, 2, 5, 4, 3}},
		{6, []int{3, 5, 1, 6, 2, 4}, []int{6, 3, 4, 2, 1, 5}, []int{4, 1, 6, 5, 3, 2}},
		{6, []int{6, 3, 4, 2, 1, 5}, []int{3, 5, 1, 6, 2, 4}, []int{4, 1, 6, 5, 3, 2}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%vx%v", tt.values1, tt.values2)
		t.Run(testname, func(t *testing.T) {
			permutation1, err1 := NewSequencePermutation(tt.n, tt.values1)
			if err1 != nil {
				t.Fatalf("got an error while initializing Permutation 1: %v", err1)
			}

			permutation2, err2 := NewSequencePermutation(tt.n, tt.values2)
			if err2 != nil {
				t.Fatalf("got an error while initializing Permutation 2: %v", err2)
			}

			got, err := permutation1.Multiply(*permutation2)
			if err != nil {
				t.Fatalf("got an error while multiplying permutations: %v", err)
			}

			if fmt.Sprintf("%v", got.values) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("got %d, want %d", got.values, tt.want)
			}
		})
	}
}
