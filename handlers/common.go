package handlers

import (
	"fmt"
	"strconv"
	"net/http"
	"io/ioutil"
)

func dataFromReq(w http.ResponseWriter, r *http.Request) (d []byte, ok bool) {
	if r.Body == nil {
		badRequest(w, "Body is Nil")
		return
	}

	defer r.Body.Close()

	var err error
	d, err = ioutil.ReadAll(r.Body)
	if err != nil {
		badRequest(w, err.Error())
		return
	}
	return d, true
}

func badRequest(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(w, strconv.Quote(msg))
}

func setupHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
