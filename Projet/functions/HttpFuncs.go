package functions

import (
	"github.com/jung-kurt/gofpdf"
	"html/template"
	"net/http"
)

// MakeTemplate create a template from one or more template files given as parameter in the form of their path in string.
func MakeTemplate(w http.ResponseWriter, templatesDir ...string) *template.Template {
	templatesDir = append(templatesDir, "templates/base.html")
	tmpl, err := template.New("base.html").ParseFiles(templatesDir...)
	if err != nil {
		ErrorPrintf("An error occurred while trying to parse the template -> %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	return tmpl
}

// ExecuteTemplate execute a template given as parameter.
func ExecuteTemplate(w http.ResponseWriter, tmpl *template.Template, content interface{}) {
	if tmpl == nil {
		ErrorPrintln("An error occurred while trying to execute a template -> Template is nil")
		http.Error(w, "Template is nil", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, content); err != nil {
		ErrorPrintf("An error occurred while trying to execute a template -> %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// MakeTemplateAndExecute is to be used for pages that do not require any particular processing
func MakeTemplateAndExecute(w http.ResponseWriter, r *http.Request, content interface{}, templatesDir ...string) {
	tmpl := MakeTemplate(w, templatesDir...)
	ExecuteTemplate(w, tmpl, content)
}

// NewContentInterface return a map[string]interface{} with a title given as parameter
// It also set the language of the user and the list of available languages, as well as the page theme.
func NewContentInterface(pageTitleKey string, w http.ResponseWriter, r *http.Request) map[string]interface{} {
	ContentInterface := make(map[string]interface{})
	// Getting the user language
	currentLang := GetAndResetUserLang(w, r)
	langText, err := GetLangContent(currentLang)
	if err != nil {
		ErrorPrintf("An error occurred while trying to get the language content -> %v/n", err)
	} else {
		ContentInterface["Lang"] = langText
		ContentInterface["Title"] = langText["pageNames"].(map[string]interface{})[pageTitleKey]
	}

	// Setting the language
	ContentInterface["LangList"] = LangListToStrList(langList)
	ContentInterface["CurrentLang"] = string(currentLang)

	// // Setting the theme
	// currentTheme := GetAndResetUserTheme(w, r)
	// ContentInterface["CurrentTheme"] = string(currentTheme)

	return ContentInterface
}

func ButtonPressed(w http.ResponseWriter, r *http.Request) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "B", 12)
	pdf.AddPage()
	pdf.Cell(40, 10, "Hello, World!")

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=example.pdf")
	pdf.Output(w)
}
