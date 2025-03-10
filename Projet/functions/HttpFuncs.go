package functions

import (
	"fmt"
	"html/template"
	"net/http"
)

// MakeTemplate crée un templates à partir d'un ou plusieurs fichiers de templates donné en paramètre sous la forme de leur chemin en string
func MakeTemplate(w http.ResponseWriter, templatesDir ...string) *template.Template {
	templatesDir = append(templatesDir, "templates/base.html")
	tmpl, err := template.New("base.html").ParseFiles(templatesDir...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	return tmpl
}

// ExecuteTemplate exécute un templates donner en paramètre
func ExecuteTemplate(w http.ResponseWriter, tmpl *template.Template, content interface{}) {
	if tmpl == nil {
		http.Error(w, "Template is nil", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, content); err != nil {
		fmt.Printf("Error: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// MakeTemplateAndExecute est à utiliser pour les pages qui ne nécessitent pas de traitement particulier
func MakeTemplateAndExecute(w http.ResponseWriter, r *http.Request, content interface{}, templatesDir ...string) {
	tmpl := MakeTemplate(w, templatesDir...)
	ExecuteTemplate(w, tmpl, content)
}

// NewContentInterface renvoie une map[string]interface{} avec un titre donné en paramètre
func NewContentInterface(pageTitleKey string, w http.ResponseWriter, r *http.Request) map[string]interface{} {
	ContentInterface := make(map[string]interface{})
	// On récupère la langue de l'utilisateur
	currentLang := GetAndSetUserLang(w, r)
	fmt.Println(currentLang)
	langText, err := GetLangContent(currentLang)
	if err != nil {
		panic(err)
	}
	ContentInterface["Lang"] = langText
	ContentInterface["Title"] = langText["pageNames"].(map[string]interface{})[pageTitleKey]
	ContentInterface["LangList"] = LangListToStrList(langList)
	ContentInterface["CurrentLang"] = string(currentLang)

	return ContentInterface
}
