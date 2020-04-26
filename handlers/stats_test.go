package handlers_test

import (
	"fmt"
	"testing"
	"reflect"
	"net/http"
	"encoding/json"
	"net/http/httptest"

	"github.com/zbioe/auction/handlers"
	"github.com/zbioe/auction/models"
)

type fakeChecker struct{}
func (fs *fakeChecker) Check() models.Stats { return models.Stats{} }

type statsTest struct {
	req	*http.Request
	code	int
	ct	string
	want	models.Stats
}

var statsTests = []statsTest{
	statsTest{
		req: httptest.NewRequest(http.MethodGet, "/stats", nil),
		code: http.StatusOK,
		ct: "application/json",
		want: models.Stats{0,0,nil},
	},
}

func runStatsTest(t statsTest) error {
	var fc *fakeChecker
	stats := handlers.NewStats(fc)
	w := httptest.NewRecorder()
	stats.Serve(w, t.req)
	if w.Code != t.code {
		return fmt.Errorf("got [%#v], want [%#v]", w.Code, t.code)
	}
	if ct := w.HeaderMap["Content-Type"][0]; ct != t.ct {
		return fmt.Errorf("got [%#v], want [%#v]", ct, t.ct)
	}
	var got models.Stats
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		return fmt.Errorf("unmarshal body: %s", err)
	}
	if !reflect.DeepEqual(got, t.want) {
		return fmt.Errorf("got [%#v], want [%#v]", got, t.want)
	}
	return nil
}

func TestStats(t *testing.T) {
	for i, test := range statsTests {
		if err := runStatsTest(test); err != nil {
			t.Fatalf("[%d]: %s", i, err)
		}
	}
}

func BenchmarkStats(b *testing.B) {
	for i, test := range statsTests {
		if err := runStatsTest(test); err != nil {
			b.Fatalf("[%d]: %s", i, err)
		}
	}
}
