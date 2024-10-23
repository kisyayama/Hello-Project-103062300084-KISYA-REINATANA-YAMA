//103062300084 - Kisya Reinatana Yama
package main

import (
	"fmt"
)

type Mahasiswa struct {
	Nama    string
	NIM     string
	Jurusan string
	IPK     float64
}

func main() {
	// Membuat objek mahasiswa
	mahasiswa1 := Mahasiswa{
		Nama:    "Weka",
		NIM:     "1202224005",
		Jurusan: "Teknologi Informasi",
		IPK:     3.85,
	}

	mahasiswa2 := Mahasiswa{
		Nama:    "Yono",
		NIM:     "10301224009",
		Jurusan: "Sistem Informasi",
		IPK:     3.67,
	}

	mahasiswa3 := Mahasiswa{
		Nama:	"Lukman",
		NIM:	"10202334029",
		Jurusan:	"DKV",
		IPK:	3.92,
	}

	// Menampilkan data mahasiswa
	fmt.Println("Data Mahasiswa 1:")
	fmt.Println("Nama:", mahasiswa1.Nama)
	fmt.Println("NIM:", mahasiswa1.NIM)
	fmt.Println("Jurusan:", mahasiswa1.Jurusan)
	fmt.Println("IPK:", mahasiswa1.IPK)

	fmt.Println("\nData Mahasiswa 2:")
	fmt.Println("Nama:", mahasiswa2.Nama)
	fmt.Println("NIM:", mahasiswa2.NIM)
	fmt.Println("Jurusan:", mahasiswa2.Jurusan)
	fmt.Println("IPK:", mahasiswa2.IPK)

	fmt.Println("\nData Mahasiswa 3:")
	fmt.Println("Nama:", mahasiswa3.Nama)
	fmt.Println("NIM:", mahasiswa3.NIM)
	fmt.Println("Jurusan:", mahasiswa3.Jurusan)
	fmt.Println("IPK:", mahasiswa3.IPK)
	
}
