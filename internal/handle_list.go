package internal

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"vigo360.es/fuentes/internal/fuente"
	"vigo360.es/fuentes/internal/templates"
)

func (s *Server) handleList() http.HandlerFunc {
	type ResponseParams struct {
		Fuentes []fuente.Fuente
		Json    template.JS
	}

	return func(w http.ResponseWriter, r *http.Request) {
		fuentes, err := s.fr.List()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error listando fuentes: %s\n", err.Error())
			s.handleError(w, 500, "Error recuperando datos")
			return
		}

		jsonFuentes, err := json.Marshal(fuentes)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error generando JSON de fuentes: %s", err.Error())
		}
		err = templates.Render(w, "index.html", ResponseParams{
			Fuentes: fuentes,
			Json:    template.JS(jsonFuentes),
		})
		if err != nil {
			fmt.Fprintf(w, "Hubo un error mostrando la página")
		}
	}
}
