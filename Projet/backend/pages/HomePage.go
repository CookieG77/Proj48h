package pages

import (
	"Proj48h/functions"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	PageInfo := functions.NewContentInterface("Acceuil", w, r)

	functions.MakeTemplateAndExecute(w, r, PageInfo, "templates/home.html")
}
