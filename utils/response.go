package utils

import (
	"encoding/json"
	"net/http"
)

// RespondJSON mengirimkan respons HTTP dalam format JSON.
// Fungsi ini menerima parameter:
// - w: objek http.ResponseWriter untuk menulis response,
// - status: kode status HTTP (misalnya 200, 404, 500),
// - data: data yang akan dikirim sebagai JSON.
//
// Header "Content-Type" akan diset ke "application/json",
// dan data akan di-encode ke dalam format JSON secara otomatis.
func RespondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
