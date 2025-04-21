package routes

import (
	"net/http"

	"github.com/RiddlerXenon/Integralize/internal/handler"
)

type Route struct {
	Path    string
	Handler http.HandlerFunc
}

func RegisterRoutes(mux *http.ServeMux) {
	routes := []Route{
		{Path: "/api/integral", Handler: handler.IntegralHandler},
		{Path: "/api/differential", Handler: handler.DiffEquationsHandler},
		{Path: "/api/predvictim", Handler: handler.PredVictimHandler},
	}

	for _, route := range routes {
		mux.HandleFunc(route.Path, route.Handler)
	}
}
