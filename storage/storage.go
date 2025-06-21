package storage

import (
	"gochi-book-api/models"
	"sync"
)

// BookStore adalah struct yang menyimpan data buku secara in-memory.
// Menggunakan slice dan lastID untuk mengelola data buku.
type BookStore struct {
	mu     sync.Mutex    // Mutex untuk menjamin keamanan data saat diakses secara bersamaan
	Books  []models.Book // Slice untuk menyimpan daftar buku
	lastID int           // ID terakhir yang digunakan untuk buku
}

var (
	once     sync.Once  // Untuk memastikan hanya satu instance yang dibuat
	instance *BookStore // Singleton instance
)

// GetBookStore mengembalikan instance tunggal (singleton) dari BookStore.
// Fungsi ini memastikan bahwa hanya ada satu BookStore yang digunakan selama aplikasi berjalan.
func GetBookStore() *BookStore {
	once.Do(func() {
		instance = &BookStore{
			Books: []models.Book{
				{ID: 1, Title: "Golang Chi Framework", Author: "Fattahillah", PublishedYear: 2024},
				{ID: 2, Title: "Learning About Go-Chi", Author: "Cakra", PublishedYear: 2025},
			},
			lastID: 2,
		}
	})
	return instance
}

// AddBook menambahkan buku baru ke dalam store.
// Secara otomatis meningkatkan lastID (AutoIncrement) dan menyimpan data buku baru.
func (s *BookStore) AddBook(book models.Book) models.Book {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.lastID++
	book.ID = s.lastID
	s.Books = append(s.Books, book)
	return book
}

// GetAllBooks mengembalikan semua data buku yang ada di dalam store.
func (s *BookStore) GetAllBooks() []models.Book {
	return s.Books
}

// GetBookByID mencari dan mengembalikan data buku berdasarkan IDnya.
// Jika tidak ditemukan, mengembalikan nilai false.
func (s *BookStore) GetBookByID(id int) (models.Book, bool) {
	for _, b := range s.Books {
		if b.ID == id {
			return b, true
		}
	}
	return models.Book{}, false
}

// UpdateBook memperbarui data buku berdasarkan ID.
// Jika berhasil, data buku diperbarui dan nilai dikembalikan.
// Jika ID tidak ditemukan, mengembalikan false.
func (s *BookStore) UpdateBook(id int, updated models.Book) (models.Book, bool) {
	for i, b := range s.Books {
		if b.ID == id {
			updated.ID = id
			s.Books[i] = updated
			return updated, true
		}
	}
	return models.Book{}, false
}

// DeleteBook menghapus buku berdasarkan ID.
// Jika berhasil dihapus, mengembalikan true.
// Jika ID tidak ditemukan, mengembalikan false.
func (s *BookStore) DeleteBook(id int) bool {
	for i, b := range s.Books {
		if b.ID == id {
			s.Books = append(s.Books[:i], s.Books[i+1:]...)
			return true
		}
	}
	return false
}