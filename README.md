# 📚 Gochi Book API

Gochi Book API adalah aplikasi RESTful API sederhana untuk manajemen data buku menggunakan Go dan framework [Chi v5](https://github.com/go-chi/chi).  
Data disimpan secara in-memory menggunakan slice.

---

## Menjalankan Aplikasi

```bash
go run main.go
```
## Table Routes
| Method | Endpoint    | Deskripsi                         |
| ------ | ----------- | --------------------------------- |
| GET    | /books      | Menampilkan semua buku            |
| GET    | /books/{id} | Menampilkan buku berdasarkan ID   |
| POST   | /books      | Menambahkan buku baru             |
| PUT    | /books/{id} | Mengubah data buku berdasarkan ID |
| DELETE | /books/{id} | Menghapus buku berdasarkan ID     |
