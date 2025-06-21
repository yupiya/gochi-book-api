package models


// Menyimpan informasi buku dan digunakan dalam komunikasi JSON.
type Book struct {
	ID            int    `json:"id"`              // ID unik setiap buku.
	Title         string `json:"title"`           // Judul buku.
	Author        string `json:"author"`          // Nama penulis buku.
	PublishedYear int    `json:"published_year"`  // Tahun terbit buku.
}
