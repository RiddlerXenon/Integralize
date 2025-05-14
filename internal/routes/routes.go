package routes

import (
	"net/http"
)

// Route описывает маршрут и соответствующий обработчик
type Route struct {
	Path    string
	Handler http.Handler
}

// RegisterRoutes — регистрирует все маршруты (API + HTML + static)
func RegisterRoutes(mux *http.ServeMux) {
	registerAPIRoutes(mux)
	registerWebRoutes(mux)
}
