package routes

import (
	"net/http"

	"github.com/RiddlerXenon/Integralize/internal/handler"
)

func registerAPIRoutes(mux *http.ServeMux) {
	apiRoutes := []Route{
		{Path: "/api/integral", Handler: http.HandlerFunc(handler.IntegralHandler)},
		{Path: "/api/differential", Handler: http.HandlerFunc(handler.DiffEquationsHandler)},
		{Path: "/api/predvictim", Handler: http.HandlerFunc(handler.PredVictimHandler)},
	}

	for _, r := range apiRoutes {
		mux.Handle(r.Path, r.Handler)
	}
}
