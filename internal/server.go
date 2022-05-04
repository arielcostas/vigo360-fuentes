package internal

import (
	"net/http"

	"vigo360.es/fuentes/internal/fuente"
)

type Server struct {
	router *http.ServeMux
	fr     fuente.Repository
}

func NewServer(fr fuente.Repository) *Server {
	s := &Server{
		fr: fr,
	}
	s.Routes()
	return s
}

func (s *Server) Routes() {
	if s.router != nil {
		return
	}

	s.router = http.NewServeMux()
	s.router.HandleFunc("/parroquia", s.handleListParroquia())
	s.router.HandleFunc("/", s.handleList())
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.router == nil {
		s.Routes()
	}
	s.router.ServeHTTP(w, r)
}
