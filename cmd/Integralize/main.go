package main

import (
	"net/http"
	"os"

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

	certFound, keyFound := true, true
	if _, err := os.Stat(certPath); os.IsNotExist(err) {
		zap.S().Errorf("Certificate file not found: %v", err)
		certFound = false
	}
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		zap.S().Errorf("Key file not found: %v", err)
		keyFound = false
	}

	if certFound && keyFound {
		zap.S().Info("Starting server on :443")
		if err := http.ListenAndServeTLS(":443", certPath, keyPath, HandleWithMiddleware); err != nil {
			zap.S().Fatalf("Failed to start server: %v", err)
		}
	} else {
		zap.S().Info("Starting server on :8080")
		if err := http.ListenAndServe(":8080", HandleWithMiddleware); err != nil {
			zap.S().Fatalf("Failed to start server: %v", err)
		}
	}
}
