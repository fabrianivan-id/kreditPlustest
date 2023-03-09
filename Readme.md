# Kredit Plus test
[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)
[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=white)](https://github.com/labstack/echo)
[![Go Reference](https://img.shields.io/badge/midtrans-reference-blue?logo=Midtrans&logoColor=white)](https://github.com/Midtrans/midtrans-go)
[![Go Reference](https://img.shields.io/badge/gmaps-reference-blue?logo=GMaps&logoColor=white)](https://github.com/googlemaps/google-maps-services-go)

# Arsitektur Aplikasi:
Untuk memenuhi kebutuhan bisnis yang telah diberikan, arsitektur aplikasi yang direkomendasikan adalah menggunakan arsitektur mikroservis. Dalam arsitektur ini, setiap layanan berdiri sendiri dan memiliki tugas dan tanggung jawab yang jelas. Setiap layanan memiliki database sendiri dan berkomunikasi dengan layanan lainnya melalui protokol yang telah ditentukan.

Dalam studi kasus ini, beberapa layanan yang dapat diidentifikasi adalah:

Layanan pengajuan peminjaman
Layanan manajemen limit pengajuan peminjaman
Layanan manajemen data konsumen
Layanan manajemen transaksi
Layanan manajemen keamanan aplikasi (opsi)


# Struktur Database:
Struktur database yang direkomendasikan untuk studi kasus ini adalah menggunakan pendekatan database per layanan. Setiap layanan memiliki database sendiri yang hanya berisi tabel yang diperlukan untuk menjalankan layanan tersebut.

Sebagai contoh, layanan manajemen data konsumen dapat memiliki tabel konsumen yang berisi data personal konsumen seperti NIK, Full name, Legal name, Tempat Lahir, Tanggal Lahir, Gaji, Foto KTP, dan Foto Selfie. 
Layanan manajemen transaksi dapat memiliki tabel transaksi yang berisi data transaksi seperti Nomor Kontrak, OTR, Admin Fee, Jumlah Cicilan, Jumlah Bunga, dan Nama Asset.

Dalam struktur database, ACID dapat dipenuhi dengan menggunakan teknologi database yang mendukung transaksi seperti PostgreSQL atau MySQL.

![ERD](https://github.com/fabrianivan21/kreditPlustest/blob/main/ER%20diagram.png)
