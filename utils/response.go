package utils

import (
	"encoding/json"
	"net/http"
)

// RespondWithError mengirimkan respons JSON dengan pesan kesalahan
func RespondWithError(w http.ResponseWriter, status int, message string) {
	RespondWithJSON(w, status, map[string]string{"error": message})
}

// RespondWithJSON mengirimkan respons JSON dengan status kode tertentu
func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
