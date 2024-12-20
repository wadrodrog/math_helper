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

// Сложение и вычитание матриц
func TestMatrixAddAndSubstract(t *testing.T) {
	tests := []struct {
		elements1 [][]float64
		elements2 [][]float64
		negative  bool
		want      [][]float64
		wantErr   error
	}{
		{[][]float64{{12, -1}, {7, 0}}, [][]float64{{-1, 0, -2}, {-5, 4, -7}, {6, -4, -6}}, false, [][]float64{{12, -1}, {7, 0}}, NotSameSizeError(2, 2, 3, 3)},
		{[][]float64{{12, -1}, {-5, 0}}, [][]float64{{-4, -3}, {15, 7}}, false, [][]float64{{8, -4}, {10, 7}}, nil},
		{[][]float64{{3, 5, -17}, {-1, 0, 10}}, [][]float64{{-4, 3, -15}, {-5, -7, 0}}, true, [][]float64{{7, 2, -2}, {4, 7, 10}}, nil},
	}

	for _, tt := range tests {
		sign := "+"
		if tt.negative {
			sign = "-"
		}
		testname := fmt.Sprintf("%v%s%v", tt.elements1, sign, tt.elements2)
		t.Run(testname, func(t *testing.T) {
			matrix1, err := NewMatrix(tt.elements1)
			if err != nil {
				t.Fatalf("got an error while initializing Matrix 1: %v", err)
			}
			matrix2, err := NewMatrix(tt.elements2)
			if err != nil {
				t.Fatalf("got an error while initializing Matrix 2: %v", err)
			}
			got := matrix1.AddMatrix(matrix2, tt.negative)
			if got != nil && tt.wantErr == nil {
				t.Fatalf("got an error while adding Matrix: %v", err)
			}
			if got != nil && got.Error() != tt.wantErr.Error() {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
			if fmt.Sprintf("%v", matrix1.elements) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("got %v, want %v", matrix1.elements, tt.want)
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
