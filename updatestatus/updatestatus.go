package updatestatus

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	store "github.com/auburnhacks/dfns/hackstorage"
	"github.com/auburnhacks/dfns/util"
)

type Handler struct{}

// ServeHTTP is a function that should be implemented for every http handler
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get storage client
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
		EventID string `json:"event_id"`
		UserID  string `json:"user_id"`
		LinkID  int    `json:"link_id"`
	}
	err = json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		log.Printf("[ERROR] error while decoding: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(util.NewHTTPResponseErr(err))
		return
	}
	defer r.Body.Close()

	if d.EventID == "user-link" {
		// perform userlinking
		err := linkUser(d.UserID, d.LinkID, store)
		if err != nil {
			log.Printf("[ERROR] error while linking user: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(util.NewHTTPResponseErr(err))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
		return
	}
	err = store.AddAttendee(d.EventID, d.LinkID)
	if err != nil {
		log.Printf("[ERROR] error while adding attendee: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(util.NewHTTPResponseErr(err))
		return
	}
}

func linkUser(userID string, randNum int, store store.Storage) error {
	err := store.LinkUser(userID, randNum)
	if err != nil {
		return err
	}
	return nil
}

func addAttendee(eventID string, randNum int, store store.Storage) error {
	return nil
}
