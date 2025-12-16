package middleware

import "net/http"

func Policy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		token := r.Header.Get("Authorization")
		1
		if token == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
		}

		if !hasAccess(token, r.URL.Path) {
			http.Error(w, "forbidden", http.StatusForbidden)
		}

		next.ServerHTTP(w, r)
	})
}

func hasAccess(token, path string) bool {
	return true
}