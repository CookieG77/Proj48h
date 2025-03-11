package pages

import (
	f "Proj48h/functions"
	"net/http"
)

// HomePage is the handler for the home page.
func HomePage(w http.ResponseWriter, r *http.Request) {
	PageInfo := f.NewContentInterface("home", w, r)
	fileID := f.GenerateHexFilename()
	f.SetCookie(w, "FileID", fileID)
	f.AddAdditionalScriptsToContentInterface(&PageInfo, "js/homeScript.js")
	f.MakeTemplateAndExecute(w, r, PageInfo, "templates/home.html")
}
