package util

import (
	"encoding/json"
	"fmt"
	"log"
)

// NewHTTPResponseErr is a utility function that returns an error in JSON
// form
func NewHTTPResponseErr(errIn error) []byte {
	resp := map[string]interface{}{
		"error": fmt.Sprintf("%v", errIn),
	}
	b, err := json.Marshal(resp)
	if err != nil {
		log.Printf("error: %v", err)
	}
	return b
}
