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

func allNumbersFrom1ToN(n int, slice []int) error {
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		if slice[i] > n || slice[i] < 1 {
			return InvalidElementError(slice[i], n)
		}
		if used[slice[i]-1] {
			return RepeatingElementError(slice[i])
		}
		used[slice[i]-1] = true
	}
	return nil
}

// Permutation function
type Permutation struct {
	size           int         // How many elements this permutation has
	arguments      []int       // Permutation arguments
	values         []int       // Permutation values
	associations   map[int]int // Arguments - values
	inversions     int         // Pairs of positions (i, j) where the entries of a permutation are in the opposite order
	count          uint        // Factorial of size
	cycles         [][]int     // Decomposed permutation into cycles
	transpositions [][]int     // Decomposed permutation into cycles of length 2
}

// Returns Permutation.
//
// Permutation arguments and values are unordered arrays of all numbers
// from 1 to n
//
// Returns error if each slice does not have all numbers from 1 to n.
//
// Example:
//
// (3 2 1)
//
// (2 1 3)
func NewPermutation(n int, arguments []int, values []int) (*Permutation, error) {
	if len(arguments) != n {
		return nil, InvalidLengthError(len(arguments), n)
	}
	if len(values) != n {
		return nil, InvalidLengthError(len(values), n)
	}

	if err := allNumbersFrom1ToN(n, arguments); err != nil {
		return nil, err
	}
	if err := allNumbersFrom1ToN(n, values); err != nil {
		return nil, err
	}

	associations := map[int]int{}
	for i := 0; i < n; i++ {
		associations[arguments[i]] = values[i]
	}

	return &Permutation{n, arguments, values, associations, -1, 0, [][]int{}, [][]int{}}, nil
}

// Returns Permutation.
//
// Permutation arguments is an ordered array of all numbers from 1 to n, defined
// automatically.
//
// Permutation values is an unordered array of all numbers from 1 to n, defined
// manually by passing a slice.
//
// Returns error if values slice does not have all numbers from 1 to n.
//
// Example:
//
// (3 2 1)
//
// (1 2 3)
func NewSequencePermutation(n int, values []int) (*Permutation, error) {
	arguments := make([]int, n)
	for i := 0; i < n; i++ {
		arguments[i] = i + 1
	}
	return NewPermutation(n, arguments, values)
}

// Returns Permutation.
//
// Permutation arguments is an ordered array of all numbers from 1 to n, defined
// automatically.
//
// Permutation values is an unordered array of all numbers from 1 to n, defined
// from transpositions cyctels.
//
// Returns error if it is impossible to construct permutation with numbers
// from 1 to n.
func NewSequencePermutationFromTranspositions(transpositions [][]int) (*Permutation, error) {
	// Find max value
	maxValue := -1
	for i := 0; i < len(transpositions); i++ {
		if len(transpositions[i]) != 2 || transpositions[i][0] < 1 || transpositions[i][1] < 1 {
			return nil, InvalidTranspositionError(i)
		}
		maxValue = max(maxValue, transpositions[i][0], transpositions[i][1])
	}
	if maxValue == -1 {
		return nil, InvalidTranspositionError(0)
	}

	// Create arguments slice using found max value
	arguments := make([]int, maxValue)
	values := make([]int, maxValue)
	for i := 0; i < maxValue; i++ {
		arguments[i] = i + 1
		values[i] = i + 1
	}

	// Recreate values slice using transpositions cycles
	for i := 0; i < len(transpositions); i++ {
		x := transpositions[i][0] - 1
		y := transpositions[i][1] - 1
		temp := values[x]
		values[x] = values[y]
		values[y] = temp
	}

	return NewPermutation(maxValue, arguments, values)
}

// Returns inversions count of current permutation.
//
// Calculates the value for the first time, then caches it.
func (p *Permutation) Inversions() int {
	// Return cached value
	if p.inversions != -1 {
		return p.inversions
	}

	// Calculate for the first time
	p.inversions = 0
	for i := 0; i < p.size; i++ {
		for j := i + 1; j < p.size; j++ {
			if p.values[i] > p.values[j] {
				p.inversions++
			}
		}
	}
	return p.inversions
}

// Returns true if the permutation's transpositions count is even
func (p *Permutation) IsEven() bool {
	return len(p.Transpositions())%2 == 0
}

// Returns permutations count of n elements. It is n!
//
// Calculates the value for the first time, then caches it.
//
// This methods will run slowly and return incorrect result for big values.
func (p *Permutation) Count() uint {
	// Return cached value
	if p.count != 0 {
		return p.count
	}

	// Calculate for the first time
	p.count = 1
	for i := uint(1); i <= uint(p.size); i++ {
		p.count *= i
	}
	return p.count
}

// Returns permutation decomposition into cycles.
//
// Calculates the value for the first time, then caches it.
func (p *Permutation) Cycles() [][]int {
	// Return cached value
	if len(p.cycles) != 0 {
		return p.cycles
	}

	used := make([]bool, p.size)

	// Build cycles
	for i := 0; i < p.size; i++ {
		// Shortcuts
		arg := p.arguments[i]
		val := p.associations[arg]

		// Do not start a cycle if initial value already has been written
		if used[arg-1] {
			continue
		}

		// This will be the cycle of length 1. They are usually omitted
		if arg == val {
			continue
		}

		// Otherwise, start a cycle with first two values
		cycle := []int{arg, val}
		used[arg-1] = true
		used[val-1] = true

		// Build a cycle step-by-step
		// Every step is value of previous as index
		// Repeat until the value does not equal initial value
		for j := p.associations[val]; j != p.arguments[i]; j = p.associations[j] {
			cycle = append(cycle, j)
			used[j-1] = true
		}

		// Write a cycle
		p.cycles = append(p.cycles, cycle)
	}

	return p.cycles
}

// Returns permutation decomposition into transpositions cycles.
//
// Calculates the value for the first time, then caches it.
func (p *Permutation) Transpositions() [][]int {
	// Return cached value
	if len(p.transpositions) != 0 {
		return p.transpositions
	}

	// Calculate cycles
	p.Cycles()

	// Build transpositions cycles
	for i := 0; i < len(p.cycles); i++ {
		for j := 1; j < len(p.cycles[i]); j++ {
			p.transpositions = append(p.transpositions, []int{p.cycles[i][j-1], p.cycles[i][j]})
		}
	}

	return p.transpositions
}

// Returns permutations composition. Only sequence permutations supported.
//
// Returns error if permutations lengths does not equal
func (p1 Permutation) Multiply(p2 Permutation) (*Permutation, error) {
	// Permutations length should be equal
	if p1.size != p2.size {
		return nil, InvalidLengthError(p2.size, p1.size)
	}

	// Calculate new permutation
	values := make([]int, p1.size)
	for i := 0; i < p2.size; i++ {
		val := p2.associations[p2.arguments[i]]
		values[i] = p1.associations[val]
	}

	return NewPermutation(p2.size, p2.arguments, values)
}
