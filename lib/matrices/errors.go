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

// Матрица должна быть квадратной.
func NotSquareMatrixError() error {
	return &matrixError{2, "Matrix must be square"}
}

// Матрицы должны быть равны по размеру.
func NotSameSizeError(rows1 int, columns1 int, rows2 int, columns2 int) error {
	return &matrixError{3, fmt.Sprintf("Matrix sizes are not the same %dx%d != %dx%d", rows1, columns1, rows2, columns2)}
}
