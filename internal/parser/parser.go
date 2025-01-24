package parser

import (
	"errors"
	"math"
	"strconv"
)

func ParseStrInt(str string) (func(float64) float64, error) {
	funcs := map[string]func(float64) float64{
		"sin":  math.Sin,
		"cos":  math.Cos,
		"tan":  math.Tan,
		"exp":  math.Exp,
		"ln":   math.Log,
		"sqrt": math.Sqrt,
	}

	// Функция для выполнения операций на стеке
	applyOperator := func(stack *[]float64, operator byte) error {
		if len(*stack) < 2 {
			return errors.New("недостаточно операндов")
		}
		b, a := (*stack)[len(*stack)-1], (*stack)[len(*stack)-2]
		*stack = (*stack)[:len(*stack)-2]

		switch operator {
		case '+':
			*stack = append(*stack, a+b)
		case '-':
			*stack = append(*stack, a-b)
		case '*':
			*stack = append(*stack, a*b)
		case '/':
			*stack = append(*stack, a/b)
		case '^':
			*stack = append(*stack, math.Pow(a, b))
		}
		return nil
	}

	// Основная функция для вычисления
	return func(x float64) float64 {
		var (
			stack     []float64
			operators []byte
			applyFunc func(float64) float64
		)

		push := func(val float64) {
			if applyFunc != nil {
				val = applyFunc(val)
				applyFunc = nil
			}
			stack = append(stack, val)
		}

		for i := 0; i < len(str); i++ {
			char := str[i]

			// Пропуск пробелов
			if char == ' ' {
				continue
			}

			// Если встречаем оператор
			if char == '+' || char == '-' || char == '*' || char == '/' || char == '^' {
				// Проверка приоритета операторов
				for len(operators) > 0 && getPrecedence(operators[len(operators)-1]) >= getPrecedence(char) {
					applyOperator(&stack, operators[len(operators)-1])
					operators = operators[:len(operators)-1]
				}
				operators = append(operators, char)
				continue
			}

			// Обработка переменной x
			if char == 'X' {
				push(x)
				continue
			}

			// Обработка функций
			if char >= 'a' && char <= 'z' {
				j := i
				for j < len(str) && str[j] >= 'a' && str[j] <= 'z' {
					j++
				}
				if f, ok := funcs[str[i:j]]; ok {
					applyFunc = f
					i = j - 1
				}
				continue
			}

			// Обработка чисел
			if (char >= '0' && char <= '9') || char == '.' {
				j := i
				for j < len(str) && ((str[j] >= '0' && str[j] <= '9') || str[j] == '.') {
					j++
				}
				num, _ := strconv.ParseFloat(str[i:j], 64)
				push(num)
				i = j - 1
				continue
			}

			// Обработка подвыражений в скобках
			if char == '(' {
				j, depth := i+1, 1
				for j < len(str) && depth > 0 {
					if str[j] == '(' {
						depth++
					} else if str[j] == ')' {
						depth--
					}
					j++
				}
				subFunc, _ := ParseStrInt(str[i+1 : j-1])
				push(subFunc(x))
				i = j - 1
			}
		}

		// Применение оставшихся операторов
		for len(operators) > 0 {
			applyOperator(&stack, operators[len(operators)-1])
			operators = operators[:len(operators)-1]
		}

		if len(stack) == 1 {
			return stack[0]
		}
		return 0
	}, nil
}

func ParseStrDiffEq(str string) (func(float64, float64) float64, error) {
	funcs := map[string]func(float64) float64{
		"sin":  math.Sin,
		"cos":  math.Cos,
		"tan":  math.Tan,
		"exp":  math.Exp,
		"ln":   math.Log,
		"sqrt": math.Sqrt,
	}

	// Функция для выполнения операций на стеке
	applyOperator := func(stack *[]float64, operator byte) error {
		if len(*stack) < 2 {
			return errors.New("недостаточно операндов")
		}
		b, a := (*stack)[len(*stack)-1], (*stack)[len(*stack)-2]
		*stack = (*stack)[:len(*stack)-2]

		switch operator {
		case '+':
			*stack = append(*stack, a+b)
		case '-':
			*stack = append(*stack, a-b)
		case '*':
			*stack = append(*stack, a*b)
		case '/':
			*stack = append(*stack, a/b)
		case '^':
			*stack = append(*stack, math.Pow(a, b))
		}
		return nil
	}

	// Основная функция для вычисления
	return func(x float64, y float64) float64 {
		var (
			stack     []float64
			operators []byte
			applyFunc func(float64) float64
		)

		push := func(val float64) {
			if applyFunc != nil {
				val = applyFunc(val)
				applyFunc = nil
			}
			stack = append(stack, val)
		}

		for i := 0; i < len(str); i++ {
			char := str[i]

			// Пропуск пробелов
			if char == ' ' {
				continue
			}

			// Если встречаем оператор
			if char == '+' || char == '-' || char == '*' || char == '/' || char == '^' {
				// Проверка приоритета операторов
				for len(operators) > 0 && getPrecedence(operators[len(operators)-1]) >= getPrecedence(char) {
					applyOperator(&stack, operators[len(operators)-1])
					operators = operators[:len(operators)-1]
				}
				operators = append(operators, char)
				continue
			}

			// Обработка переменной x
			if char == 'X' {
				push(x)
				continue
			}

			if char == 'Y' {
				push(y)
				continue
			}

			// Обработка функций
			if char >= 'a' && char <= 'z' {
				j := i
				for j < len(str) && str[j] >= 'a' && str[j] <= 'z' {
					j++
				}
				if f, ok := funcs[str[i:j]]; ok {
					applyFunc = f
					i = j - 1
				}
				continue
			}

			// Обработка чисел
			if (char >= '0' && char <= '9') || char == '.' {
				j := i
				for j < len(str) && ((str[j] >= '0' && str[j] <= '9') || str[j] == '.') {
					j++
				}
				num, _ := strconv.ParseFloat(str[i:j], 64)
				push(num)
				i = j - 1
				continue
			}

			// Обработка подвыражений в скобках
			if char == '(' {
				j, depth := i+1, 1
				for j < len(str) && depth > 0 {
					if str[j] == '(' {
						depth++
					} else if str[j] == ')' {
						depth--
					}
					j++
				}
				subFunc, _ := ParseStrDiffEq(str[i+1 : j-1])
				push(subFunc(x, y))
				i = j - 1
			}
		}

		// Применение оставшихся операторов
		for len(operators) > 0 {
			applyOperator(&stack, operators[len(operators)-1])
			operators = operators[:len(operators)-1]
		}

		if len(stack) == 1 {
			return stack[0]
		}
		return 0
	}, nil
}

// getPrecedence возвращает приоритет оператора
func getPrecedence(op byte) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	case '^':
		return 3 // высокий приоритет для степени
	default:
		return 0
	}
}

// func main() {
// 	numRuns := 100
// 	totalDuration := time.Duration(0)

// 	var startMem, endMem runtime.MemStats
// 	runtime.ReadMemStats(&startMem)

// 	for i := 0; i < numRuns; i++ {
// 		start := time.Now()
// 		exprFunc, err := ParseStr("2^3/1^4")
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			return
// 		}
// 		fmt.Println(exprFunc(0.5))
// 		totalDuration += time.Since(start)
// 	}
// 	runtime.ReadMemStats(&endMem)

// 	avgDuration := totalDuration / time.Duration(numRuns)
// 	fmt.Println("Среднее время выполнения:", avgDuration.Nanoseconds())

// 	memoryUsage := endMem.Alloc - startMem.Alloc
// 	fmt.Println("Используемая память (в среднем за тест):", memoryUsage, "байт")
// }
