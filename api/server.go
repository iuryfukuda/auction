package api

import (
	"net/http"

	"github.com/iuryfukuda/auction/handlers"
	"github.com/iuryfukuda/auction/db"
)

type server struct {
	mux	*http.ServeMux
	db	*db.Mem
}

func NewServer() *server {
	return &server{
		mux: http.NewServeMux(),
		db: db.NewMem(),
	}
}

func (s *server) SetupHandlers() {
	s.mux.HandleFunc("/bid", handlers.NewBid(s.db).Serve)
	s.mux.HandleFunc("/stats", handlers.NewStats(s.db).Serve)
}

func (s *server) Start(port string) error {
	return http.ListenAndServe(port, s.mux)
}

func (s *server) Run(port string) error {
	s.SetupHandlers()
	return s.Start(port)
}
