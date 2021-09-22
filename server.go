package pdcnotification

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

func Notification(w http.ResponseWriter, r *http.Request) {

	var d struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Hello World!")
		return
	}
	if d.Message == "" {
		fmt.Fprint(w, "Hello World!")
		return
	}
	fmt.Fprint(w, html.EscapeString(d.Message))
}
