package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSON レスポンスボディの書込
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
