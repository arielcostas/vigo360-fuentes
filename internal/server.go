package internal

import (
	"net/http"

	"github.com/gorilla/mux"
	"vigo360.es/fuentes/internal/fuente"
)

type Server struct {
	router *mux.Router
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

	s.router = mux.NewRouter().StrictSlash(true)
	s.router.HandleFunc("/fuentes/{id}", s.handleShowFuente())
	s.router.HandleFunc("/parroquias/{id}", s.handleListParroquia())
	s.router.HandleFunc("/", s.handleList())
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.router == nil {
		s.Routes()
	}
	s.router.ServeHTTP(w, r)
}
