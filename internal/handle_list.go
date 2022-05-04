package internal

import (
	"fmt"
	"net/http"
	"os"
)

func (s *Server) handleList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fuentes, err := s.fr.List()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error listando fuentes: %s\n", err.Error())
			s.handleError(w, 500, "Error recuperando datos")
			return
		}

		for i, fuente := range fuentes {
			fmt.Fprintf(w, "%d => %s\n", i, fuente)
		}
	}
}
