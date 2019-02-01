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
		{EventID: "check-in", DisplayName: "Hacker Check-in"},
		{EventID: "user-link", DisplayName: "HackerLink"},
		{EventID: "sat-lunch", DisplayName: "Saturday Lunch"},
		{EventID: "sat-dinner", DisplayName: "Saturday Dinner"},
		{EventID: "midnight-snack", DisplayName: "Midnight Snack"},
		{EventID: "sun-breakfast", DisplayName: "Sunday BreakFast"},
		{EventID: "sun-lunch", DisplayName: "Sunday Lunch"},
		{EventID: "techtalk-1", DisplayName: "1:30pm Tech Talk"},
		{EventID: "techtalk-2", DisplayName: "5:00pm Tech Talk"},
		{EventID: "techtalk-3", DisplayName: "7:30pm Tech Talk"},
		{EventID: "techtalk-4", DisplayName: "9:00pm Tech Talk"},
		{EventID: "techtalk-5", DisplayName: "11:00pm Tech Talk"},
		{EventID: "techtalk-6", DisplayName: "10:00am Tech Talk"},
		{EventID: "mc1", DisplayName: "Mini Challenge #1"},
		{EventID: "mc2", DisplayName: "Mini Challenge #2"},
		{EventID: "mc3", DisplayName: "Mini Challenge #3"},
		{EventID: "mc4", DisplayName: "Mini Challenge: Tech Trivia Contest"},
	}
)

// Handler is a http handler
type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	bb, err := json.Marshal(allEvents)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error: %v", err)))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bb)
}
