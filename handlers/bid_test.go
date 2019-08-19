package handlers_test

import (
	"fmt"
	"testing"
	"strings"
	"net/http"
	"net/http/httptest"

	"github.com/iuryfukuda/auction/handlers"
	"github.com/iuryfukuda/auction/models"
)

type fakeSavior struct{}
func (fs *fakeSavior) Save(bd models.BidData) error { return nil }

type bidTest struct {
	req	*http.Request
	code	int
	ct	string
}

var bidTests = []bidTest{
	bidTest{
		req: httptest.NewRequest(http.MethodPost, "/bid", nil),
		code: http.StatusBadRequest,
		ct: "application/json",
	},
	bidTest{
		req: httptest.NewRequest(
			http.MethodPost, "/bid", strings.NewReader(`{"item_id": "123", "price": 1.99, "client_id": "123"}`),
		),
		code: http.StatusOK,
		ct: "application/json",
	},
}

func runBidTest(t bidTest) error {
	var fs *fakeSavior
	bid := handlers.NewBid(fs)
	w := httptest.NewRecorder()
	bid.Serve(w, t.req)
	if w.Code != t.code {
		return fmt.Errorf("got [%#v], want [%#v]", w.Code, t.code)
	}
	if ct := w.HeaderMap["Content-Type"][0]; ct != t.ct {
		return fmt.Errorf("got [%#v], want [%#v]", ct, t.ct)
	}
	return nil
}

func TestBid(t *testing.T) {
	for i, test := range bidTests {
		if err := runBidTest(test); err != nil {
			t.Fatalf("[%d]: %s", i, err)
		}
	}
}

func BenchmarkBid(b *testing.B) {
	for i, test := range bidTests {
		if err := runBidTest(test); err != nil {
			b.Fatalf("[%d]: %s", i, err)
		}
	}
}
