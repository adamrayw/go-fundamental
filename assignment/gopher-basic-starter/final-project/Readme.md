# Final Project Gopher-Basic

## Introduction
Pada final project ini teman-teman akan membuat sebuah aplikasi sederhana yang berfungsi untuk mengelola data mahasiswa. Aplikasi ini akan memiliki 4 fitur utama yaitu:
- Menambahkan data mahasiswa
- Menghapus data mahasiswa
- Menampilkan data mahasiswa
- Menampilkan data mahasiswa berdasarkan kriteria tertentu

## Project Structure
Project ini terdiri dari 3 package utama yaitu:
1. `main` package
2. `mahasiswa` package
3. `course` package

### `main` package
Package main berisi fungsi `main` yang akan menjalankan aplikasi. Pada package ini teman-teman bisa mencoba untuk menjalankan program nya.

### `mahasiswa` package
Package `mahasiswa` berisi fungsi-fungsi yang berhubungan dengan data mahasiswa. Pada package ini terdapat fungsi-fungsi berikut:
- `AddMahasiswa` untuk menambahkan data mahasiswa
- `DeleteMahasiswa` untuk menghapus data mahasiswa
- `GetMahasiswa` untuk menampilkan data mahasiswa
- `GetMahasiswaByCourse` untuk menampilkan data mahasiswa berdasarkan kriteria tertentu

### `course` package
Package `course` berisi fungsi-fungsi yang berhubungan dengan data mata kuliah. Pada package ini terdapat fungsi-fungsi berikut:
- `AddCourse` untuk menambahkan data mata kuliah
- `GetCourseByName` untuk menampilkan data mata kuliah berdasarkan nama mata kuliah

## Cara Mengecek Jawaban
- Terdapat file `course_test.go` dan `mahasiswa_test.go` yang berisi test case untuk mengecek jawaban teman-teman. Teman-teman bisa mengecek jawaban dengan menjalankan perintah `go test` pada terminal.
  - `go test ./course -v` untuk mengecek jawaban pada package `course` 
  - `go test ./mahasiswa -v` untuk mengecek jawaban pada package `mahasiswa`
  - Jangan lupa untuk masuk ke folder `final-project` terlebih dahulu sebelum menjalankan perintah di atas.
- Jika terdapat test case yang gagal, teman-teman bisa memperbaiki jawaban teman-teman sesuai dengan pesan error yang ditampilkan oleh test case tersebut.
- Jika semua test case berhasil, ditandai dengan tulisan `PASS`, teman-teman sudah berhasil menyelesaikan final project ini.

Selamat mengerjakan!