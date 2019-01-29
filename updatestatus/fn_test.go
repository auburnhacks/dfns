package updatestatus

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestUpdateStatusUserLink(t *testing.T) {
	d := struct {
		EventID string `json:"event_id,omitempty"`
		UserID  string `json:"user_id,omitempty"`
		LinkID  int    `json:"link_id,omitempty"`
	}{EventID: "user-link", UserID: "23rfe", LinkID: 2342}
	bb, err := json.Marshal(&d)
	t.Logf("sending: %s to function", bb)
	if err != nil {
		t.Fatalf("error while marshaling: %v", err)
	}
	r := bytes.NewReader(bb)
	req := httptest.NewRequest("GET", "/", r)

	rr := httptest.NewRecorder()
	UpdateStatus(rr, req)

	out, err := ioutil.ReadAll(rr.Result().Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "ok" {
		t.Fatalf("error: expected ok got %s", out)
	}
}

func TestUpdateStatusAddAttendee(t *testing.T) {
	d := struct {
		EventID string `json:"event_id,omitempty"`
		UserID  string `json:"user_id,omitempty"`
		LinkID  int    `json:"link_id,omitempty"`
	}{EventID: "sat-lunch", LinkID: 242}

	bb, err := json.Marshal(&d)
	if err != nil {
		t.Fatal(err)
	}
	r := bytes.NewReader(bb)
	req := httptest.NewRequest("GET", "/", r)

	rr := httptest.NewRecorder()
	UpdateStatus(rr, req)

	out, err := ioutil.ReadAll(rr.Result().Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("out: %s", out)
}
