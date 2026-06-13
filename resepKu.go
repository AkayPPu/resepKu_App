package main
import (
	"fmt"
	"strings"
)

type Resep struct{
	title string

	nBahan int
	bahan [50]string
	
	nSteps int
	steps [50]string

	kategori string
	duration int
	searchCount int
}

func tampilkanTeks(teks string) {
	for i := 0; i < len(teks); i++ {
		if teks[i] == '_' {
			fmt.Print(" ")
		} else {
			fmt.Print(string(teks[i]))
		}
	}
}


func tambah(data []Resep) []Resep {
	var r Resep

	fmt.Println("===================")
	fmt.Println("  TAMBAHKAN RESEP  ")
	fmt.Println("===================")
	fmt.Println("Pastikan untuk menggunakan '_' sebagai pengganti Spasi!!!")
	fmt.Println()

	fmt.Print("Masukkan Judul : ")
	fmt.Scan(&r.title)

	fmt.Print("Jumlah Bahan : ")
	fmt.Scan(&r.nBahan)

	for i := 0; i < r.nBahan; i++ {
		fmt.Printf("Bahan %d : ", i+1)
		fmt.Scan(&r.bahan[i])
	}

	fmt.Print("Jumlah Langkah : ")
	fmt.Scan(&r.nSteps)

	for i := 0; i < r.nSteps; i++ {
		fmt.Printf("Langkah %d : ", i+1)
		fmt.Scan(&r.steps[i])
	}

	fmt.Print("Kategori : ")
	fmt.Scan(&r.kategori)

	fmt.Print("Durasi (menit) : ")
	fmt.Scan(&r.duration)

	r.searchCount = 0

	data = append(data, r)

	fmt.Println("----------------------------")
	fmt.Println("Resep berhasil ditambahkan!")
	fmt.Println("----------------------------")

	return data
}


func lihat(data []Resep){
	var pilih int
	var confirm string

	fmt.Println("=================")
	fmt.Println("   SEMUA RESEP   ")
	fmt.Println("=================")

	if len(data) == 0 {
		fmt.Println("Belum ada resep.")
		return
	}

	for i := 0; i < len(data); i++ {
		fmt.Print(i+1, ". ")
		tampilkanTeks(data[i].title)
		fmt.Println()
	}

	fmt.Println("=================")
	fmt.Print("Pilih Resep(1-", len(data), ") atau ketik '0' untuk kembali : ")
	fmt.Scan(&pilih)

	if pilih == 0 {
		return
	}

	if pilih < 1 || pilih > len(data) {
		fmt.Println("Nomor resep tidak valid!")
		return
	}

	fmt.Println("#####################")
	fmt.Print("    ",)
	tampilkanTeks(data[pilih-1].title)
	fmt.Println("        ")
	fmt.Println("#####################")
	fmt.Println("Bahan : ")
	for i := 0; i < data[pilih-1].nBahan; i++{
		fmt.Printf("%d. ", i+1)
		tampilkanTeks(data[pilih-1].bahan[i])
		fmt.Println()
		fmt.Println()
	}
	
	fmt.Println("Langkah :")
	for i := 0; i < data[pilih-1].nSteps; i++ {
		fmt.Printf("%d. ", i+1)
		tampilkanTeks(data[pilih-1].steps[i])
		fmt.Println()
	}

	fmt.Print("Kategori : ")
	tampilkanTeks(data[pilih-1].kategori)
	fmt.Println()

	fmt.Printf("Durasi : %d menit\n", data[pilih-1].duration)

	fmt.Println()
	fmt.Println("----------------------------")
	fmt.Print("Ketik 'back' untuk kembali :")
	fmt.Scan(&confirm)
	confirm = strings.ToUpper(confirm)
	fmt.Println("----------------------------")
	if confirm == "BACK" {
		lihat(data)
	}

}


func ubah(data []Resep) []Resep{
	var pilih, memilih int
	var confirm string

	fmt.Println("=================")
	fmt.Println("   SEMUA RESEP   ")
	fmt.Println("=================")

	if len(data) == 0 {
		fmt.Println("Belum ada resep.")
		return data
	}

	for i := 0; i < len(data); i++ {
		fmt.Print(i+1, ". ")
		tampilkanTeks(data[i].title)
		fmt.Println()
	}

	fmt.Println("=================")
	fmt.Print("Pilih Resep yang Ingin Diubah(1-", len(data), ") atau pilih '0' untuk kembali : ")
	fmt.Scan(&pilih)

	if pilih == 0{
		return data
	}

	if pilih < 1 || pilih > len(data) {
		fmt.Println("Nomor resep tidak valid!")
		return data
	}


	fmt.Println("#####################")
	fmt.Print("    ")
	tampilkanTeks(data[pilih-1].title)
	fmt.Println("        ")
	fmt.Println("#####################")
	fmt.Print("Apakah Kamu Ingin Mengubah Judul(y/n) : ")
	fmt.Scan(&confirm)
	if confirm == "Y" || confirm == "y"{
		fmt.Print("Judul Baru : ")
		fmt.Scan(&data[pilih-1].title)
	}

	fmt.Print("Apakah Kamu Ingin Mengubah Bahan(y/n) : ")
	fmt.Scan(&confirm)
	if confirm == "Y" || confirm == "y"{
		fmt.Println("Bahan :")
		for{
			for i := 0; i < data[pilih-1].nBahan; i++ {
			fmt.Printf("%d. ", i+1)
			tampilkanTeks(data[pilih-1].bahan[i])
			fmt.Println()
			}
			fmt.Print("Bahan Nomor Berapa yang Ingin Diubah(1-", data[pilih-1].nBahan, ") atau Ketik '0' untuk Menambah Bahan : ")
			fmt.Scan(&memilih)
			fmt.Println("-------------------------")

			if memilih == 0 {
				...


			}else if != 0 {
				if memilih < 1 || memilih > data[pilih-1].nBahan {
				fmt.Println("Nomor bahan tidak valid!")
				continue
				}

				fmt.Print("Ubah ", data[pilih-1].bahan[memilih-1], " Jadi : ")
				fmt.Scan(&data[pilih-1].bahan[memilih-1])
				fmt.Print("Masih Ingin Ubah Bahan(y/n) : ")
				fmt.Scan(&confirm)
			}

			if confirm == "n" || confirm == "N"{
				break
			}
		}
		
	}

	fmt.Print("Apakah Kamu Ingin Mengubah Steps Pembuatan(y/n) : ")
	fmt.Scan(&confirm)
	if confirm == "y" || confirm == "Y"{
		fmt.Println("Langkah :")
		for{
			for i := 0; i < data[pilih-1].nSteps; i++ {
				fmt.Printf("%d. ", i+1)
				tampilkanTeks(data[pilih-1].steps[i])
				fmt.Println()
			}
			fmt.Print("Steps Nomor Berapa yang Ingin Diubah(1-", data[pilih-1].nSteps, ") : ")
			fmt.Scan(&memilih)
			fmt.Println("-------------------------")
			fmt.Print("Ubah ", data[pilih-1].steps[memilih-1], " Jadi : ")
			fmt.Scan(&data[pilih-1].steps[memilih-1])
			fmt.Print("Masih Ingin Ubah Steps(y/n) : ")
			fmt.Scan(&confirm)
			if confirm == "n" || confirm == "N"{
				break
			}
		}
	}

	fmt.Print("Apakah Kamu Ingin Mengubah Kategori(y/n) : ")
	fmt.Scan(&confirm)
	if confirm == "Y" || confirm == "y"{
		fmt.Print("Kategori Baru : ")
		fmt.Scan(&data[pilih-1].kategori)
	}

	fmt.Print("Apakah Kamu Ingin Mengubah Durasi(y/n) : ")
	fmt.Scan(&confirm)
	if confirm == "Y" || confirm == "y"{
		fmt.Print("Durasi Baru : ")
		fmt.Scan(&data[pilih-1].duration)
	}

	return data
}


// func hapus(data []Resep) []Resep{

// }

// func cari(data []Resep){

// }

// func urut(data []Resep){

// }

// func statistik(data []Resep){

// }



func main(){
	var pilih int
	var data []Resep


	for{
		fmt.Println("==================")
		fmt.Println("      RESEPKU     ")
		fmt.Println("==================")
		fmt.Println("1. Lihat Resep")
		fmt.Println("2. Tambah Resep")
		fmt.Println("3. Ubah Resep")
		fmt.Println("4. Hapus Resep")
		fmt.Println("5. Cari Resep")
		fmt.Println("6. Urutkan Resep")
		fmt.Println("7. Statistik")
		fmt.Println("8. Keluar")
		fmt.Println("==================")
		fmt.Print("Pilih(1-8): ")
		fmt.Scan(&pilih)

		if pilih < 1 || pilih > 8 {
			fmt.Println("Pilihan tidak valid!")
			continue
		}

		if pilih == 8 {
			break
		}	


		if pilih == 1{
			lihat(data)
		}else if pilih == 2{
			data = tambah(data)
		}else if pilih == 3{
			data = ubah(data)
		}
		// else if pilih == 4{

		// }else if pilih == 5{

		// }else if pilih == 6 {
			
		// }else if pilih == 7{

		// }


	}


}