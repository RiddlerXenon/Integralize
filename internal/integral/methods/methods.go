package methods

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"unicode"
)

// Поддерживаемые функции
var functions = map[string]string{
	"sin":  "math.Sin",
	"cos":  "math.Cos",
	"tan":  "math.Tan",
	"exp":  "math.Exp",
	"log":  "math.Log",
	"sqrt": "math.Sqrt",
	"pow":  "math.Pow",
}

// Определение токенов
type TokenType int

const (
	Operator TokenType = iota
	Function
	Variable
	Number
	LParen
	RParen
)

// Токен
type Token struct {
	value string
	typ   TokenType
}

// Лексический анализ
func tokenize(expr string) ([]Token, error) {
	tokens := []Token{}
	current := strings.Builder{}

	for i := 0; i < len(expr); i++ {
		char := expr[i]

		// Игнорируем пробелы
		if unicode.IsSpace(rune(char)) {
			continue
		}

		if unicode.IsLetter(rune(char)) {
			current.WriteByte(char)
			for i+1 < len(expr) && unicode.IsLetter(rune(expr[i+1])) {
				i++
				current.WriteByte(expr[i])
			}
			funcName := current.String()
			if _, ok := functions[funcName]; ok {
				tokens = append(tokens, Token{value: funcName, typ: Function})
			} else {
				tokens = append(tokens, Token{value: funcName, typ: Variable}) // обработка переменных
			}
			current.Reset()
		} else if unicode.IsDigit(rune(char)) || char == '.' {
			current.WriteByte(char)
			for i+1 < len(expr) && (unicode.IsDigit(rune(expr[i+1])) || expr[i+1] == '.') {
				i++
				current.WriteByte(expr[i])
			}
			tokens = append(tokens, Token{value: current.String(), typ: Number})
			current.Reset()
		} else {
			switch char {
			case '+', '-', '*', '/', '^':
				tokens = append(tokens, Token{value: string(char), typ: Operator})
			case '(':
				tokens = append(tokens, Token{value: string(char), typ: LParen})
			case ')':
				tokens = append(tokens, Token{value: string(char), typ: RParen})
			default:
				return nil, errors.New("недопустимый символ: " + string(char))
			}
		}
	}
	return tokens, nil
}

// Преобразование токенов в математическое выражение
func evalTokens(tokens []Token, x float64) (float64, error) {
	stack := []float64{}
	operators := []string{}

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		switch token.typ {
		case Number:
			val, _ := strconv.ParseFloat(token.value, 64)
			stack = append(stack, val)
		case Variable:
			// Замена переменной x на её значение
			if token.value == "x" {
				stack = append(stack, x)
			} else {
				return 0, errors.New("неизвестная переменная: " + token.value)
			}
		case Function:
			// Применяем функцию к значению
			if i+1 < len(tokens) && tokens[i+1].typ == LParen {
				// Пропускаем до аргументов функции
				operators = append(operators, token.value)
			}
		case LParen:
			// Ожидаем открывающую скобку и пропускаем
		case RParen:
			// Когда закрывающая скобка, выполняем функцию
			if len(operators) == 0 {
				return 0, errors.New("нет операторов для выполнения функции")
			}
			lastFunc := operators[len(operators)-1]
			operators = operators[:len(operators)-1]
			arg := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			switch lastFunc {
			case "sin":
				stack = append(stack, math.Sin(arg))
			case "cos":
				stack = append(stack, math.Cos(arg))
			case "tan":
				stack = append(stack, math.Tan(arg))
			case "exp":
				stack = append(stack, math.Exp(arg))
			case "log":
				stack = append(stack, math.Log(arg))
			case "sqrt":
				stack = append(stack, math.Sqrt(arg))
			}
		case Operator:
			// Операторы: +, -, *, /, ^
			if len(stack) < 2 {
				return 0, errors.New("недостаточно аргументов для оператора " + token.value)
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch token.value {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				stack = append(stack, a/b)
			case "^":
				stack = append(stack, math.Pow(a, b))
			}
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("ошибка в синтаксисе выражения")
	}
	return stack[0], nil
}

// Функция f для вычисления значения выражения
func f(x float64, expr string) (float64, error) {
	tokens, err := tokenize(expr)
	if err != nil {
		return 0, err
	}

	result, err := evalTokens(tokens, x)
	if err != nil {
		return 0, err
	}

	return result, nil
}
