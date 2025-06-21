package middlewares

import (
	"log"
	"net/http"
	"time"
)

// Middleware ini mencatat metode HTTP, path yang diakses, dan waktu eksekusi handler.
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() // Waktu mulai eksekusi handler.

		next.ServeHTTP(w, r)

		// Hitung durasi dan cetak log setelah request selesai diproses.
		duration := time.Since(start)
		log.Printf("[%s] %s - %v", r.Method, r.URL.Path, duration)
	})
}
