package pages

import (
	f "Proj48h/functions"
	"net/http"
)

// HomePage is the handler for the home page.
func HomePage(w http.ResponseWriter, r *http.Request) {
	PageInfo := f.NewContentInterface("home", w, r)
	tmpl := f.MakeTemplate(w, "templates/home.html")

	f.TemplateToPDF(tmpl, "statics/css/style.css", PageInfo)

	f.ExecuteTemplate(w, tmpl, PageInfo)
}
