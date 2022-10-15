package main

// Import fmt untuk keperluan i/o
// Import os untuk menginput os.Args
// import strconv untuk mengconvert string ke integer
import (
	"fmt"
	"os"
	"strconv"
)

// Membuat struct yang menampung data teman
type Person struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// Fungsi untuk menunjukkan error ketika nomor absen tidak ditemukan
func showError() {
	fmt.Println("Maaf, Nomor Absen yang anda pilih tidak ditemukan")
}

// Fungsi untuk menunjukkan data ketika nomor absen ditemukan
func showData(teman Person) {
	fmt.Println("Nama \t: ", teman.Nama)
	fmt.Println("Alamat \t: ", teman.Alamat)
	fmt.Println("Pekerjaan : ", teman.Pekerjaan)
	fmt.Println("Alasan \t: ", teman.Alasan)
}
func main() {

	// Mendapatkan input argument
	var args = os.Args
	var nomor = args[1]

	// Convert string menjadi integer
	number, _ := strconv.Atoi(nomor)
	// Untuk menunjukkan input
	// fmt.Printf("-> %#v\n", number)

	// Data Teman kalian
	var dataTeman = []Person{
		{Nama: "Bastian", Alamat: "Ponorogo", Pekerjaan: "Mahasiswa", Alasan: "Ingin menggali ilmu"},
		{Nama: "Arfianto", Alamat: "Sidoarjo", Pekerjaan: "Mahasiswa", Alasan: "Ingin mencari teman"},
		{Nama: "Kadek", Alamat: "Denpasar", Pekerjaan: "Freelance", Alasan: "Memperbanyak relasi"},
		{Nama: "Laba", Alamat: "Badung", Pekerjaan: "Wiraswasta", Alasan: "Gabut aja"},
		{Nama: "Subagia", Alamat: "Singaraja", Pekerjaan: "Software Developer", Alasan: "Menambah relasi dan menambah teman"},
		{Nama: "Labib", Alamat: "Bogor", Pekerjaan: "Front End Developer", Alasan: "Ingin switch karir dari Back End ke Front End"},
		{Nama: "Ihsan", Alamat: "Bandung", Pekerjaan: "Back End Developer", Alasan: "Mengejar sertifikat"},
		{Nama: "Ramadhan", Alamat: "Sumedang", Pekerjaan: "Mechanical Engineer", Alasan: "Ingin menambah ilmu serta relasi"},
		{Nama: "Evan", Alamat: "Yogyakarta", Pekerjaan: "Mahasiswa", Alasan: "Ingin menambah ilmu dan wawasan"},
		{Nama: "Anindya", Alamat: "Sleman", Pekerjaan: "Pedagang", Alasan: "Menghabiskan waktu"},
		{Nama: "Hasan", Alamat: "Banjarmasin", Pekerjaan: "Mahasiswa", Alasan: "Disuruh dosen untuk ikutan"},
		{Nama: "Angga", Alamat: "Jakarta", Pekerjaan: "Pelajar", Alasan: "Persiapan untuk kuliah"},
	}

	// Cek data jika ada
	if number > len(dataTeman) {
		showError()
	} else {
		showData(dataTeman[number-1])
	}
}
