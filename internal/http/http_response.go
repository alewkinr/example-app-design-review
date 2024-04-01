package http

import (
	"encoding/json"
	"net/http"
)

func EncodeJSONResponse(i interface{}, status int, w http.ResponseWriter) error {
	wHeader := w.Header()
	wHeader.Set("Content-Type", "application/json; charset=UTF-8")

	if status != 0 {
		w.WriteHeader(status)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	if i != nil {
		return json.NewEncoder(w).Encode(i)
	}

	return nil
}
