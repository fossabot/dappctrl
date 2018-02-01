package serv

import (
	"fmt"
	"net/http"
)

func (s *Server) handleAuth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Don't even try!")
}
