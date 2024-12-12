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
)

type permutationError struct {
	Code    byte
	Message string
}

func (e *permutationError) Error() string {
	return fmt.Sprintf("%s (code: %d)", e.Message, e.Code)
}

func InvalidLengthError(got int, expected int) error {
	return &permutationError{
		1, fmt.Sprintf("Invalid length: permutation length is %d, but n=%d", got, expected),
	}
}

func InvalidElementError(element int, n int) error {
	return &permutationError{
		2, fmt.Sprintf("Element %d does not belong the range 1..%d", element, n),
	}
}

func RepeatingElementError(element int) error {
	return &permutationError{
		3, fmt.Sprintf("Repeating element: %d", element),
	}
}

func InvalidTranspositionError(position int) error {
	return &permutationError{
		4, fmt.Sprintf("Invalid transposition at position %d", position),
	}
}
