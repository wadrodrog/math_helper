// Пакет matrices предоставляет реализацию алгоритмов матриц:
//   - Создание матрицы
//   - Умножение и деление на число
//   - Транспонирование матрицы
//   - Сложение и вычитание матриц
//   - Нахождение определителей 1-3 порядков
//   - Умножение матриц
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

// Прибавляет к каждому элементу матрицы элементы другой матрицы.
//
// Если аргумент negative равен true, то будет произведено вычитание матриц.
func (m *Matrix) AddMatrix(other Matrix, negative bool) error {
	// У матриц должны быть равно количество строк и столбцов
	if m.rows != other.rows || m.columns != other.columns {
		return NotSameSizeError(m.rows, m.columns, other.rows, other.columns)
	}

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			if negative {
				m.elements[i][j] -= other.elements[i][j]
			} else {
				m.elements[i][j] += other.elements[i][j]
			}
		}
	}

	return nil
}

// Возвращает матрицу, являющуюся результатом умножения одной текущей матрицы
// на другую заданную.
//
// Возвращает ошибку, если матрицы нельзя перемножить (количество столбцов
// первой матрицы не количеству строк второй)
func (m Matrix) MultiplyMatrix(other Matrix) (Matrix, error) {
	// Число столбцов первой матрицы должно совпадать с числом строк второй
	if m.columns != other.rows {
		return Matrix{}, UnableToMultiplyError(m.columns, other.rows)
	}

	result := ZeroMatrix(m.rows, other.columns)

	for i := 0; i < result.rows; i++ {
		for j := 0; j < result.columns; j++ {
			for k := 0; k < m.columns; k++ {
				result.elements[i][j] += m.elements[i][k] * other.elements[k][j]
			}
		}
	}

	return result, nil
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
