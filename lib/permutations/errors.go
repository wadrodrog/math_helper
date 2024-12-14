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

// Неправильная длина перестановки.
func InvalidLengthError(got int, expected int) error {
	return &permutationError{
		1, fmt.Sprintf("Invalid length: permutation length is %d, but n=%d", got, expected),
	}
}

// Неправильный элемент перестановки.
func InvalidElementError(element int, n int) error {
	return &permutationError{
		2, fmt.Sprintf("Element %d does not belong the range 1..%d", element, n),
	}
}

// Элемент перестановки повторяется.
func RepeatingElementError(element int) error {
	return &permutationError{
		3, fmt.Sprintf("Repeating element: %d", element),
	}
}

// Неправильные транспозиции.
func InvalidTranspositionError(position int) error {
	return &permutationError{
		4, fmt.Sprintf("Invalid transposition at position %d", position),
	}
}
