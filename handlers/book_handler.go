package handlers

import (
	"gochi-book-api/models"
	"gochi-book-api/storage"
	"gochi-book-api/utils"
	"net/http"
)

var store = storage.GetBookStore()

// GetBookByID menangani permintaan GET /books/{id}.
// Fungsi ini mengambil data buku berdasarkan ID dari penyimpanan (store).
// Jika ID tidak valid atau buku tidak ditemukan, akan mengembalikan respons error.
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ParseIDParam(w, r)
	if !ok {
		return
	}

	book, found := store.GetBookByID(id)
	if !found {
		utils.RespondNotFound(w, "Book")
		return
	}

	utils.RespondJSON(w, http.StatusOK, book)
}

// GetBooks menangani permintaan GET /books.
// Fungsi ini mengembalikan seluruh daftar buku yang tersedia dalam bentuk JSON.
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := store.GetAllBooks()
	utils.RespondJSON(w, http.StatusOK, books)
}

// CreateBook menangani permintaan POST /books.
// Fungsi ini akan membaca data buku dari body request (dalam format JSON),
// lalu menyimpannya ke dalam penyimpanan.
// Jika berhasil, akan mengembalikan buku yang baru dibuat.
// Jika data JSON tidak valid, akan mengembalikan status 400.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if !utils.DecodeJSONBody(w, r, &book) {
		return
	}

	book = store.AddBook(book)
	utils.RespondJSON(w, http.StatusCreated, book)
}

// UpdateBook menangani permintaan PUT /books/{id}.
// Fungsi ini akan memperbarui data buku berdasarkan ID dengan data baru dari body request.
// Jika ID tidak valid atau buku tidak ditemukan, maka akan mengembalikan error.
// Jika berhasil, akan mengembalikan data buku yang sudah diperbarui.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ParseIDParam(w, r)
	if !ok {
		return
	}

	var updatedBook models.Book
	if !utils.DecodeJSONBody(w, r, &updatedBook) {
		return
	}

	book, ok := store.UpdateBook(id, updatedBook)
	if !ok {
		utils.RespondNotFound(w, "Book")
		return
	}

	utils.RespondJSON(w, http.StatusOK, book)
}

// DeleteBook menangani permintaan DELETE /books/{id}.
// Fungsi ini menghapus data buku berdasarkan ID dari penyimpanan.
// Jika ID tidak valid atau buku tidak ditemukan, maka akan mengembalikan error.
// Jika berhasil, akan mengembalikan pesan sukses.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ParseIDParam(w, r)
	if !ok {
		return
	}

	if ok := store.DeleteBook(id); !ok {
		utils.RespondNotFound(w, "Book")
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "Book deleted successfully"})
}
