package internal

import (
	"fmt"
	"net/http"
)

func (s *Server) handleError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	fmt.Fprintf(w, `<!DOCTYPE html><html><body><h1>%s</h1></body></html>`, message)
}
