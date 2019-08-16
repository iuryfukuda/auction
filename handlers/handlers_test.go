package handlers_test

import (
	"fmt"
	"testing"
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

