package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/iuryfukuda/auction/models"
)

type Savior interface {
	Save(bd models.BidData)
}

type bid struct {
	Savior
}

func NewBid(s Savior) *bid {
	return &bid{s}
}

func (b *bid) Serve(w http.ResponseWriter, r *http.Request) {
	setupHeaders(w)
	d, ok := dataFromReq(w, r)
	if !ok {
		badRequest(w, "Missing body in request")
		return
	}

	var bd models.BidData
	if err := json.Unmarshal(d, &bd); err != nil {
		badRequest(w, err.Error())
		return
	}

	b.Save(bd)
}
