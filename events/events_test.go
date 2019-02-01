package events

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetEvents(t *testing.T) {
	payload := strings.NewReader("")
	req := httptest.NewRequest("GET", "/", payload)
	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	h := &Handler{}
	h.ServeHTTP(rr, req)

	_, err := ioutil.ReadAll(rr.Result().Body)
	if err != nil {
		t.Fatal(err)
	}
}
