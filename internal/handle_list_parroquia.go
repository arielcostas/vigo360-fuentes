package internal

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) handleListParroquia() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parroquia := mux.Vars(r)["id"]
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
