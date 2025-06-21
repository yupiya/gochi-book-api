# Gochi Book API

Gochi Book API adalah aplikasi RESTful API sederhana untuk manajemen data buku menggunakan Go dan framework [Chi v5](https://github.com/go-chi/chi).  
Data disimpan secara in-memory menggunakan slice.

---

## Menjalankan Aplikasi/Website

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

## Request via PowerShell

### Menampilkan Semua Data Buku

Untuk mendapatkan daftar semua buku, kirimkan permintaan **GET** ke `/books`.

```powershell
Invoke-RestMethod http://localhost:8080/books
```

### Menampilkan Data Buku Berdasarkan ID

Mendapatkan detail buku tertentu dengan menambahkan ID buku ke URL (Contoh ID=1).

```powershell
Invoke-RestMethod http://localhost:8080/books/1
```

### Menambah Data Buku
Dapat menambahkan buku baru dengan mengirimkan permintaan **POST** ke `/books`.

```powershell
$body = @{
  title = "Tere"
  author = "Yupi"
  published_year=2025
} | ConvertTo-Json -Depth 2

Invoke-RestMethod -Uri http://localhost:8080/books `
  -Method POST `
  -Body $body `
  -ContentType "application/json"
```

### Mengubah Data Buku Berdasarkan ID (Update)

Untuk memperbarui data buku yang sudah ada, kirimkan permintaan **PUT** ke `/books/{id}`.

```powershell
$body = @{
  title = "Tereliye"
  author = "Yupiyapiyay"
  published_year=2024
} | ConvertTo-Json

Invoke-RestMethod -Uri http://localhost:8080/books/3 `
  -Method PUT `
  -Body $body `
  -ContentType "application/json"
```

### Menghapus Data Buku Berdasarkan ID

Untuk menghapus buku, kirimkan permintaan **DELETE** ke `/books/{id}`.

```powershell
Invoke-RestMethod -Uri http://localhost:8080/books/2 -Method DELETE
```
