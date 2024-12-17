// Пакет matrices предоставляет реализацию алгоритмов матриц:
//   - Создание матрицы
package matrices

// Матрица действительных чисел
type Matrix struct {
	rows     int         // Количество строк
	columns  int         // Количество столбцов
	elements [][]float64 // Элементы матрицы
}

// Возвращает матрицу действительных чисел.
//
// Возвращает ошибку, если в матрице не одинаковое количество столбцов.
func NewMatrix(elements [][]float64) (Matrix, error) {
	// Проверка правильности заданной матрицы
	rows := len(elements)
	columns := -1
	for i := 0; i < rows; i++ {
		if i == 0 {
			columns = len(elements[i])
		} else if len(elements[i]) != columns {
			return Matrix{}, InvalidMatrixError(i + 1)
		}
	}
	return Matrix{rows, columns, elements}, nil
}

// Возвращает true, если матрица является квадратной
// (количество строк совпадает с количеством столбцов).
func (m Matrix) IsSquare() bool {
	return m.rows == m.columns
}
