package internal

import (
	"fmt"
	"net/http"
)

func (s *Server) handleListParroquia() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parroquia := r.URL.Query().Get("parroquia")
		fuentes, err := s.fr.ListByParroquia(parroquia)
		if err != nil {
			s.handleError(w, 500, "Error recuperando datos")
			return
		}

		if len(fuentes) < 1 {
			s.handleError(w, 404, "Ninguna fuente encontrada")
			return
		}

		for i, fuente := range fuentes {
			fmt.Fprintf(w, "%d => %s\n", i, fuente)
		}
	}
}
