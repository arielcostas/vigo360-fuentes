/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
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
		Fuentes    []fuente.Fuente
		Parroquias map[string]int
		Json       template.JS
	}

	return func(w http.ResponseWriter, r *http.Request) {
		fuentes, err := s.fr.List()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error listando fuentes: %s\n", err.Error())
			s.handleError(w, 500, "Error recuperando datos")
			return
		}

		var parroquias = make(map[string]int)
		for _, fuente := range fuentes {
			if p, ok := parroquias[fuente.Parroquia]; ok {
				parroquias[fuente.Parroquia] = p + 1
			} else {
				parroquias[fuente.Parroquia] = 1
			}
		}

		// Generar JSON para incluir en JavaScript (para mapa)
		jsonFuentes, err := json.Marshal(fuentes)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error generando JSON de fuentes: %s", err.Error())
		}
		err = templates.Render(w, "index.html", ResponseParams{
			Fuentes:    fuentes,
			Parroquias: parroquias,
			Json:       template.JS(jsonFuentes),
		})
		if err != nil {
			fmt.Fprintf(w, "Hubo un error mostrando la página")
		}
	}
}
