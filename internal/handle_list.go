package internal

import (
	"fmt"
	"net/http"

	"vigo360.es/fuentes/internal/fuente"
)

func (s *Server) handleList(fuentes fuente.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fuentes, err := fuentes.List()
		if err != nil {
			s.handleError(w, 500, "Error recuperando datos")
			return
		}

		for i, fuente := range fuentes {
			fmt.Fprintf(w, "%d => %s\n", i, fuente)
		}
	}
}
