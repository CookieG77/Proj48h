package functions

import (
	"encoding/json"
	"fmt"
	"os"
)

// Lang is a type representing a language
type Lang string

// Constants representing the different languages
const (
	Fr Lang = "fr"
	En Lang = "en"
)

// List of all the languages
var langList = []Lang{Fr, En}

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
