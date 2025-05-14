package main

import (
	"net/http"

	"github.com/RiddlerXenon/Integralize/internal/middleware"
	"github.com/RiddlerXenon/Integralize/internal/routes"
	"go.uber.org/zap"
)

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func main() {
	mux := http.NewServeMux()

	// Register routes
	routes.RegisterRoutes(mux)

	// Apply middleware
	HandleWithMiddleware := middleware.CORS(mux)

	zap.S().Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", HandleWithMiddleware); err != nil {
		zap.S().Fatalf("Failed to start server: %v", err)
	}
}
