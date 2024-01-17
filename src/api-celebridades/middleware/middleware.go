package middleware

import "net/http"

func SetContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(write http.ResponseWriter, request *http.Request) {
		write.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(write, request)
	})
}
