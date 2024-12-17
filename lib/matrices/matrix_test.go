package matrices

import (
	"fmt"
	"testing"
)

// Создание новой матрицы
func TestNewMatrix(t *testing.T) {
	tests := []struct {
		elements [][]float64
		want     error
	}{
		// Ошибка
		{[][]float64{{1, 2, 3}, {4, 5}}, InvalidMatrixError(2)},
		// Нет ошибки
		{[][]float64{{1, 2, 3}, {4, 5, 6}}, nil},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.elements)
		t.Run(testname, func(t *testing.T) {
			_, err := NewMatrix(tt.elements)
			if err == nil && tt.want != nil {
				t.Fatalf("no error %q", tt.want)
			}
			if err != nil && err.Error() != tt.want.Error() {
				t.Errorf("got %q, want %q", err, tt.want)
			}
		})
	}

}
