package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/iuryfukuda/auction/models"
)

type Checker interface {
	Check() models.Stats
}

type stats struct {
	Checker
}

func NewStats(c Checker) *stats {
	return &stats{c}
}

func (s *stats) Serve(w http.ResponseWriter, r *http.Request) {
	setupHeaders(w)
	stats := s.Check()
	b, err := json.Marshal(stats)
	if err != nil {
		badRequest(w, err.Error())
		return
	}
	w.Write(b)
}
