package canattend

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("MONGO_URL", "mongodb://localhost:27017/dayof")
	// set up a test event so that we can
	d := struct {
		EventID string `json:"event_id,omitempty"`
		UserID  string `json:"user_id,omitempty"`
		LinkID  int    `json:"link_id,omitempty"`
	}{EventID: "sat-lunch", UserID: "12121", LinkID: 234}
	bb, err := json.Marshal(&d)
	if err != nil {
		log.Fatalf("error while marshaling: %v", err)
	}
	r := bytes.NewReader(bb)
	req := httptest.NewRequest("GET", "/", r)
	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	UpdateStatus(rr, req)

	_, err = ioutil.ReadAll(rr.Result().Body)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}

func TestCanAttendNotOk(t *testing.T) {
	d := struct {
		EventID string `json:"event_id,omitempty"`
		UserID  int    `json:"user_id,omitempty"`
	}{EventID: "sat-lunch", UserID: 234}

	bb, err := json.Marshal(&d)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest("GET", "/", bytes.NewReader(bb))
	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	CanAttend(rr, req)

	out, err := ioutil.ReadAll(rr.Result().Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) == "ok" {
		t.Fatalf("expected anything other than ok but got %s", out)
	}
}

func TestCanAttendOk(t *testing.T) {
	d := struct {
		EventID string `json:"event_id,omitempty"`
		UserID  int    `json:"user_id,omitempty"`
	}{EventID: "sat-lunch", UserID: 23442}

	bb, err := json.Marshal(&d)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest("GET", "/", bytes.NewReader(bb))
	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	CanAttend(rr, req)

	out, err := ioutil.ReadAll(rr.Result().Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "ok" {
		t.Fatalf("expected an ok from server got: %s", out)
	}
}
