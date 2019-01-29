package events

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type event struct {
	EventID     string `json:"event_id"`
	DisplayName string `json:"display_name"`
}

var (
	allEvents = []event{
		// eventId user-link is a special event where we link a user to
		// a random number
		{EventID: "user-link", DisplayName: "HackerLink"},
		{EventID: "check-in", DisplayName: "Hacker Check-in"},
		{EventID: "sat-lunch", DisplayName: "Saturday Lunch"},
	}
)

// Handler is a http handler
type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bb, err := json.Marshal(allEvents)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error: %v", err)))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bb)
}
