package internal

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) handleShowFuente() http.HandlerFunc {
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

		fmt.Fprintf(w, "fuente => %s\n", fuente)
	}
}
