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
}

var handlersTests = []handlersTest{
	handlersTest{
		req: httptest.NewRequest(http.MethodGet, "/bid", nil),
		code: http.StatusBadRequest,
	},
	handlersTest{
		req: httptest.NewRequest(
			http.MethodGet, "/bid", strings.NewReader(`{"item_id": "123", "price": 1.99, "client_id": "123"}`),
		),
		code: http.StatusOK,
	},
}

func runHandlersTest(t handlersTest) error {
    w := httptest.NewRecorder()
	handlers.Bid(w, t.req)
	if w.Code != t.code {
		return fmt.Errorf("got [%#v], want [%#v]", w.Code, t.code)
	}
	return nil
}

func TestHandler(t *testing.T) {
	for _, test := range handlersTests {
		if err := runHandlersTest(test); err != nil {
			t.Fatal(err)
		}
	}
}

