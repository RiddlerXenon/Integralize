package methods

import (
	"log"

	"github.com/Knetic/govaluate"
)

// FunctionParser парсит введенную пользователем строку и создает функцию для интегрирования
type FunctionParser struct {
	expression *govaluate.EvaluableExpression
}

// NewFunctionParser инициализирует новый парсер для пользовательской функции
func NewFunctionParser(functionStr string) *FunctionParser {
	expr, err := govaluate.NewEvaluableExpression(functionStr)
	if err != nil {
		log.Fatalf("Ошибка парсинга выражения: %v", err)
	}
	return &FunctionParser{
		expression: expr,
	}
}

// F вычисляет значение пользовательской функции в точке x
func (fp *FunctionParser) F(x float64) float64 {
	parameters := map[string]interface{}{
		"x": x,
	}
	result, err := fp.expression.Evaluate(parameters)
	if err != nil {
		log.Fatalf("Ошибка вычисления функции: %v", err)
	}
	return result.(float64)
}
