package internal

import (
	"database/sql"
	"fmt"
	"net/http"

	"vigo360.es/fuentes/internal/fuente"
)

func (s *Server) handleListParroquia(fuentes fuente.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parroquia := r.URL.Query().Get("parroquia")
		fuentes, err := fuentes.ListByParroquia(parroquia)
		if err != nil {
			s.handleError(w, 500, "Error recuperando datos")
			return
		}

		if len(fuentes) < 1 {
			s.handleError(w, 404, err.Error())

			http.Error(w, sql.ErrNoRows.Error(), http.StatusNotFound)
			return
		}

		for i, fuente := range fuentes {
			fmt.Fprintf(w, "%d => %s\n", i, fuente)
		}
	}
}
