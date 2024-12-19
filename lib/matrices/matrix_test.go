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
		{[][]float64{{2, 5, 4}, {1, 3, 2}, {2, 10, 9}}, 5, nil},
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

// Действия с числами (умножение, деление)
func TestMatrixNumbersOperations(t *testing.T) {
	tests := []struct {
		elements  [][]float64
		number    float64
		operation byte
		want      [][]float64
	}{
		{[][]float64{{12, -1}, {7, 0}}, 3, 0, [][]float64{{36, -3}, {21, 0}}},
		{[][]float64{{36, -3}, {21, 0}}, 3, 1, [][]float64{{12, -1}, {7, 0}}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%f*%v", tt.number, tt.elements)
		t.Run(testname, func(t *testing.T) {
			matrix, err := NewMatrix(tt.elements)
			if err != nil {
				t.Fatalf("got an error while initializing Matrix: %v", err)
			}
			switch tt.operation {
			case 0:
				matrix.MultiplyByNumber(tt.number)
			case 1:
				matrix.DivideByNumber(tt.number)
			}
			got := matrix.elements
			if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

// Транспонирование матрицы
func TestMatrixTranspose(t *testing.T) {
	tests := []struct {
		elements [][]float64
		want     [][]float64
	}{
		{[][]float64{{7, 3, -12, 0, 34}}, [][]float64{{7}, {3}, {-12}, {0}, {34}}},
		{[][]float64{{-1, 0, -2}, {-5, 4, -7}, {6, -4, -6}}, [][]float64{{-1, -5, 6}, {0, 4, -4}, {-2, -7, -6}}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.elements)
		t.Run(testname, func(t *testing.T) {
			matrix, err := NewMatrix(tt.elements)
			if err != nil {
				t.Fatalf("got an error while initializing Matrix: %v", err)
			}
			got := matrix.Transpose().elements
			if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
