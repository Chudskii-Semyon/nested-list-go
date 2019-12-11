package middleware

import "net/http"

type Adapter func(h http.Handler) http.Handler

func CORS(w *http.ResponseWriter) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		})
	}
}
