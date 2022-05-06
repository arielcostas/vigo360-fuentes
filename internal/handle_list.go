package internal

import (
	"fmt"
	"net/http"
	"os"

	"vigo360.es/fuentes/internal/fuente"
	"vigo360.es/fuentes/internal/templates"
)

func (s *Server) handleList() http.HandlerFunc {
	type ResponseParams struct {
		Fuentes []fuente.Fuente
	}

	return func(w http.ResponseWriter, r *http.Request) {
		fuentes, err := s.fr.List()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error listando fuentes: %s\n", err.Error())
			s.handleError(w, 500, "Error recuperando datos")
			return
		}

		err = templates.Render(w, "index.html", ResponseParams{
			Fuentes: fuentes,
		})
		if err != nil {
			fmt.Fprintf(w, "Hubo un error mostrando la página")
		}
	}
}
