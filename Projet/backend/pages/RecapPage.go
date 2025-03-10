package pages

import (
	f "Proj48h/functions"
	"net/http"
)

func ReportPage(w http.ResponseWriter, r *http.Request) {
	PageInfo := f.NewContentInterface("report", w, r)

	f.MakeTemplateAndExecute(w, r, PageInfo, "templates/report.html")
}
