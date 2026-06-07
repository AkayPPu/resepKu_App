package main
import "fmt"

type Resep struct{
	title string

	nbahan int
	bahan [50]string

	nsteps int
	steps [50]string

	kategori string
	duration int
	searchCount int
}


func tambah(data []Resep) []Resep {
	var r Resep

	fmt.Println("===================")
	fmt.Println("  TAMBAHKAN RESEP  ")
	fmt.Println("===================")

	fmt.Print("Masukkan Judul : ")
	fmt.Scan(&r.title)

	fmt.Print("Jumlah Bahan : ")
	fmt.Scan(&r.nbahan)

	for i := 0; i < r.nbahan; i++ {
		fmt.Printf("Bahan %d : ", i+1)
		fmt.Scan(&r.bahan[i])
	}

	fmt.Print("Jumlah Langkah : ")
	fmt.Scan(&r.nsteps)

	for i := 0; i < r.nsteps; i++ {
		fmt.Printf("Langkah %d : ", i+1)
		fmt.Scan(&r.steps[i])
	}

	fmt.Print("Kategori : ")
	fmt.Scan(&r.kategori)

	fmt.Print("Durasi (menit) : ")
	fmt.Scan(&r.duration)

	r.searchCount = 0

	data = append(data, r)

	fmt.Println("Resep berhasil ditambahkan!")

	return data
}



func ubah(data []Resep) []Resep{

}

func lihat(data []Resep){

}

func hapus(data []Resep) []Resep{

}

func cari(data []Resep){

}

func urut(data []Resep){

}

func statistik(data []Resep){

}



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
		fmt.Scanln(&pilih)

		if pilih < 1 || pilih > 8{ //keluar
			break
		}


		if pilih == 1{
			lihat(data)
		}else if pilih == 2{
			data = tambah(data)
		}else if pilih == 3{
			data = ubah(data)
		}else if pilih == 4{

		}else if pilih == 5{

		}else if pilih == 6 {
			
		}else if pilih == 7{

		}else if pilih == 8{
			break
		}


	}


}