package internal

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"vigo360.es/fuentes/internal/fuente"
	"vigo360.es/fuentes/internal/templates"
)

func (s *Server) handleShowFuente() http.HandlerFunc {
	type ResponseParams struct {
		Fuente fuente.Fuente
		Json   template.JS
	}

	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		fuente, err := s.fr.GetById(id)
		if err != nil {
			fmt.Printf("erroor obteniendo fuente %s: %s", id, err.Error())
			if errors.Is(err, sql.ErrNoRows) {
				http.Redirect(w, r, "/", http.StatusFound)
				return
			} else {
				s.handleError(w, 500, "Error recuperando datos")
			}
			return
		}

		// Generar JSON para incluir en JavaScript (para mapa)
		jsonFuentes, err := json.Marshal(fuente)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error generando JSON de fuentes: %s", err.Error())
		}
		err = templates.Render(w, "fuente.html", ResponseParams{
			Fuente: fuente,
			Json:   template.JS(jsonFuentes),
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "error mostrando página: %s", err.Error())
			fmt.Fprintf(w, "Hubo un error mostrando la página")
		}
	}
}
