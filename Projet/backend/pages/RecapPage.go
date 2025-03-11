package pages

import (
	f "Proj48h/functions"
	"net/http"
)

// ReportPage is the handler for the report page.
func ReportPage(w http.ResponseWriter, r *http.Request) {
	PageInfo := f.NewContentInterface("report", w, r)

	f.MakeTemplateAndExecute(w, r, PageInfo, "templates/report.html")
}
