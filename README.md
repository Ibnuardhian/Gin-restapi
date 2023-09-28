# Gin-restapi: API CRUD Produk

Gin-restapi adalah sebuah aplikasi RESTful API yang memungkinkan Anda untuk melakukan operasi CRUD (Create, Read, Update, Delete) terhadap data produk di dalam database. Aplikasi ini dibangun menggunakan framework web Go yang populer, yaitu Gin.

## Daftar Isi

- [Pengenalan](#pengenalan)
- [Fitur](#fitur)
- [Persyaratan](#persyaratan)
- [Cara Menggunakan](#cara-menggunakan)
- [Konfigurasi](#konfigurasi)
- [Endpoints](#endpoints)
- [Contoh Penggunaan](#contoh-penggunaan)

## Pengenalan

Aplikasi ini bertujuan untuk menyediakan sebuah API yang memungkinkan pengguna untuk mengelola data produk dalam sebuah database. Anda dapat menggunakan API ini untuk membuat, membaca, mengupdate, dan menghapus produk.

## Fitur

- Membuat produk baru beserta detailnya.
- Mengambil daftar semua produk.
- Mengambil detail produk berdasarkan ID.
- Mengupdate produk yang sudah ada.
- Menghapus produk berdasarkan ID.

## Persyaratan

Sebelum Anda dapat menggunakan aplikasi ini, pastikan Anda telah memenuhi persyaratan berikut:

- Go (versi >= 1.16)
- MySQL atau database lain yang mendukung Go (seperti PostgreSQL, SQLite, dsb.)
- Git (opsional, untuk mengklon repo)
- Postman atau perangkat serupa untuk menguji endpoint API.

## Cara Menggunakan

Berikut adalah langkah-langkah dasar untuk mengunduh, mengonfigurasi, dan menjalankan aplikasi ini:

1. **Klon Repo**: Klon repo ini ke komputer Anda dengan perintah berikut:

   ```bash
   git clone https://github.com/Ibnuardhian/Gin-restapi.git
   ```

2. **Konfigurasi Database**: Konfigurasi database Anda di dalam file `config/config.go`. Anda harus mengganti detail koneksi sesuai dengan database yang Anda gunakan.

3. **Instalasi Dependencies**: Jalankan perintah berikut untuk menginstal semua dependensi yang diperlukan:

   ```bash
   go mod tidy
   ```

4. **Jalankan Aplikasi**: Jalankan aplikasi dengan perintah:

   ```bash
   go run main.go
   ```

   Aplikasi akan berjalan di `http://localhost:8080` secara default. Anda dapat mengganti port tersebut di dalam file `main.go` jika diperlukan.

## Konfigurasi

Anda dapat mengkonfigurasi aplikasi ini dengan mengedit file `config/config.go`. Di sini, Anda dapat mengganti pengaturan database, port, dan konfigurasi lain yang diperlukan.

## Endpoints

Berikut adalah daftar endpoint yang tersedia:

- `GET /products`: Mendapatkan daftar semua produk.
- `GET /product/{id}`: Mendapatkan detail produk berdasarkan ID.
- `POST /product`: Membuat produk baru.
- `PUT /product/{id}`: Mengupdate produk berdasarkan ID.
- `DELETE /product/{id}`: Menghapus produk berdasarkan ID.

## Contoh Penggunaan

Anda dapat menggunakan Postman atau perangkat serupa untuk menguji endpoint-endpoint API. Berikut adalah contoh penggunaan beberapa endpoint:

- **Mendapatkan Daftar Produk**:

  ```
  GET http://localhost:8080/products
  ```

- **Membuat Produk Baru**:

  ```
  POST http://localhost:8080/products
  Content-Type: application/json

  {
      "name": "Nama Produk",
      "description": "Deskripsi Produk",
  }
  ```

- **Mengupdate Produk**:

  ```
  PUT http://localhost:8080/products/{id}
  Content-Type: application/json

  {
      "name": "Nama Produk yang Diperbarui",
      "description": "Deskripsi Produk yang Diperbarui",
  }
  ```

- **Menghapus Produk**:

  ```
  DELETE http://localhost:8080/products/{id}
  ```

Pastikan untuk mengganti `{id}` dengan ID produk yang sesuai.

Selamat menggunakan Gin-restapi!
