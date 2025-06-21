package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// ParseIDParam digunakan untuk mengambil parameter "id" dari URL dan mengubahnya ke tipe integer.
// Jika parameter tidak valid (bukan angka), akan mengembalikan respons error 400 dan false.
func ParseIDParam(w http.ResponseWriter, r *http.Request) (int, bool) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid book ID"})
		return 0, false
	}
	return id, true
}

// DecodeJSONBody melakukan parsing (decoding) isi body permintaan HTTP dalam format JSON
// dan menyimpannya ke dalam struct tujuan (dst).
// Jika gagal decoding (misal format JSON salah), akan mengembalikan error 400 dan false.
func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return false
	}
	return true
}

// RespondNotFound mengirimkan response 404 dengan pesan "not found" untuk entitas yang disebutkan.
// Fungsi ini digunakan untuk menghindari pengulangan kode dalam menangani entitas yang tidak ditemukan.
func RespondNotFound(w http.ResponseWriter, entity string) {
	RespondJSON(w, http.StatusNotFound, map[string]string{"error": entity + " not found"})
}
