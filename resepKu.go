package main

import "fmt"

type Resep struct {
	title string

	nBahan int
	bahan  [50]string

	nSteps int
	steps  [50]string

	kategori    string
	duration    int
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

func lihat(data []Resep) {
	var pilih int
	var confirm string

	fmt.Println("=================")
	fmt.Println("   SEMUA RESEP   ")
	fmt.Println("=================")

	if len(data) == 0 {
		fmt.Println("Belum ada resep.")
		fmt.Println("=================")
		fmt.Println("Ketik '-2' untuk kembali ke Menu Utama")
		fmt.Print("Pilihan Anda : ")
		fmt.Scan(&pilih)
		if pilih == -2 {
			return
		}
		fmt.Println("Pilihan tidak valid!")
		lihat(data)
		return
	}

	for i := 0; i < len(data); i++ {
		fmt.Print(i+1, ". ")
		tampilkanTeks(data[i].title)
		fmt.Printf(" | Durasi : %d menit", data[i].duration)
		fmt.Println()
	}

	fmt.Println("=================")
	fmt.Println("Pilih nomor resep (1-",len(data),") untuk melihat detail")
	fmt.Println("Ketik '0'  -> Cari Resep")
	fmt.Println("Ketik '-1' -> Urutkan Resep")
	fmt.Println("Ketik '-2' -> Kembali ke Menu Utama")
	fmt.Println("=================")
	fmt.Print("Pilihan Anda : ")
	fmt.Scan(&pilih)

	if pilih == -2 {
		return
	}

	if pilih == -1 {
		var opsi int
		fmt.Println("===========================")
		fmt.Println("      URUTKAN RESEP        ")
		fmt.Println("===========================")
		fmt.Println("0. Kembali ke Menu Sebelumnya")
		fmt.Println("1. Berdasarkan Durasi Masak (Tercepat) - Selection Sort")
		fmt.Println("2. Berdasarkan Judul Resep (A-Z) - Insertion Sort")
		fmt.Println("===========================")
		fmt.Print("Pilih (0-2): ")
		fmt.Scan(&opsi)

		if opsi == 0 {
			lihat(data)
			return
		}

		if opsi == 1 {
			for i := 0; i < len(data)-1; i++ {
				n := i
				for j := i + 1; j < len(data); j++ {
					if data[j].duration < data[n].duration {
						n = j
					}
				}
				if n != i {
					data[i], data[n] = data[n], data[i]
				}
			}
			fmt.Println()
			fmt.Println("--- Resep berhasil diurutkan berdasarkan Durasi (Tercepat) ---")
			fmt.Println()
		} else if opsi == 2 {
			for i := 1; i < len(data); i++ {
				r := data[i]
				j := i - 1
				for j >= 0 && data[j].title > r.title {
					data[j+1] = data[j]
					j--
				}
				data[j+1] = r
			}
			fmt.Println()
			fmt.Println("--- Resep berhasil diurutkan berdasarkan Judul (A-Z) ---")
			fmt.Println()
		} else {
			fmt.Println("Pilihan tidak valid!")
		}

		lihat(data)
		return
	}

	if pilih == 0 {
		var menu int
		var kata string
		fmt.Println("-------------------------")
		fmt.Println("    METODE PENCARIAN     ")
		fmt.Println("-------------------------")
		fmt.Println("0. Kembali ke Menu Sebelumnya")
		fmt.Println("1. Cari Bahan Utama (Binary Search)")
		fmt.Println("2. Cari Bahan Utama (Sequential Search)")
		fmt.Println("-------------------------")
		fmt.Print("Pilih (0-2): ")
		fmt.Scan(&menu)

		if menu == 0 {
			lihat(data)
			return
		}

		if menu < 1 || menu > 2 {
			fmt.Println("Pilihan tidak valid!")
			lihat(data)
			return
		}

		fmt.Print("Masukkan Bahan Utama: ")
		fmt.Scan(&kata)

		var ketemu int = -1

		if menu == 1 {
			for i := 0; i < len(data)-1; i++ {
				for j := i + 1; j < len(data); j++ {
					if data[i].bahan[0] > data[j].bahan[0] {
						data[i], data[j] = data[j], data[i]
					}
				}
			}

			var low int = 0
			var high int = len(data) - 1
			for low <= high && ketemu == -1 {
				var mid int = (low + high) / 2
				if data[mid].bahan[0] == kata {
					ketemu = mid
				} else if data[mid].bahan[0] < kata {
					low = mid + 1
				} else {
					high = mid - 1
				}
			}
		} else if menu == 2 {
			for i := 0; i < len(data); i++ {
				if data[i].bahan[0] == kata {
					ketemu = i
					break
				}
			}
		}

		if ketemu == -1 {
			fmt.Println("-------------------------")
			fmt.Println("Resep tidak ditemukan!")
			fmt.Println("-------------------------")
			lihat(data)
			return
		}

		pilih = ketemu + 1
	}

	if pilih < 1 || pilih > len(data) {
		fmt.Println("Nomor resep tidak valid!")
		return
	}

	fmt.Println("#####################")
	fmt.Print("    ")
	tampilkanTeks(data[pilih-1].title)
	fmt.Println("        ")
	fmt.Println("#####################")
	fmt.Println("Bahan : ")
	for i := 0; i < data[pilih-1].nBahan; i++ {
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
	fmt.Print("Ketik 'back' untuk kembali : ")
	fmt.Scan(&confirm)
	fmt.Println("----------------------------")
	if confirm == "back" || confirm == "BACK" || confirm == "Back" || confirm == "BAck" || confirm == "BaCK" || confirm == "bACk" || confirm == "bacK" || confirm == "baCk" {
		lihat(data)
	}

}

func ubah(data []Resep) []Resep {
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

	if pilih == 0 {
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
	if confirm == "Y" || confirm == "y" {
		fmt.Print("Judul Baru : ")
		fmt.Scan(&data[pilih-1].title)
	}

	fmt.Print("Apakah Kamu Ingin Mengubah Bahan(y/n) : ")
	fmt.Scan(&confirm)
	if confirm == "Y" || confirm == "y" {
		fmt.Println("Bahan :")
		for {
			for i := 0; i < data[pilih-1].nBahan; i++ {
				fmt.Printf("%d. ", i+1)
				tampilkanTeks(data[pilih-1].bahan[i])
				fmt.Println()
			}
			fmt.Print("Bahan Nomor Berapa yang Ingin Diubah(1-", data[pilih-1].nBahan, ") atau Ketik '0' untuk Menambah Bahan : ")
			fmt.Scan(&memilih)
			fmt.Println("-------------------------")

			if memilih == 0 {
				if data[pilih-1].nBahan >= 50 {
					fmt.Println("Bahan sudah penuh (maksimal 50)!")
				} else {
					fmt.Print("Masukkan Bahan Baru: ")
					fmt.Scan(&data[pilih-1].bahan[data[pilih-1].nBahan])
					data[pilih-1].nBahan++
					fmt.Print("Masih Ingin Ubah/Tambah Bahan(y/n) : ")
					fmt.Scan(&confirm)
				}
			} else {
				if memilih < 1 || memilih > data[pilih-1].nBahan {
					fmt.Println("Nomor bahan tidak valid!")
					continue
				}

				fmt.Print("Ubah ", data[pilih-1].bahan[memilih-1], " Jadi : ")
				fmt.Scan(&data[pilih-1].bahan[memilih-1])
				fmt.Print("Masih Ingin Ubah Bahan(y/n) : ")
				fmt.Scan(&confirm)
			}

			if confirm == "n" || confirm == "N" {
				break
			}
		}

	}

	fmt.Print("Apakah Kamu Ingin Mengubah Steps Pembuatan(y/n) : ")
	fmt.Scan(&confirm)
	if confirm == "y" || confirm == "Y" {
		fmt.Println("Langkah :")
		for {
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
			if confirm == "n" || confirm == "N" {
				break
			}
		}
	}

	fmt.Print("Apakah Kamu Ingin Mengubah Kategori(y/n) : ")
	fmt.Scan(&confirm)
	if confirm == "Y" || confirm == "y" {
		fmt.Print("Kategori Baru : ")
		fmt.Scan(&data[pilih-1].kategori)
	}

	fmt.Print("Apakah Kamu Ingin Mengubah Durasi(y/n) : ")
	fmt.Scan(&confirm)
	if confirm == "Y" || confirm == "y" {
		fmt.Print("Durasi Baru : ")
		fmt.Scan(&data[pilih-1].duration)
	}

	return data
}

func hapus(data []Resep) []Resep {
	var pilih int
	var confirm string

	for {
		fmt.Println("=================")
		fmt.Println("   HAPUS RESEP   ")
		fmt.Println("=================")

		if len(data) == 0 {
			fmt.Println("Belum ada resep.")
			fmt.Println("=================")
			fmt.Println("Ketik '0' untuk kembali ke Menu Utama")
			fmt.Print("Pilihan Anda : ")
			fmt.Scan(&pilih)

			if pilih == 0 {
				return data
			}

			fmt.Println("Pilihan tidak valid!")
			continue
		}

		for i := 0; i < len(data); i++ {
			fmt.Print(i+1, ". ")
			tampilkanTeks(data[i].title)
			fmt.Printf(" | Durasi : %d menit", data[i].duration)
			fmt.Println()
		}

		fmt.Println("=================")
		fmt.Print("Resep nomor berapa yang ingin dihapus (1-", len(data), ") atau ketik '0' untuk kembali ke menu utama : ")
		fmt.Scan(&pilih)

		if pilih == 0 {
			return data
		}

		if pilih < 1 || pilih > len(data) {
			fmt.Println("Nomor resep tidak valid!")
			continue
		}

		fmt.Println()
		fmt.Println("-------------------------")
		fmt.Print("Yakin ingin menghapus resep '")
		tampilkanTeks(data[pilih-1].title)
		fmt.Print("'? (y/n) : ")
		fmt.Scan(&confirm)

		if confirm == "y" || confirm == "Y" {

			fmt.Print("Resep '")
			tampilkanTeks(data[pilih-1].title)
			fmt.Println("' berhasil dihapus!")

			// Geser semua elemen setelahnya ke kiri
			for i := pilih - 1; i < len(data)-1; i++ {
				data[i] = data[i+1]
			}

			// Kurangi panjang slice
			data = data[:len(data)-1]

			return data

		} else if confirm == "n" || confirm == "N" {

			fmt.Println("Penghapusan dibatalkan.")
			continue

		} else {

			fmt.Println("Input tidak valid!")
			fmt.Println("Penghapusan dibatalkan.")
			continue

		}
	}
}

func statistik(data []Resep) {
	if len(data) == 0 {
		fmt.Println("=========================")
		fmt.Println("        STATISTIK")
		fmt.Println("=========================")
		fmt.Println("Belum ada resep.")
		fmt.Println("=========================")
		return
	}

	var kategori []string
	var jumlah []int

	for i := 0; i < len(data); i++ {
		ketemu := -1

		for j := 0; j < len(kategori); j++ {
			if kategori[j] == data[i].kategori {
				ketemu = j
				break
			}
		}

		if ketemu != -1 {
			jumlah[ketemu]++
		} else {
			kategori = append(kategori, data[i].kategori)
			jumlah = append(jumlah, 1)
		}
	}

	idxMax := 0
	idxMin := 0

	for i := 1; i < len(jumlah); i++ {
		if jumlah[i] > jumlah[idxMax] {
			idxMax = i
		}

		if jumlah[i] < jumlah[idxMin] {
			idxMin = i
		}
	}

	top := make([]Resep, len(data))
	copy(top, data)

	for i := 0; i < len(top)-1; i++ {
		max := i

		for j := i + 1; j < len(top); j++ {
			if top[j].searchCount > top[max].searchCount {
				max = j
			}
		}

		top[i], top[max] = top[max], top[i]
	}

	fmt.Println("=========================")
	fmt.Println("        STATISTIK")
	fmt.Println("=========================")
	fmt.Printf("Total Resep : %d\n", len(data))

	fmt.Println()
	fmt.Print("Kategori Terbanyak : ")
	tampilkanTeks(kategori[idxMax])
	fmt.Printf(" (%d resep)\n", jumlah[idxMax])

	fmt.Print("Kategori Tersedikit : ")
	tampilkanTeks(kategori[idxMin])
	fmt.Printf(" (%d resep)\n", jumlah[idxMin])

	fmt.Println()
	fmt.Println("-------------------------")
	fmt.Println("Jumlah Resep per Kategori")
	fmt.Println("-------------------------")

	for i := 0; i < len(kategori); i++ {
		tampilkanTeks(kategori[i])
		fmt.Printf(" : %d resep\n", jumlah[i])
	}

	fmt.Println()
	fmt.Println("-------------------------")
	fmt.Println("Top 5 Pencarian")
	fmt.Println("-------------------------")

	batas := 5
	if len(top) < 5 {
		batas = len(top)
	}

	for i := 0; i < batas; i++ {
		fmt.Printf("%d. ", i+1)
		tampilkanTeks(top[i].title)
		fmt.Printf(" (%dx)\n", top[i].searchCount)
	}

	fmt.Println("=========================")
}

func main() {
	var pilih int
	var data []Resep

	for {
		fmt.Println("==================")
		fmt.Println("      RESEPKU     ")
		fmt.Println("==================")
		fmt.Println("1. Lihat Resep")
		fmt.Println("2. Tambah Resep")
		fmt.Println("3. Ubah Resep")
		fmt.Println("4. Hapus Resep")
		fmt.Println("5. Statistik")
		fmt.Println("6. Keluar")
		fmt.Println("==================")
		fmt.Print("Pilih(1-6): ")
		fmt.Scan(&pilih)

		if pilih < 1 || pilih > 6 {
			fmt.Println("Pilihan tidak valid!")
			continue
		}

		if pilih == 6 {
			break
		}

		if pilih == 1 {
			lihat(data)
		} else if pilih == 2 {
			data = tambah(data)
		} else if pilih == 3 {
			data = ubah(data)
		}else if pilih == 4{
			data = hapus(data)
		}else if pilih == 5{
			statistik(data)
		}

	}

}
