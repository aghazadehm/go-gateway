package middleware

import "net/http"

func Policy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		token := r.Header.Get("Authorization")
		
		if token == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		if !hasAccess(token, r.URL.Path) {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func hasAccess(token, path string) bool {
	// Extract bearer token
	if len(token) < 7 || token[:7] != "Bearer " {
		return false
	}
	
	tokenValue := token[7:]
	
	// Simple token validation - in production, validate JWT signature
	if tokenValue == "" {
		return false
	}
	
	// Define path-based access rules
	// Example: admin paths require "admin" role
	adminPaths := map[string]bool{
		"/admin":    true,
		"/settings": true,
		"/users":    true,
	}
	
	// Check if path requires admin access
	if adminPaths[path] {
		// In production, decode JWT and check claims for role
		// For now, check if token contains "admin"
		return tokenValue == "admin-token"
	}
	
	// Public paths are accessible with valid token
	return true
}