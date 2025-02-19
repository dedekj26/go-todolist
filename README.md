# Go Todo List API

Sebuah RESTful API sederhana untuk manajemen aktivitas/todo list yang dibangun menggunakan Go dan Fiber framework dengan PostgreSQL sebagai database.

## Teknologi yang Digunakan

- Go (Golang)
- Fiber v2 (Web Framework)
- PostgreSQL (Database)
- Supabase (Database Hosting)
- Go Validator v10 (Request Validation)

## Fitur

- Create Activity (Membuat aktivitas baru)
- Read Activities (Melihat semua aktivitas)
- Update Activity (Mengubah detail aktivitas)
- Delete Activity (Menghapus aktivitas)
- Filter aktivitas berdasarkan status
- Pencarian aktivitas berdasarkan judul

## Struktur Data Aktivitas

```json
{
  "id": 1,
  "title": "Judul Aktivitas",
  "category": "Kategori",
  "activity_date": "2025-02-19T00:00:00Z",
  "status": "pending",
  "created_at": "2025-02-19T00:00:00Z",
  "description": "Deskripsi aktivitas"
}
```

## Endpoint API

- `GET /activities` - Mendapatkan semua aktivitas
- `GET /activities/:id` - Mendapatkan detail aktivitas berdasarkan ID
- `POST /activities` - Membuat aktivitas baru
- `PUT /activities/:id` - Mengubah aktivitas yang ada
- `DELETE /activities/:id` - Menghapus aktivitas
- `GET /activities/search` - Mencari aktivitas berdasarkan judul
- `GET /activities/filter` - Memfilter aktivitas berdasarkan status

## Cara Menjalankan Aplikasi

1. Pastikan Go sudah terinstall di sistem Anda
2. Clone repository ini
3. Salin file `.env.example` menjadi `.env` dan sesuaikan konfigurasi database
4. Jalankan perintah:
   ```bash
   go mod download
   go run main.go
   ```
5. API akan berjalan di `http://localhost:3000`

## Pengembangan

Proyek ini menggunakan:
- Go Modules untuk manajemen dependensi
- Fiber v2 sebagai web framework
- PostgreSQL sebagai database
- Supabase sebagai hosting database

## Lisensi

MIT License
