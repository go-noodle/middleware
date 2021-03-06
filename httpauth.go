package middleware

import (
	"net/http"
	"net/url"

	"github.com/go-noodle/noodle"
)

// HTTPAuth is a middleware factory function that accepts the authentication realm
// and function for username and password verification. Resulting middleware injects
// username into request context if authentication successful.
func HTTPAuth(realm string, authFunc func(username, password string) bool) noodle.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			username, password, ok := r.BasicAuth()
			if !ok || !authFunc(username, password) {
				w.Header().Set("WWW-Authenticate", "Basic realm="+url.QueryEscape(realm))
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// Inject user name into request context
			next(w, noodle.WithValue(r, userKey, username))
		}
	}
}

// GetUser extract authentication information from context
func GetUser(r *http.Request) string {
	res, _ := noodle.Value(r, userKey).(string)
	return res
}
