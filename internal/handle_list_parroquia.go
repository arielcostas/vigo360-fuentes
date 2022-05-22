package internal

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"vigo360.es/fuentes/internal/fuente"
	"vigo360.es/fuentes/internal/templates"
)

func (s *Server) handleListParroquia() http.HandlerFunc {
	type ResponseParams struct {
		Nombre  string
		Fuentes []fuente.Fuente
		Json    template.JS
	}

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

		jsonFuentes, err := json.Marshal(fuentes)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error generando JSON de fuentes: %s", err.Error())
		}
		err = templates.Render(w, "parroquia.html", ResponseParams{
			Nombre:  parroquia,
			Json:    template.JS(jsonFuentes),
			Fuentes: fuentes,
		})
		if err != nil {
			fmt.Fprintf(w, "Hubo un error mostrando la página")
		}
	}
}
