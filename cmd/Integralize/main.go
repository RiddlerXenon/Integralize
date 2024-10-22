package main

import (
	"net/http"

	"go.uber.org/zap"
)

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func main() {
	// http.HandleFunc("/api/get", handler.GetHandler(c))

	zap.S().Info("Server starting at http://127.0.0.1:8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		zap.S().Fatal(err)
	}
}
