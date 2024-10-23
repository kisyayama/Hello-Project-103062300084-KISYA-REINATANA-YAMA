package main //mulai program

import (
    "fmt" //memanggil program input output data
    "sort" //memanggil program pengurutan data
)

//menggunakan tipe data bentukan (struct) untuk mendefinisikan struktur data barang dan transaksi
type Barang struct { //definisi struct Barang untuk menyimpan informasi barang
    ID       int //deklarasi ID dengan type integer bilangan bulat
    Nama     string //deklarasi Nama barang dengan type string
    Kategori string //deklarasi Kategori dengan type string
    Harga    float64 //deklarasi Harga dengan type float64 bilangan berkoma
    Stok     int //deklarasi Stok dengan type integer bilangan bulat
}

type Transaksi struct { //definisi struct Transaksi untuk menyimpan informasi transaksi
    ID        int //deklarasi ID dengan type integer bilangan bulat
    BarangID  int //deklarasi Barang ID dengan type integer bilangan bulat
    Jumlah    int //deklarasi Jumlah dengan type integer bilangan bulat
    HargaJual float64 //deklarasi Harga Jual dengan type float64 bilangan berkoma
}

//penggunaan array statis yaitu ukurannya ditentukan deklarasi dan tidak bisa diubah selama runtime
//menggunakan array untuk menyimpan data barang dan transaksi dalam program
var barangList [100]Barang //deklarasi array yang hanya menyimpan max 100 elemen bertipe barang
var transaksiList [100]Transaksi //deklarasi array yang hanya menyimpan max 100 elemen bertipe transaksi
var barangCount int //deklarai variabel barangCount dengan type integer bilangan bulat
var transaksiCount int //deklarai variabel transaksiCount dengan type integer bilangan bulat

//Insertion sort digunakan didalam fungsi tambah barang untuk memastikan bahwa setiap kali sebuah barang baru ditambahkan tetap terurut berdasarkan ID.
func tambahBarang(barang Barang) { //fungsi untuk menambah barang baru
    //mengurutkan daftar barang berdasarkan ID
    sort.Slice(barangList[:barangCount], func(i, j int) bool {
        return barangList[i].ID < barangList[j].ID
    })

    // Mencari posisi untuk memasukkan barang baru menggunakan binary search
    left, right := 0, barangCount-1
    for left <= right {
        mid := (left + right) / 2
        if barangList[mid].ID == barang.ID {
            fmt.Println("ID barang sudah ada. Gunakan ID yang berbeda.")
            fmt.Print("ID: ")
            fmt.Scanln(&barang.ID)
            left, right = 0, barangCount-1 // Reset the search range
            continue
        } else if barangList[mid].ID < barang.ID {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    // Menambah barang baru jika masih ada kapasitas
    if barangCount < len(barangList) {
        barangList[barangCount] = barang
        barangCount++
    } else {
        fmt.Println("Kapasitas maksimum barang telah tercapai")
    }
}

func editBarang(id int, barang Barang) { // Fungsi untuk mengedit data barang berdasarkan ID
    index := binarySearchBarang(id) // Mencari index barang menggunakan binary search
    for index == -1 {
        // Jika ID barang tidak ditemukan, meminta ID baru
        fmt.Println("ID barang tidak ditemukan. Masukkan ID yang sudah ada: ")
        fmt.Scanln(&id)
        index = binarySearchBarang(id)
    }

    barangList[index] = barang // Mengupdate barang pada index yang ditemukan

}

func hapusBarang(id int) { // Fungsi untuk menghapus barang berdasarkan ID
    index := binarySearchBarang(id) // Mencari index barang menggunakan binary search
    if index != -1 {
        // Menggeser elemen-elemen setelah index yang dihapus
        copy(barangList[index:], barangList[index+1:barangCount])
        barangCount--
    } else {
        fmt.Println("Barang tidak ditemukan")
    }
}

func tambahTransaksi(transaksi Transaksi) { // Fungsi untuk menambah transaksi baru
    if transaksiCount < len(transaksiList) {
        transaksiList[transaksiCount] = transaksi
        transaksiCount++
    } else {
        fmt.Println("Kapasitas maksimum transaksi telah tercapai")
    }
}

func editTransaksi(id int, transaksi Transaksi) { // Fungsi untuk mengedit data transaksi berdasarkan ID
    for i := 0; i < transaksiCount; i++ {
        if transaksiList[i].ID == id {
            // Meminta nama barang baru untuk diupdate
            fmt.Print("Masukkan nama barang baru: ")
            var namaBarang string
            fmt.Scanln(&namaBarang)
            transaksi.BarangID = cariBarangID(namaBarang)
            if transaksi.BarangID == -1 {
                fmt.Println("Barang tidak ditemukan")
                return
            }
            transaksiList[i] = transaksi
            return
        }
    }
    fmt.Println("Transaksi tidak ditemukan")
}

func hapusTransaksi(id int) { // Fungsi untuk menghapus transaksi berdasarkan ID
    for i := 0; i < transaksiCount; i++ {
        if transaksiList[i].ID == id {
            // Menggeser elemen-elemen setelah index yang dihapus
            copy(transaksiList[i:], transaksiList[i+1:transaksiCount])
            transaksiCount--
            return
        }
    }
    fmt.Println("Transaksi tidak ditemukan")
}

//menggunakan binary search untuk mencari index dari suatu barang berdasarkan ID nya
func binarySearchBarang(id int) int { // Fungsi untuk melakukan binary search pada barang berdasarkan ID
    left, right := 0, barangCount-1
    for left <= right {
        mid := (left + right) / 2
        if barangList[mid].ID == id {
            return mid
        } else if barangList[mid].ID < id {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}

func lihatBarang(urutkanBerdasarkan, urutan string) { // Fungsi untuk melihat daftar barang yang diurutkan berdasarkan kriteria tertentu
    validCriteria := map[string]bool{ // Validasi kriteria urutan
        "nama":     true,
        "kategori": true,
        "harga":    true,
        "stok":     true,
    }

    if !validCriteria[urutkanBerdasarkan] {
        fmt.Println("Urutan tidak valid. Gunakan salah satu dari: nama, kategori, harga, stok.")
        return
    }


    // Mengurutkan barang berdasarkan kriteria dan urutan yang dipilih
    sort.Slice(barangList[:barangCount], func(i, j int) bool {
        switch urutkanBerdasarkan {
        case "nama":
            if urutan == "ascending" {
                return barangList[i].Nama < barangList[j].Nama
            }
            return barangList[i].Nama > barangList[j].Nama
        case "kategori":
            if urutan == "ascending" {
                return barangList[i].Kategori < barangList[j].Kategori
            }
            return barangList[i].Kategori > barangList[j].Kategori
        case "harga":
            if urutan == "ascending" {
                return barangList[i].Harga < barangList[j].Harga
            }
            return barangList[i].Harga > barangList[j].Harga
        case "stok":
            if urutan == "ascending" {
                return barangList[i].Stok < barangList[j].Stok
            }
            return barangList[i].Stok > barangList[j].Stok
        }
        return false
    })

    for i := 0; i < barangCount; i++ { // Mencetak daftar barang yang telah diurutkan
        fmt.Println(barangList[i])
    }
}

func cariBarang(keyword string) { // Fungsi untuk mencari barang berdasarkan nama
    found := false
    for i := 0; i < barangCount; i++ {
        if barangList[i].Nama == keyword {
            fmt.Println(barangList[i])
            found = true
        }
    }
    if !found {
        fmt.Println("Barang tidak ditemukan")
    }
}

func laporanKeuangan() { // Fungsi untuk membuat laporan keuangan
    var modal, pendapatan float64
    var modalDetails, pendapatanDetails string

    // Menghitung total modal berdasarkan harga dan stok barang
    for i := 0; i < barangCount; i++ {
        itemModal := barangList[i].Harga * float64(barangList[i].Stok)
        modal += itemModal
        if i == 0 {
            modalDetails += fmt.Sprintf("(%.2f * %d)", barangList[i].Harga, barangList[i].Stok)
        } else {
            modalDetails += fmt.Sprintf(" + (%.2f * %d)", barangList[i].Harga, barangList[i].Stok)
        }
    }

    // Menghitung total pendapatan berdasarkan transaksi
    for i := 0; i < transaksiCount; i++ {
        itemPendapatan := transaksiList[i].HargaJual * float64(transaksiList[i].Jumlah)
        pendapatan += itemPendapatan
        if i == 0 {
            pendapatanDetails += fmt.Sprintf("(%.2f * %d)", transaksiList[i].HargaJual, transaksiList[i].Jumlah)
        } else {
            pendapatanDetails += fmt.Sprintf(" + (%.2f * %d)", transaksiList[i].HargaJual, transaksiList[i].Jumlah)
        }
    }

    // Mencetak detail perhitungan modal dan pendapatan
    fmt.Printf("Hitung Modal: %s = %.2f\n", modalDetails, modal)
    fmt.Printf("Hitung Pendapatan : %s = %.2f\n", pendapatanDetails, pendapatan)
    fmt.Printf("Modal: %.2f, Pendapatan: %.2f\n", modal, pendapatan)
}

func barangPalingBanyakTerjual() { // Fungsi untuk menemukan barang yang paling banyak terjual
    type penjualan struct {
        BarangID int
        Jumlah   int
    }

    // Map untuk menyimpan jumlah penjualan setiap barang
    penjualanMap := make(map[int]int)
    for i := 0; i < transaksiCount; i++ {
        penjualanMap[transaksiList[i].BarangID] += transaksiList[i].Jumlah
    }

    // Membuat slice dari map untuk diurutkan
    var penjualanList []penjualan
    for id, jumlah := range penjualanMap {
        penjualanList = append(penjualanList, penjualan{id, jumlah})
    }

    // Mengurutkan berdasarkan jumlah penjualan
    sort.Slice(penjualanList, func(i, j int) bool {
        return penjualanList[i].Jumlah > penjualanList[j].Jumlah
    })

    // Mencetak barang yang paling banyak terjual
    fmt.Println("Barang yang paling banyak terjual:")
    for i := 0; i < 5 && i < len(penjualanList); i++ {
        for _, barang := range barangList {
            if barang.ID == penjualanList[i].BarangID {
                fmt.Println(barang)
                break
            }
        }
    }
}

func cariBarangID(nama string) int { // Fungsi untuk mencari ID barang berdasarkan nama
    for i := 0; i < barangCount; i++ {
        if barangList[i].Nama == nama {
            return barangList[i].ID
        }
    }
    return -1
}

func lihatTransaksi() { // Fungsi untuk melihat daftar transaksi

    //menggunakan sequential search untuk mencari nama barang yang sesuai dengan ID barang dalam array
    for i := 0; i < transaksiCount; i++ {
        var namaBarang string
        for j := 0; j < barangCount; j++ {
            if barangList[j].ID == transaksiList[i].BarangID {
                namaBarang = barangList[j].Nama
                break
            }
        }

        if namaBarang != "" {
            fmt.Printf("ID: %d, Nama Barang: %s, Jumlah: %d, Harga Jual: %.2f\n", transaksiList[i].ID, namaBarang, transaksiList[i].Jumlah, transaksiList[i].HargaJual)
        } else {
            fmt.Printf("ID: %d, Nama Barang: Tidak Ditemukan, Jumlah: %d, Harga Jual: %.2f\n", transaksiList[i].ID, transaksiList[i].Jumlah, transaksiList[i].HargaJual)
        }
    }
}

func main() { // Fungsi utama yang berfungsi sebagai menu program
    var choice int //deklarasi variabel choise tengan type integer bilangan bulat

    //mencetak semua kalimat didalam kutip
    for {
        fmt.Println("\nMenu:")
        fmt.Println("1. Tambah Barang")
        fmt.Println("2. Edit Barang")
        fmt.Println("3. Hapus Barang")
        fmt.Println("4. Tambah Transaksi")
        fmt.Println("5. Lihat Transaksi")
        fmt.Println("6. Edit Transaksi")
        fmt.Println("7. Hapus Transaksi")
        fmt.Println("8. Lihat Barang")
        fmt.Println("9. Cari Barang")
        fmt.Println("10. Laporan Keuangan")
        fmt.Println("11. Barang Paling Banyak Terjual")
        fmt.Println("12. Keluar")
        fmt.Print("Pilih menu: ")
        fmt.Scanln(&choice)

        // Memproses pilihan pengguna
        switch choice {
        case 1:
            var barang Barang //deklarasi variabel barang yaitu Barang
            fmt.Print("Masukkan ID barang: ") //cetak kalimat dalam kutip
            fmt.Scanln(&barang.ID) //menginput ID dari pengguna
            fmt.Print("Masukkan nama barang: ") //cetak kalimat dalam kutip
            fmt.Scanln(&barang.Nama) //menginput nama barang dari pengguna
            fmt.Print("Masukkan kategori barang: ") //cetak kalimat dalam kutip
            fmt.Scanln(&barang.Kategori) //menginput kategori
            fmt.Print("Masukkan harga barang: ") //cetak kalimat dalam kutip
            fmt.Scanln(&barang.Harga) //menginput harga
            fmt.Print("Masukkan stok barang: ") //cetak kalimat dalam kutip
            fmt.Scanln(&barang.Stok) //menginput stok
            tambahBarang(barang)
        case 2:
            var id int //deklarasi variable id dengan type integer bilangan
            var barang Barang //deklarasi variable barang yaitu Barang
            fmt.Print("Masukkan ID barang yang akan diubah: ") //cetak kalimat dalam kutip
            fmt.Scanln(&id) //print atau input id dari pengguna
            fmt.Print("Masukkan ID baru barang: ") //cetak kalimat dalam kutip
            fmt.Scanln(&barang.ID) //input ID dari pengguna
            fmt.Print("Masukkan nama baru barang: ") //cetak kalimat dalam kutip
            fmt.Scanln(&barang.Nama) //input nama dari pengguna
            fmt.Print("Masukkan kategori baru barang: ") //cetak kalimat dalam kutip
            fmt.Scanln(&barang.Kategori) //input kategori dari pengguna
            fmt.Print("Masukkan harga baru barang: ") //cetak kalimat dalam kutip
            fmt.Scanln(&barang.Harga) //input harga dari pengguna
            fmt.Print("Masukkan stok baru barang: ") //cetak kalimat dalam kutip
            fmt.Scanln(&barang.Stok) //input stok dari pengguna
            editBarang(id, barang)
        case 3:
            var id int
            fmt.Print("Masukkan ID barang yang akan dihapus: ") //cetak kalimat dalam kutip
            fmt.Scanln(&id) //input id dari pengguna
            hapusBarang(id)
        case 4:
            var transaksi Transaksi //deklarasi variablel transaksi yaitu Transaksi
            fmt.Print("Masukkan ID transaksi: ") //cetak kalimat dalam kutip
            fmt.Scanln(&transaksi.ID) //input ID dari pengguna
            var namaBarang string //deklarasi variabel namaBaranng dengan type string
            fmt.Print("Masukkan nama barang: ") //cetak kalimat dalam kutip
            fmt.Scanln(&namaBarang) //input nama baran dari pengguna
            transaksi.BarangID = cariBarangID(namaBarang)
            if transaksi.BarangID == -1 {
                fmt.Println("Barang tidak ditemukan") //cetak kalimat dalam kutip
                continue
            }
            fmt.Print("Masukkan jumlah barang: ") //cetak kalimat dalam kutip
            fmt.Scanln(&transaksi.Jumlah) //input jumlah barang transaksi dari pengguna
            fmt.Print("Masukkan harga jual barang: ") //cetak kalimat dalam kutip
            fmt.Scanln(&transaksi.HargaJual) //input harga jual dari pengguna
            tambahTransaksi(transaksi)
        case 5:
            lihatTransaksi()
        case 6:
            var id int //deklarasi variabel id dengan type integer bilangan bulat
            var transaksi Transaksi //deklarasi transaksi yaitu Transaksi
            fmt.Print("Masukkan ID transaksi yang akan diubah: ") //cetak kalimat dalam kutip
            fmt.Scanln(&id) //input id yang akan diubah dari pengguna
            fmt.Print("Masukkan ID baru transaksi: ") //cetak kalimat dalam kutip
            fmt.Scanln(&transaksi.ID) //input id baru dari pengguna
            fmt.Print("Masukkan nama barang baru: ") //cetak kalimat dalam kutip
            var namaBarang string //deklarasi nama barang variabel dengan string
            fmt.Scanln(&namaBarang) //input nama barang dari pengguna
            transaksi.BarangID = cariBarangID(namaBarang)
            if transaksi.BarangID == -1 {
                fmt.Println("Barang tidak ditemukan") //cetak kalimat dalam kutip
                continue
            }
            fmt.Print("Masukkan jumlah baru barang: ") //cetak kalimat dalam kutip
            fmt.Scanln(&transaksi.Jumlah) //input jumlah barang baru dari pengguna
            fmt.Print("Masukkan harga jual baru barang: ") //cetak kalimat dalam kutip
            fmt.Scanln(&transaksi.HargaJual) //input harga jual  dari pengguna
            editTransaksi(id, transaksi)
        case 7:
            var id int
            fmt.Print("Masukkan ID transaksi yang akan dihapus: ") //cetak kalimat dalam kutip
            fmt.Scanln(&id) //input id dari pengguna
            hapusTransaksi(id)
        case 8:
            var urutkanBerdasarkan, urutan string
            fmt.Print("Urutkan berdasarkan (nama/kategori/harga/stok): ") //cetak kalimat dalam kutip
            fmt.Scanln(&urutkanBerdasarkan) //input urutkan berdasarkan (nama,kategori,harga,stok) dari pengguna
            fmt.Print("Urutan (ascending/descending): ") //cetak kalimat dalam kutip
            fmt.Scanln(&urutan) //input urutan (ascending/descending) dari pengguna
            lihatBarang(urutkanBerdasarkan, urutan)
        case 9:
            var keyword string
            fmt.Print("Masukkan nama barang yang dicari: ") //cetak kalimat dalam kutip
            fmt.Scanln(&keyword) //input namabarang yang dicari dari pengguna
            cariBarang(keyword)
        case 10:
            laporanKeuangan()
        case 11:
            barangPalingBanyakTerjual()
        case 12:
            fmt.Println("Keluar dari program.") //cetak kalimat dalam kutip
            return
        default:
            fmt.Println("Pilihan tidak valid. Silakan coba lagi.") //cetak kalimat dalam kutip
        }
    }
} //program selesai
