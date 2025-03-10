package functions

import (
	"net/http"
)

// GetCookie renvoie le cookie demandé s'il existe, sinon renvoie nil
func GetCookie(w http.ResponseWriter, r *http.Request, name string) *http.Cookie {
	cookie, err := r.Cookie(name)
	if err != nil {
		return nil
	}
	return cookie
}

// SetCookie définit le cookie avec le nom 'name' donné avec la valeur 'value' donnée.
func SetCookie(w http.ResponseWriter, name string, value string) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
}

// GetAndSetUserLang renvoie la langue de l'utilisateur si elle existe.
// Sinon définit la langue de l'utilisateur à la langue par défaut (En) et la renvoie.
func GetAndSetUserLang(w http.ResponseWriter, r *http.Request) Lang {
	cookie := GetCookie(w, r, "lang")
	if cookie == nil {
		SetCookie(w, "lang", "En")
		return En
	}
	return StrToLang(cookie.Value)
}
