package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

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

func serveFrontendFile(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join("../frontend", name) // относительно backend/
		// Установим корректный Content-Type для html
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeFile(w, r, path)
	}
}

func serveStaticFile(name string, contentType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join("../frontend", name)
		w.Header().Set("Content-Type", contentType)
		http.ServeFile(w, r, path)
	}
}

func main() {
	http.HandleFunc("/", serveFrontendFile("index.html"))
	http.HandleFunc("/integrals", serveFrontendFile("integrals.html"))
	http.HandleFunc("/diffeq", serveFrontendFile("diffeq.html"))
	http.HandleFunc("/models", serveFrontendFile("models.html"))
	http.HandleFunc("/style.css", serveStaticFile("style.css", "text/css; charset=utf-8"))

	// Статика: assets (картинки, иконки и т.д.)
	assetsDir := filepath.Join("../frontend/assets")
	if _, err := os.Stat(assetsDir); os.IsNotExist(err) {
		log.Fatalf("Директория assets не найдена по пути: %s", assetsDir)
	}
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(assetsDir))))

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
