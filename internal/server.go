package internal

import (
	"net/http"

	"vigo360.es/fuentes/internal/database"
	"vigo360.es/fuentes/internal/fuente"
)

type Server struct {
	router *http.ServeMux
}

func (s *Server) Routes() {
	if s.router != nil {
		return
	}

	var db = database.GetDB()
	var fr = fuente.NewMysqlRepository(db)

	s.router = http.NewServeMux()
	s.router.HandleFunc("/parroquia", s.handleListParroquia(fr))
	s.router.HandleFunc("/", s.handleList(fr))
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.router == nil {
		s.Routes()
	}
	s.router.ServeHTTP(w, r)
}
