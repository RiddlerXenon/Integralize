package main

import (
	"net/http"

	"github.com/RiddlerXenon/Integralize/internal/handler"
	"go.uber.org/zap"
)

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

// enableCors добавляет заголовки CORS к ответу
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func main() {
	http.HandleFunc("/api/integral", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		// Если метод OPTIONS, просто верните 200 OK
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		handler.IntegralHandler(w, r)
	})
	http.HandleFunc("/api/differential", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		// Если метод OPTIONS, просто верните 200 OK
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		handler.DiffEquationsHandler(w, r)
	})

	http.HandleFunc("/api/predvictim", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		// Если метод OPTIONS, просто верните 200 OK
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		handler.PredVictimHandler(w, r)
	})

	zap.S().Info("Server starting at http://127.0.0.1:8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		zap.S().Fatal(err)
	}
}
