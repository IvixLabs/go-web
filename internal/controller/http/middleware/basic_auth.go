package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
)

func GetBasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user, pass, ok := r.BasicAuth()
		if ok {
			username := sha256.Sum256([]byte("admin"))
			password := sha256.Sum256([]byte("admin"))
			userHash := sha256.Sum256([]byte(user))
			passHash := sha256.Sum256([]byte(pass))
			validUser := subtle.ConstantTimeCompare(userHash[:], username[:]) == 1
			validPass := subtle.ConstantTimeCompare(passHash[:], password[:]) == 1
			if validPass && validUser {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "No/Invalid Credentials", http.StatusUnauthorized)
	})
}
