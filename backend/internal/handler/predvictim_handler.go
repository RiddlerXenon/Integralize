package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RiddlerXenon/Integralize/internal/differential"
	"go.uber.org/zap"
)

func PredVictimHandler(w http.ResponseWriter, r *http.Request) {
	var request predatorVictimRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		zap.S().Error(err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	zap.S().Infof("Request: %+v", request)

	if _, ok := predatorVictim[request.EquationType]; !ok {
		zap.S().Error("invalid equation type")
		http.Error(w, "invalid equation type", http.StatusBadRequest)
		return
	}

	params := differential.Parameters{
		Alpha: request.Args[0],
		Beta:  request.Args[1],
		Gamma: request.Args[2],
		Delta: request.Args[3],
		Step:  request.Step,
		Steps: request.Steps,
		Prey:  request.PredVictim[0],
		Pred:  request.PredVictim[1],
	}

	preyY, predY, err := predatorVictim[request.EquationType](params)
	if err != nil {
		zap.S().Error(err)
		http.Error(w, "failed to process expression", http.StatusBadRequest)
		return
	}

	response := predatorVictimResponse{
		PreyY: preyY,
		PredY: predY,
	}

	zap.S().Infof("Response: %+v", response)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
