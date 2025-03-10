package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Lang is a type representing a language
type Lang string

// Constants representing the different languages
const (
	En Lang = "en"
	Fr Lang = "fr"
)

// langList is a list of all the languages
var langList = []Lang{En, Fr}

// GetLangContent returns the map[string]string containing each field text adapted to the given language
// Also returns an error if the file can't be read or if the json can't be unmarshalled
func GetLangContent(language Lang) (map[string]interface{}, error) {
	filePath := fmt.Sprintf("statics/lang/%s.json", language)

	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var content map[string]interface{}
	err = json.Unmarshal(file, &content)
	if err != nil {
		return nil, err
	}

	return content, nil
}

// StrToLang converts a string to a Lang
// Returns En if the string doesn't match any Lang
func StrToLang(str string) Lang {
	for _, l := range langList {
		if string(l) == str {
			return l
		}
	}
	return En
}

// LangListToStrList return the given list of Lang and return it as list of string
func LangListToStrList(langList []Lang) []string {
	strList := make([]string, len(langList))
	for i, l := range langList {
		strList[i] = string(l)
	}
	return strList
}

// GetAndResetUserLang return the language of the user if it exists
// // Otherwise it will set it at its default value (En)
func GetAndResetUserLang(w http.ResponseWriter, r *http.Request) Lang {
	cookie := GetCookie(w, r, "lang")
	if cookie == nil {
		SetCookie(w, "lang", string(En))
		InfoPrintf("Resetting user language -> %v\n", En)
		return En
	}
	InfoPrintf("User language -> %v\n", StrToLang(cookie.Value))
	return StrToLang(cookie.Value)
}
