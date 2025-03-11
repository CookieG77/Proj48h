package functions

import (
	"net/http"
)

// GetCookie returns the cookie with the given name.
func GetCookie(w http.ResponseWriter, r *http.Request, name string) *http.Cookie {
	cookie, err := r.Cookie(name)
	if err != nil {
		return nil
	}
	return cookie
}

// SetCookie set the cookie with the given name to the given value.
// This cookie is not meant to be used for marketing or data analysing of the user.
// This implementation only serve to store a value in the user browser.
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
