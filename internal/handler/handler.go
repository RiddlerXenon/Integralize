package handler

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

// Инициализация шаблонов (вызывается из main.go)
func InitTemplates() {
	var err error
	templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Ошибка при парсинге шаблонов: %v", err)
	}
}

// HTTP-обработчик
func RenderPage(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Path[1:]
	if page == "" {
		page = "index"
	}
	filename := page + ".html"

	tmpl := templates.Lookup(filename)
	if tmpl == nil {
		http.NotFound(w, r)
		return
	}

	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Ошибка рендера шаблона", http.StatusInternalServerError)
	}
}
