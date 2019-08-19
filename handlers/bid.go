package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/iuryfukuda/auction/app"
)

func Bid(w http.ResponseWriter, r *http.Request) {
	d, ok := dataFromReq(w, r)
	if !ok {
		return
	}

	var bd app.BidData
	if err := json.Unmarshal(d, &bd); err != nil {
		badRequest(w, err.Error())
		return
	}

	w.Write(d)
}
