package api

import (
	"net/http"

	_ "github.com/iuryfukuda/auction/handlers"
)

type server struct {
	mux	*http.ServeMux
//	db	db.Mem
}

func NewServer() *server {
	return &server{
		mux: http.NewServeMux(),
//		db: db.NewMem(),
	}
}

func (s *server) SetupHandlers() {
	// s.mux.HandleFunc("/bid", handlers.NewBid(s.db))
	// s.mux.HandleFunc("/stats", handlers.NewStats(s.db))
}

func (s *server) Start(port string) error {
	return http.ListenAndServe(port, s.mux)
}
