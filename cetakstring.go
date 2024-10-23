//103062300084 - Kisya Reinatana Yama
package main // Mulai program

import "fmt" // untuk memanggil program, input output data

func main()  { // fungsi untuk menjalankan program
	var z int // deklarasi variabel z bilangan bulat non-negatif dengan type integer 
	var k string // deklarasi variabel k kata yang diulang dengan type string 

	fmt.Println("Perulangan Kata") // cetak kalimat Perulangan Kata
	fmt.Println("===============================") // cetak simbol =
	fmt.Print("Masukan Jumlah Kata yang akan diulang :") // cetak kalimat Masukan Jumlah Kata yang akan diulang
	fmt.Scan(&z) // masukan atau input NILAI kata yang akan diulang
	fmt.Print("Masukan Kata yang akan diulang :") // cetak kalimat Masukan Kata yang akan diulang
	fmt.Scan(&k) // masukan atau input KATA yang akan diulang

	for i := 1; i < z; i++ { // looping atau batas pengulangan yang akan dijalankan program
		fmt.Println(k) // cetak hasil dari banyak kata yang di minta 
	}
	
} // program selesai 
