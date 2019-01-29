package getqr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	apiURL  = "https://api.qrserver.com/v1/create-qr-code/"
	fgColor = "dd550c"
	bgColor = "03244d"
)

// GetQR is a function that accepts a UserID, Width, Height
// and returns a qr code in UTF-8 encoding. This handler uses the
// goqr.me api to get create qr codes.
func GetQR(w http.ResponseWriter, r *http.Request) {
	var d struct {
		UserID string `json:"user_id"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error: could not decode json body"))
	}
	defer r.Body.Close()

	qrURL := constructURL(d.UserID, d.Width, d.Height)
	log.Printf("pinging url: %s", qrURL)
	resp, err := http.Get(qrURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bb)
}

func constructURL(data string, width, height int) string {
	return fmt.Sprintf("%s?size=%dx%d&data=%s&bgcolor=%s&color=%s",
		apiURL, width, height, data, bgColor, fgColor)
}
