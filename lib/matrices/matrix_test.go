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

// Проверка квадратной матрицы
func TestMatrixIsSquare(t *testing.T) {
	tests := []struct {
		elements [][]float64
		want     bool
	}{
		{[][]float64{{1, 2, 3}, {4, 5, 6}}, false}, // Не квадратная
		{[][]float64{{1, 2}, {3, 4}}, true},        // Квадратная
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.elements)
		t.Run(testname, func(t *testing.T) {
			matrix, err := NewMatrix(tt.elements)
			if err != nil {
				t.Fatalf("got an error while initializing Matrix: %v", err)
			}
			got := matrix.IsSquare()
			if got != tt.want {
				t.Errorf("got %v, want %v", err, tt.want)
			}
		})
	}
}

// Проверка вычисления определителя матрицы
func TestMatrixDeterminator(t *testing.T) {
	tests := []struct {
		elements [][]float64
		want     float64
		wantErr  error
	}{
		// Ошибка
		{[][]float64{{1, 2, 3}, {4, 5, 6}}, 0, NotSquareMatrixError()},
		// Нет ошибок
		{[][]float64{{5}}, 5, nil},
		{[][]float64{{11, -3}, {-15, -2}}, -67, nil},
		{[][]float64{{1, -2, 3}, {4, 0, 6}, {-7, 8, 9}}, 204, nil},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.elements)
		t.Run(testname, func(t *testing.T) {
			matrix, err := NewMatrix(tt.elements)
			if err != nil {
				t.Fatalf("got an error while initializing Matrix: %v", err)
			}
			got, err := matrix.Determinator()
			if err != nil && tt.wantErr == nil {
				t.Fatalf("got an error while calculating Matrix Determinator: %v", err)
			}
			if err == nil && tt.wantErr != nil {
				t.Fatalf("no error %q", tt.wantErr)
			}
			if err == nil && got != tt.want {
				t.Errorf("got %f, want %f", got, tt.want)
			}
		})
	}
}
