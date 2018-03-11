package payment

import (
	"net/http"
)

func (s *Server) validateMethod(f http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != string(method) {
			w.WriteHeader(http.StatusNotAcceptable)
		}
		f(w, r)
	}
}
