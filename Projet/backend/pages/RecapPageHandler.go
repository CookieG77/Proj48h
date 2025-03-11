package pages

import (
	f "Proj48h/functions"
	"net/http"
)

// ReportPage is the handler for the report page.
func ReportPage(w http.ResponseWriter, r *http.Request) {
	// Checking if request is a POST request if not redirect to the home page
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		f.ErrorPrintln("Method not allowed when trying to reach the report page")
		return
	}

	// Getting the file ID from the cookie
	fileID := f.GetCookie(w, r, "FileID").Value
	if fileID == "" {
		http.Error(w, "No file ID provided", http.StatusBadRequest)
		f.ErrorPrintln("No file ID provided when trying to download a file")
		return
	}

	// Getting the website URL from the form
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error while parsing form", http.StatusInternalServerError)
		f.ErrorPrintln("Error while parsing form when trying to download a file")
		return
	}
	websiteUrl := r.Form.Get("realWebsiteUrl")
	if websiteUrl == "" {
		http.Error(w, "No website URL provided", http.StatusBadRequest)
		f.ErrorPrintln("No website URL provided when trying to download a file")
		return
	}

	f.InfoPrintf("Trying to generate a report for the website %s with id -> %s\n", websiteUrl, fileID)

	// TODO: Implement the report generation here

	PageInfo := f.NewContentInterface("report", w, r)

	tmpl := f.MakeTemplate(w, "templates/report.html")

	f.TemplateToPDF(tmpl, "statics/css/style.css", PageInfo, fileID)

	f.AddAdditionalScriptsToContentInterface(&PageInfo, "/js/reportScript.js")
	f.ExecuteTemplate(w, tmpl, PageInfo)
}
