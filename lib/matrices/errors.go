package matrices

import "fmt"

type matrixError struct {
	Code    byte
	Message string
}

func (e *matrixError) Error() string {
	return fmt.Sprintf("%s (code: %d)", e.Message, e.Code)
}

// Неправильно заданная матрица.
func InvalidMatrixError(row int) error {
	return &matrixError{
		1, fmt.Sprintf("Invalid matrix at row=%d", row),
	}
}
