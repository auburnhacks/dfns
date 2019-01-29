package getqr

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetQR(t *testing.T) {
	payload := strings.NewReader(`{"user_id": "fsdfadfadf", "width": 100, "height": 100}`)
	req := httptest.NewRequest("POST", "/", payload)
	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	GetQR(rr, req)

	out, err := ioutil.ReadAll(rr.Result().Body)
	if err != nil {
		t.Fatal(err)
	}
	if len(out) == 0 {
		t.Fatalf("no response recorded")
	}
}
