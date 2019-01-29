package canattend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	store "github.com/auburnhacks/dfns/hackstorage"
)

// Handler is a http handler
type Handler struct{}

// CanAttend is a cloud function that allows to check whether a hacker
// can attend an event
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dbURL, ok := os.LookupEnv("MONGO_URL")
	if !ok {
		log.Fatalf("error: MONGO_URL env variable not found")
	}
	store, err := store.New(dbURL, "dayof")
	if err != nil {
		log.Fatal(err)
	}
	// checking to see if the database if available
	err = store.Ping()
	if err != nil {
		log.Fatal(err)
	}

	var d struct {
		EventID string `json:"event_id,omitempty"`
		UserID  int    `json:"user_id,omitempty"`
	}

	err = json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w,
			fmt.Sprintf("cloud/can_attend: error while decoding json: %v", err),
			http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	// check and see if the user can attend the event
	ok, err = store.CanAttend(d.EventID, d.UserID)
	if err != nil {
		http.Error(w,
			fmt.Sprintf("cloud/can_attend: error user cannot attend: %v", err),
			http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "false", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
