package parser

import "math"

func parseStr(str string) interface{} {
	result := 0.0
	operator := '+'
	funcs := map[string]interface{}{
		"sin":  math.Sin,
		"cos":  math.Cos,
		"tan":  math.Tan,
		"exp":  math.Exp,
		"log":  math.Log,
		"sqrt": math.Sqrt,
		"pow":  math.Pow,
	}
	operators := map[byte]bool{
		byte('+'): false,
		byte('-'): false,
		byte('*'): false,
	}

	return func(x float64) float64 {
		s := ""
		for i := 0; i < len(str); i++ {
			if str[i] == ' ' {
				continue
			}

			s += string(str[i])
			if f, ok := funcs[s]; ok {
				switch operator {
				case '+':
					result += f(x)
				case '-':
					result -= f(x)
				}

				s = ""
			} else if _, ok := operations[str[i]]; ok {
				operator = str[i]
			}
		}

		return result
	}
}
