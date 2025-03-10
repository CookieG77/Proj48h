package pages

import (
	f "Proj48h/functions"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	PageInfo := f.NewContentInterface("home", w, r)

	f.MakeTemplateAndExecute(w, r, PageInfo, "templates/home.html")
}
