// Package json handles the encode & decode for json
package json

import (
	"encoding/json"
	"net/http"
)

func Write(w http.ResponseWriter, status int, data any) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
