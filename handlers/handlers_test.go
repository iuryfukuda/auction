package handlers_test

import (
	"fmt"
	"testing"
	"strings"
	"net/http"
	"net/http/httptest"

	"github.com/iuryfukuda/auction/handlers"
)

type handlersTest struct {
	req	*http.Request
	code	int
	ct	string
}

var handlersTests = []handlersTest{
	handlersTest{
		req: httptest.NewRequest(http.MethodGet, "/bid", nil),
		code: http.StatusBadRequest,
		ct: "application/json",
	},
	handlersTest{
		req: httptest.NewRequest(
			http.MethodGet, "/bid", strings.NewReader(`{"item_id": "123", "price": 1.99, "client_id": "123"}`),
		),
		code: http.StatusOK,
		ct: "application/json",
	},
}

func runHandlersTest(t handlersTest) error {
	w := httptest.NewRecorder()
	handlers.Bid(w, t.req)
	if w.Code != t.code {
		return fmt.Errorf("got [%#v], want [%#v]", w.Code, t.code)
	}
	if ct := w.HeaderMap["Content-Type"][0]; ct != t.ct {
		return fmt.Errorf("got [%#v], want [%#v]", ct, t.ct)
	}
	return nil
}

func TestHandler(t *testing.T) {
	for i, test := range handlersTests {
		if err := runHandlersTest(test); err != nil {
			t.Fatalf("[%d]: %s", i, err)
		}
	}
}

func BenchmarkHandler(b *testing.B) {
	for i, test := range handlersTests {
		if err := runHandlersTest(test); err != nil {
			b.Fatalf("[%d]: %s", i, err)
		}
	}
}
