package functions

import "net/http"

// Theme is a type representing a theme
type Theme string

// Constants representing the different theme
const (
	Light Theme = "light"
	Dark  Theme = "dark"
)

// themeList is a list of all the theme
var themeList = []Theme{Light, Dark}

func StrToTheme(s string) Theme {
	for _, theme := range themeList {
		if s == string(theme) {
			return theme
		}
	}
	return Light
}

// GetAndResetUserTheme return the theme of the user if it exists
// Otherwise it will set it at its default value (Light)
func GetAndResetUserTheme(w http.ResponseWriter, r *http.Request) Theme {
	cookie := GetCookie(w, r, "theme")
	if cookie == nil {
		SetCookie(w, "theme", string(Light))
		return Light
	}
	return StrToTheme(cookie.Value)
}
