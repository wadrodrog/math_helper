// Пакет permutations предоставляет реализацию алгоритмов перестановок:
//   - Создание перестановки
//   - Подсчёт количества перестановок
//   - Подсчёт количества инверсий
//   - Определение чётности перестановки
//   - Разложение на циклы
//   - Разложение на транспозиции
//   - Сборка перестановки из транспозиции
//   - Умножение перестановок
package permutations

func allNumbersFrom1ToN(n int, slice []int) error {
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		if slice[i] > n || slice[i] < 1 {
			return InvalidElementError(slice[i], n)
		}
		if used[slice[i]-1] {
			return RepeatingElementError(slice[i])
		}
		used[slice[i]-1] = true
	}
	return nil
}

// Permutation представляет собой перестановку.
type Permutation struct {
	size           int         // Количество элементов в перестановке (n)
	arguments      []int       // Аргументы перестановки (верхний ряд чисел)
	values         []int       // Значения перестановки (нижний ряд чисел)
	associations   map[int]int // Ассоциации: Аргумент - Значение
	inversions     int         // Инверсии: пары индексов (i, j), при которых элементы перестановки расположены в обратном порядке.
	count          uint        // Количество перестановок (факториал количества элементов)
	cycles         [][]int     // Разложение перестановки на циклы
	transpositions [][]int     // Разложение перестановки на транспозиции (циклы длины 2)
}

// Возвращает перестановку. Перестановка задаётся двумя массивами чисел
// от 1 до n.
//
// Возвращает ошибку, если в массивах нет всех чисел от 1 до n.
//
// Пример:
//
//	(2 1 3)
//	(3 2 1)
func NewPermutation(n int, arguments []int, values []int) (*Permutation, error) {
	if len(arguments) != n {
		return nil, InvalidLengthError(len(arguments), n)
	}
	if len(values) != n {
		return nil, InvalidLengthError(len(values), n)
	}

	if err := allNumbersFrom1ToN(n, arguments); err != nil {
		return nil, err
	}
	if err := allNumbersFrom1ToN(n, values); err != nil {
		return nil, err
	}

	associations := map[int]int{}
	for i := 0; i < n; i++ {
		associations[arguments[i]] = values[i]
	}

	return &Permutation{n, arguments, values, associations, -1, 0, [][]int{}, [][]int{}}, nil
}

// Возвращает перестановку. Перестановка задаётся массивом чисел от 1 до n.
// Аргументы перестановки (верхний ряд чисел) упорядочен.
//
// Возвращает ошибку, если в массиве нет всех чисел от 1 до n.
//
// Пример:
//
//	(1 2 3)
//	(2 1 3)
func NewSequencePermutation(n int, values []int) (*Permutation, error) {
	arguments := make([]int, n)
	for i := 0; i < n; i++ {
		arguments[i] = i + 1
	}
	return NewPermutation(n, arguments, values)
}

// Возвращает перестановку. Перестановка составляется из транспозиций.
// Аргументы перестановки (верхний ряд чисел) упорядочен.
//
// Возвращает ошибку, если невозможно составить перестановку из заданных
// транспозиций.
//
// Пример:
//
//	(1 3)(3 2)(2 1) => (1 3 2)
func NewSequencePermutationFromTranspositions(transpositions [][]int) (*Permutation, error) {
	// Находим максимальное значение
	maxValue := -1
	for i := 0; i < len(transpositions); i++ {
		if len(transpositions[i]) != 2 || transpositions[i][0] < 1 || transpositions[i][1] < 1 {
			return nil, InvalidTranspositionError(i)
		}
		maxValue = max(maxValue, transpositions[i][0], transpositions[i][1])
	}
	if maxValue == -1 {
		return nil, InvalidTranspositionError(0)
	}

	// Создаём массив аргументов, используя найденное максимальное значение
	arguments := make([]int, maxValue)
	values := make([]int, maxValue)
	for i := 0; i < maxValue; i++ {
		arguments[i] = i + 1
		values[i] = i + 1
	}

	// Восстанавливаем массив значений, меняя числа местами согласно транспозициям
	for i := 0; i < len(transpositions); i++ {
		x := transpositions[i][0] - 1
		y := transpositions[i][1] - 1
		temp := values[x]
		values[x] = values[y]
		values[y] = temp
	}

	return NewPermutation(maxValue, arguments, values)
}

// Возвращает количество инверсий перестановки.
func (p *Permutation) Inversions() int {
	// Возващаем кэшированное значение
	if p.inversions != -1 {
		return p.inversions
	}

	// Вычисляем значение в первый раз
	p.inversions = 0
	for i := 0; i < p.size; i++ {
		for j := i + 1; j < p.size; j++ {
			if p.values[i] > p.values[j] {
				p.inversions++
			}
		}
	}
	return p.inversions
}

// Возвращает true, если количество транспозиций перестановки чётно.
func (p *Permutation) IsEven() bool {
	return len(p.Transpositions())%2 == 0
}

// Возвращает количество перестановок n элементов.
//
// Вычисляет факториал из n медленно и неправильно для больших значений.
func (p *Permutation) Count() uint {
	// Возващаем кэшированное значение
	if p.count != 0 {
		return p.count
	}

	// Вычисляем значение в первый раз
	p.count = 1
	for i := uint(1); i <= uint(p.size); i++ {
		p.count *= i
	}
	return p.count
}

// Возвращает разложение перестановки на циклы.
func (p *Permutation) Cycles() [][]int {
	// Возващаем кэшированное значение
	if len(p.cycles) != 0 {
		return p.cycles
	}

	used := make([]bool, p.size)

	// Строим циклы
	for i := 0; i < p.size; i++ {
		// Сокращения
		arg := p.arguments[i]
		val := p.associations[arg]

		// Не начинать цикл, если начальное значение уже было использовано
		if used[arg-1] {
			continue
		}

		// Это будет цикл длины 1. Обычно их не записывают
		if arg == val {
			continue
		}

		// Иначе начинаем цикл с первыми двумя значениями
		cycle := []int{arg, val}
		used[arg-1] = true
		used[val-1] = true

		// Пошаговое построение цикла.
		// Каждый шаг - это значение от предыдущего аргумента.
		// Повторять, пока значение не равно первоначальному.
		for j := p.associations[val]; j != p.arguments[i]; j = p.associations[j] {
			cycle = append(cycle, j)
			used[j-1] = true
		}

		// Записать цикл
		p.cycles = append(p.cycles, cycle)
	}

	return p.cycles
}

// Возвращает разложение перестановки на транспозиции.
func (p *Permutation) Transpositions() [][]int {
	// Возващаем кэшированное значение
	if len(p.transpositions) != 0 {
		return p.transpositions
	}

	// Сначала нужно разложить на циклы
	p.Cycles()

	// Строим транспозиции
	for i := 0; i < len(p.cycles); i++ {
		for j := 1; j < len(p.cycles[i]); j++ {
			p.transpositions = append(p.transpositions, []int{p.cycles[i][j-1], p.cycles[i][j]})
		}
	}

	return p.transpositions
}

// Возвращает композицию перестановок. Это умножение текущей перестановки на
// другую.
//
// Возвращает ошибку, если длины перестановок не равны.
func (p1 Permutation) Multiply(p2 Permutation) (*Permutation, error) {
	// Проверка длины
	if p1.size != p2.size {
		return nil, InvalidLengthError(p2.size, p1.size)
	}

	// Вычислить новую перестановку
	values := make([]int, p1.size)
	for i := 0; i < p2.size; i++ {
		val := p2.associations[p2.arguments[i]]
		values[i] = p1.associations[val]
	}

	return NewPermutation(p2.size, p2.arguments, values)
}
