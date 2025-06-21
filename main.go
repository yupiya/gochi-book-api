package main

import (
	"gochi-book-api/routes"
	"log"
	"net/http"
)

func main() {
	// Inisialisasi router menggunakan fungsi NewRouter dari package routes.
	r := routes.NewRouter()

	// Menampilkan pesan bahwa server telah berjalan
	log.Println("Server running at http://localhost:8080")

	// Menjalankan server HTTP di port 8080.
	// log.Fatal akan menghentikan program jika terjadi error saat menjalankan server.
	log.Fatal(http.ListenAndServe(":8080", r))
}
