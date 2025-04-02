package parser

import (
	"math"
	"strings"

	"github.com/Knetic/govaluate"
)

// Prepare a LaTeX expression for evaluation
func PrepareLatexExpression(latex string) (func(map[string]float64) float64, error) {
	// Convert LaTeX operators and functions to Go operators and functions
	replacements := map[string]string{
		"\\cdot":   "*",
		"\\div":    "/",
		"\\left(":  "(",
		"\\right)": ")",
		"\\sin":    "sin",
		"\\cos":    "cos",
		"\\tan":    "tan",
		"\\log":    "log",
		"\\ln":     "ln",
		"{":        "",
		"}":        "",
		"^":        "**",
	}

	for latexOp, goOp := range replacements {
		latex = strings.ReplaceAll(latex, latexOp, goOp)
	}

	// Remove LaTeX spaces
	latex = strings.ReplaceAll(latex, "\\,", "")

	// Create a new govaluate expression
	functions := map[string]govaluate.ExpressionFunction{
		"sin": func(args ...interface{}) (interface{}, error) {
			return math.Sin(args[0].(float64)), nil
		},
		"cos": func(args ...interface{}) (interface{}, error) {
			return math.Cos(args[0].(float64)), nil
		},
		"tan": func(args ...interface{}) (interface{}, error) {
			return math.Tan(args[0].(float64)), nil
		},
		"log": func(args ...interface{}) (interface{}, error) {
			return math.Log10(args[0].(float64)), nil
		},
		"ln": func(args ...interface{}) (interface{}, error) {
			return math.Log(args[0].(float64)), nil
		},
		"**": func(args ...interface{}) (interface{}, error) {
			base := args[0].(float64)
			exponent := args[1].(float64)
			return math.Pow(base, exponent), nil
		},
	}

	expression, err := govaluate.NewEvaluableExpressionWithFunctions(latex, functions)
	if err != nil {
		return nil, err
	}

	// Return a function that takes variables and evaluates the expression
	return func(variables map[string]float64) float64 {
		// Prepare parameters for the expression
		parameters := make(map[string]interface{})
		for k, v := range variables {
			parameters[k] = v
		}

		// Evaluate the expression
		result, err := expression.Evaluate(parameters)
		if err != nil {
			return 0
		}

		// Convert the result to float64
		return result.(float64)
	}, nil
}
