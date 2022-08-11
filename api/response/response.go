package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ToJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		ToJSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	ToJSON(w, http.StatusBadRequest, nil)
}
