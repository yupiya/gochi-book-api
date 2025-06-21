package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gochi-book-api/handlers"
	"gochi-book-api/middlewares"
)

// Mengatur dan mengembalikan router utama aplikasi.
// Router ini menggunakan chi v5 dan mengatur middleware serta rute endpoint /books.
func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// Middleware Recoverer untuk menangani agar tidak crash.
	r.Use(middleware.Recoverer)

	// Middleware custom untuk mencatat setiap request (method, path, response time).
	r.Use(middlewares.LoggerMiddleware)

	// Group route untuk endpoint /books.
	r.Route("/books", func(r chi.Router) {
		r.Get("/", handlers.GetBooks)           // GET semua buku.
		r.Post("/", handlers.CreateBook)        // POST tambah buku baru.
		r.Get("/{id}", handlers.GetBookByID)    // GET buku berdasarkan ID.
		r.Put("/{id}", handlers.UpdateBook)     // PUT (update) buku berdasarkan ID.
		r.Delete("/{id}", handlers.DeleteBook)  // DELETE buku berdasarkan ID.
	})

	return r
}
