package handler

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/RiddlerXenon/Integralize/internal/parser"
)

// Хэндлеры для обработки запросов
func IntegralHandler(w http.ResponseWriter, r *http.Request) {
	var request integralRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		zap.S().Error(err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	zap.S().Infof("Request: %+v", request)

	// Парсинг выражения
	expressionFunc, err := parser.PrepareLatexExpression(request.Expression)
	if err != nil {
		zap.S().Error(err)
		http.Error(w, "failed to parse expression", http.StatusBadRequest)
		return
	}

	if _, ok := integralMethods[request.EquationType]; !ok {
		zap.S().Error("invalid equation type")
		http.Error(w, "invalid equation type", http.StatusBadRequest)
		return
	}

	result, err := integralMethods[request.EquationType](request.Args[0], request.Args[1], request.Args[2], expressionFunc)

	if err != nil {
		zap.S().Error(err)
		http.Error(w, "failed to process expression", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := integralResponse{
		Result: result,
	}

	zap.S().Infof("Response: %+v", response)

	json.NewEncoder(w).Encode(response)

	// fmt.Fprintf(w, "Expression processed successfully")
}
