package common

import (
	"encoding/json"
	"net/http"
)

// WebJSONResponse parse the response payload and encode it to JSON to return to HTTP client
func WebJSONResponse(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(status)
	w.Write([]byte(response))
}
