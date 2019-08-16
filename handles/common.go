package handles

import (
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
	http.Error(w, strconv.Quote(msg), http.StatusBadRequest)
}
