package handler

import (
	"encoding/json"
	"math"
	"net/http"

	"github.com/RiddlerXenon/Integralize/internal/parser"
	"go.uber.org/zap"
)

func sanitizeFloatArray(arr []float64) []interface{} {
	out := make([]interface{}, len(arr))
	for i, v := range arr {
		switch {
		case math.IsInf(v, 1):
			out[i] = "+Inf"
		case math.IsInf(v, -1):
			out[i] = "-Inf"
		case math.IsNaN(v):
			out[i] = "NaN"
		default:
			out[i] = v
		}
	}
	return out
}

func DiffEquationsHandler(w http.ResponseWriter, r *http.Request) {
	var request diffEquationsRequest
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

	if _, ok := diffEquationsMethods[request.EquationType]; !ok {
		zap.S().Error("invalid equation type")
		http.Error(w, "invalid equation type", http.StatusBadRequest)
		return
	}

	x, y := diffEquationsMethods[request.EquationType](request.Args[0], request.Args[1], request.Args[2], request.Args[3], expressionFunc)

	response := map[string]interface{}{
		"x": sanitizeFloatArray(x),
		"y": sanitizeFloatArray(y),
	}

	zap.S().Infof("Response: %+v", response)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
