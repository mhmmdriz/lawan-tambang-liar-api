# Lawan Tambang Liar API

## About Project
Lawan Tambang Liar API adalah sebuah project API(Back-End) yang dibuat untuk digunakan di aplikasi Lawan Tambang Liar. Lawan Tambang Liar adalah sebuah aplikasi yang dibuat untuk memudahkan masyarakat dalam melaporkan keberadaan tambang liar di Provinsi Kepulauan Bangka Belitung. Tak hanya itu, kehadiran aplikasi ini juga diharapkan dapat menjadi alat bantu bagi pemerintah dalam menangani permasalahan tambang ilegal yang merugikan lingkungan serta kesejahteraan masyarakat di Provinsi Kepulauan Bangka Belitung. Melalui kolaborasi antara masyarakat dan pemerintah, diharapkan dapat tercipta sinergi yang kuat dalam menangani masalah ini demi menjaga keberlanjutan ekosistem dan kesejahteraan sosial.

## Features
### Super Admin
- Super Admin dapat melakukan login
- Super Admin dapat mengelola akun Admin
- Super Admin dapat mereset password akun Admin
- Super Admin dapat mengubah password akunnya

### Admin
- Admin dapat melakukan login
- Admin dapat mengelola akun User
- Admin dapat mereset password akun User
- Admin dapat mengubah password akunnya
- Admin dapat melakukan seeding data regency dan district dari API external
- Admin dapat melihat data regency dan district yang telah di seeding
- Admin dapat melihat seluruh data laporan yang telah dibuat oleh User
- Admin dapat melakukan searching, filtering, dan sorting data laporan yang telah dibuat oleh User
- Admin dapat melihat detail data laporan yang telah dibuat oleh User
- Admin dapat melihat estimasi waktu dan jarak dari lokasi Admin ke lokasi laporan tambang liar
- Admin dapat melakukan verifikasi laporan yang telah dibuat oleh User
- Admin dapat menambahkan progress penyelesaian laporan yang telah dibuat oleh User
- Admin dapat menyelesaikan progress laporan yang telah dibuat oleh User
- Admin dapat mengedit dan menghapus data proses penyelesaian laporan
- Admin dapat menghapus laporan yang telah dibuat oleh User
- Admin dapat meminta rekomendasi pesan proses penyelesaian laporan dari AI ketika ingin menambahkan proses penyelesaian laporan

### User
- User dapat melakukan registrasi
- User dapat melakukan login
- User dapat mengubah password akunnya
- User dapat melihat data regency dan district untuk membuat laporan
- User dapat membuat laporan tambang liar
- User dapat melihat seluruh data laporan yang telah dibuat oleh User lain
- User dapat melakukan searching, filtering, dan sorting data laporan
- User dapat melihat detail data laporan
- User dapat mengedit data laporan yang telah dibuat
- User dapat menghapus laporan yang telah dibuat
- User dapat melihat proses penyelesaian laporan yang dilakukan Admin
- User dapat meminta rekomendasi deskripsi laporan dari AI ketika ingin membuat laporan

## Tech Stacks
sebutkan daftar tools dan framework yang digunakan dalam bentuk list seperti ini:
- App Framework : Echo
- ORM Library : GORM
- DB : MySQL
- Storage : Google Cloud Platform, Cloud Storage
- Deployment : Google Cloud Platform, Cloud Run
- Code Structure : Clean Architecture
- Authentication : JWT
- External API : OpenAI API, Google Maps API, Indonesia Area API

## API Documentation
https://documenter.getpostman.com/view/31634961/2sA3JM6gYF

## ERD
![ERD](https://storage.googleapis.com/lawan-tambang-liar-assets/ERD.png)
Alternatif link ERD : https://bit.ly/erd-lawan-tambang-liar-api

## Setup 
Cara setup project :
1. Clone repository ini
2. Buka terminal dan arahkan ke folder project
3. Download golang jika belum terinstall di komputer anda
4. Jalankan perintah `go mod tidy` untuk mendownload semua dependencies yang dibutuhkan
5. Copy file `.env.example` dan rename menjadi `.env`
6. Sesuaikan konfigurasi yang ada di file `.env`
7. Jalankan perintah `go run main.go` untuk menjalankan project