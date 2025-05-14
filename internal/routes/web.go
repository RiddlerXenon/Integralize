package routes

import (
	"net/http"

	"github.com/RiddlerXenon/Integralize/internal/handler"
)

func registerWebRoutes(mux *http.ServeMux) {
	// Инициализация шаблонов для HTML-страниц
	handler.InitTemplates()

	// Рендеринг страниц через шаблоны
	mux.HandleFunc("/", handler.RenderPage)

	// Раздача статики (CSS, JS, изображения и т.д.)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
