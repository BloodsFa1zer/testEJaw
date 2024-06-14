package middleware

import (
	"net/http"
	"test/response"
)

func BasicAuth(username, password string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || user != username || pass != password {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			response.JSON(w, http.StatusUnauthorized, "Unauthorized", nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}
