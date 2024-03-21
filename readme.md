# Super Indo Product API

API ini digunakan untuk menambahkan, mengambil, dan mengelola data produk di Super Indo.

## Spesifikasi

-   Dapat melakukan penambahkan data produk.
-   Dapat menampilkan daftar data produk.
-   Dapat melakukan pencarian berdasarkan nama dan ID produk.
-   Dapat melakukan filter produk berdasarkan tipe produk: Sayuran, Protein, Buah, dan Snack.
-   Dapat melakukan sorting berdasarkan tanggal, harga, dan nama produk.

## Tech Stack

-   **Bahasa Pemrograman**: Go (Golang)
-   **Database**: SQL / NoSQL + Seeder + Migration
-   **Cache**: Redis
-   **Dependency Injection**: wire (Opsional)
-   **Unit Test**: Go Testing Framework (Opsional)
-   **Docker**: Containerization (Opsional)

## Daftar Endpoint

1. **[POST] /product**

    - Menambahkan data produk baru atau memperbarui data produk.
    - Jika ID produk sudah ada, maka akan memperbarui data produk.
    - Contoh body request:
        ```json
        {
            "name": "kokocrunch2",
            "typeId": 1,
            "price": 5700
        }
        ```

2. **[GET] /products**

    - Mendapatkan daftar semua produk.

3. **[GET] /product/{id}**

    - Mendapatkan detail produk berdasarkan ID.

4. **[GET] /products?type_id={typeId}**

    - Mendapatkan daftar produk berdasarkan tipeId.

5. **[GET] /products?name={keyword}**

    - Mencari produk berdasarkan nama.

6. **[GET] /products?sort={field}&dir={direction}**

    - Mengurutkan daftar produk berdasarkan tanggal, harga, atau nama.

7. **[GET] /product?typeId={typeId}&name={name}&sort={column}&dir={direction}**
    - Menggabungkan pencarian, filter, dan sorting.

Untuk dokumentasi lebih rinci, dapat diakses melalui endpoint `/api/swagger` atau melihat koleksi postman yang terdapat di dalam folder `docs`.

## Implementasi

1. Pastikan Anda telah menginstal Go dan Redis di sistem Anda.
2. Buatlah database sesuai kebutuhan, dan konfigurasikan koneksi database di file konfigurasi.
3. Pastikan Redis telah berjalan dan dikonfigurasi dengan benar.
4. Mulailah proyek dengan menjalankan `make run`.

### Catatan

Kredensial default untuk login admin adalah:

```json
{
    "usernameOrEmail": "admin",
    "password": "password"
}
```

Token dapat diakses melalui endpoint `[POST] /api/auth/login`.
