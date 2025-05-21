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

	certPath := "/etc/letsencrypt/live/integralize.ru/fullchain.pem"
	keyPath := "/etc/letsencrypt/live/integralize.ru/privkey.pem"

	zap.S().Info("Starting server on :443")
	if err := http.ListenAndServeTLS(":443", certPath, keyPath, HandleWithMiddleware); err != nil {
		zap.S().Fatalf("Failed to start server: %v", err)
	}
}
