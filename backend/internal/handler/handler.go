package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RiddlerXenon/Integralize/internal/parser"
)

// Хэндлеры для обработки запросов
func IntegralHandler(w http.ResponseWriter, r *http.Request) {
	var request integralRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// Парсинг выражения
	expressionFunc, err := parser.ParseStrInt(request.Expression)
	if err != nil {
		http.Error(w, "failed to parse expression", http.StatusBadRequest)
		return
	}

	if _, ok := integralMethods[request.EquationType]; !ok {
		http.Error(w, "invalid equation type", http.StatusBadRequest)
		return
	}

	result, err := integralMethods[request.EquationType](request.Args[0], request.Args[1], request.Args[2], expressionFunc)

	if err != nil {
		http.Error(w, "failed to process expression", http.StatusBadRequest)
		return
	}

	response := integralResponse{
		Result: result,
	}

	json.NewEncoder(w).Encode(response)

	fmt.Fprintf(w, "Expression processed successfully")
}

func DiffEquationsHandler(w http.ResponseWriter, r *http.Request) {
	var request diffEquationsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// Парсинг выражения
	expressionFunc, err := parser.ParseStrDiffEq(request.Expression)
	if err != nil {
		http.Error(w, "failed to parse expression", http.StatusBadRequest)
		return
	}

	if _, ok := diffEquationsMethods[request.EquationType]; !ok {
		http.Error(w, "invalid equation type", http.StatusBadRequest)
		return
	}

	x, y := diffEquationsMethods[request.EquationType](request.Args[0], request.Args[1], request.Args[2], request.Args[3], expressionFunc)

	response := diffEquationsResponse{
		X: x,
		Y: y,
	}

	json.NewEncoder(w).Encode(response)
}
