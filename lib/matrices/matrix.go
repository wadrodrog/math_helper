// Пакет matrices предоставляет реализацию алгоритмов матриц:
//   - Создание матрицы
//   - Умножение и деление на число
//   - Транспонирование матрицы
//   - Нахождение определителей 1-3 порядков
package matrices

// Матрица действительных чисел
type Matrix struct {
	rows     int         // Количество строк
	columns  int         // Количество столбцов
	elements [][]float64 // Элементы матрицы
}

// Возвращает нулевую матрицу.
func ZeroMatrix(rows int, columns int) Matrix {
	elements := make([][]float64, rows)
	for i := range elements {
		elements[i] = make([]float64, columns)
	}
	return Matrix{rows, columns, elements}
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

// Умножает каждый элемент матрицы на заданное число.
func (m *Matrix) MultiplyByNumber(number float64) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			m.elements[i][j] *= number
		}
	}
}

// Делит каждый элемент матрицы на заданное число.
func (m *Matrix) DivideByNumber(number float64) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			m.elements[i][j] /= number
		}
	}
}

// Возвращает транспонированную матрицу, то есть матрицу, в которой строки
// записаны как столбцы, а столбцы - как строки.
func (m Matrix) Transpose() Matrix {
	transposed := ZeroMatrix(m.columns, m.rows)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			transposed.elements[j][i] = m.elements[i][j]
		}
	}
	return transposed
}

// Возвращает определитель квадратной матрицы.
//
// Возвращает ошибку, если матрица не квадратная.
func (m *Matrix) Determinator() (float64, error) {
	// Матрица должна быть квадратной
	if m.rows != m.columns {
		return 0, NotSquareMatrixError()
	}

	// Вычисляем определитель
	determinator := 0.0
	switch m.columns {
	case 2:
		determinator = m.elements[0][0]*m.elements[1][1] - m.elements[0][1]*m.elements[1][0]
	case 3:
		determinator = m.elements[0][0]*m.elements[1][1]*m.elements[2][2] + m.elements[2][0]*m.elements[0][1]*m.elements[1][2] + m.elements[0][2]*m.elements[1][0]*m.elements[2][1] - m.elements[0][2]*m.elements[1][1]*m.elements[2][0] - m.elements[0][0]*m.elements[1][2]*m.elements[2][1] - m.elements[2][2]*m.elements[0][1]*m.elements[1][0]
	default:
		determinator = m.elements[0][0]
	}

	return determinator, nil
}
