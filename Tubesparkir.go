package main //hidup harus hidup

import "fmt"

const NMAX int = 99

type parkir struct {
	jenisKendaraan       string
	nomorKendaraan       string
	jam1, menit1, detik1 int // waktu masuk
	jam2, menit2, detik2 int // waktu keluar
	totalUang            int
}

type tabParkir [NMAX]parkir

type user struct {
	username, password, jabatan string
}

func menu(jabatan string) {
	fmt.Println("--------------------------")
	fmt.Println("          M E N U        ")
	fmt.Println("--------------------------")
	if jabatan == "admin" {
		fmt.Println("1. Tambah Data Petugas")
		fmt.Println("2. Edit Data Petugas")
		fmt.Println("3. Hapus Data Petugas")
	} else if jabatan == "petugas" {
		fmt.Println("1. Input Kendaraan Masuk")
		fmt.Println("2. Input Kendaraan Keluar")
		fmt.Println("3. Hitung Biaya Parkir")
		fmt.Println("4. Cetak Total Uang")
		fmt.Println("5. Cari Nomor Polisi")
		fmt.Println("6. Cetak Daftar Terurut")
	}
	fmt.Println("7. Exit")
	fmt.Println("--------------------------")
}

func login(users []user) user {
	var uname, pass string
	fmt.Print("Username: ")
	fmt.Scan(&uname)
	fmt.Print("Password: ")
	fmt.Scan(&pass)
	var i int
	for i = 0; i < len(users); i++ {
		if users[i].username == uname && users[i].password == pass {
			return users[i]
		}
	}
	return user{"", "", ""}
}

func inputKendaraanMasuk(A *tabParkir, n *int) {
	if *n < NMAX {
		fmt.Print("Masukkan jenis kendaraan (mobil/motor): ")
		fmt.Scan(&A[*n].jenisKendaraan)
		fmt.Print("Masukkan nomor kendaraan: ")
		fmt.Scan(&A[*n].nomorKendaraan)
		fmt.Print("Masukkan jam, menit, detik kendaraan masuk: ")
		fmt.Scan(&A[*n].jam1, &A[*n].menit1, &A[*n].detik1)
		*n++
	}
}

func inputKendaraanKeluar(A *tabParkir, n int) {
	var nomor string
	fmt.Print("Masukkan nomor kendaraan: ")
	fmt.Scan(&nomor)
	var i int
	for i = 0; i < n; i++ {
		if A[i].nomorKendaraan == nomor {
			fmt.Print("Masukkan jam, menit, detik kendaraan keluar: ")
			fmt.Scan(&A[i].jam2, &A[i].menit2, &A[i].detik2)
			return
		}
	}
	fmt.Println("Kendaraan tidak ditemukan.")
}

func hitungBiayaParkir(A *tabParkir, n int, total *int) {
	var nomor string
	fmt.Print("Masukkan nomor kendaraan: ")
	fmt.Scan(&nomor)
	var i int
	for i = 0; i < n; i++ {
		if A[i].nomorKendaraan == nomor {
			var masuk = A[i].jam1*3600 + A[i].menit1*60 + A[i].detik1
			var keluar = A[i].jam2*3600 + A[i].menit2*60 + A[i].detik2
			var durasi = keluar - masuk
			var jam = durasi / 3600
			if durasi%3600 > 0 {
				jam++
			}
			if A[i].jenisKendaraan == "mobil" {
				A[i].totalUang = 5000 + (jam-1)*3000
			} else if A[i].jenisKendaraan == "motor" {
				A[i].totalUang = 2000 + (jam-1)*1000
			}
			*total += A[i].totalUang
			fmt.Println("Biaya: Rp", A[i].totalUang)
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}

func cariNomorKendaraan(A tabParkir, n int) {
	var nomor string
	fmt.Print("Masukkan nomor kendaraan: ")
	fmt.Scan(&nomor)
	var i int
	for i = 0; i < n; i++ {
		if A[i].nomorKendaraan == nomor {
			fmt.Println("Kendaraan ditemukan:", A[i])
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}

func cetakTotalUang(total int) {
	fmt.Println("Total uang parkir hari ini: Rp", total)
}

func readData(A *tabParkir, n *int) {
	/*
		IS: Array A sebanyak n elemen terdefinisi sembarang
		Proses: 1. Membaca banyak elemen (n)
				2. Jika n > NMAX, n = NMAX
				3. Baca seluruh atribut A dan masukkan ke array sebanyak n elemen
		FS: Array A sebanyak n elemen berisi data mahasiswa
	*/
	if *n > NMAX {
		*n = NMAX
	}
	for i := 0; i < *n; i++ {
		fmt.Scan(&A[i].jenisKendaraan, &A[i].nomorKendaraan, &A[i].waktuMasuk, &A[i].waktuKeluar)
	}
}

func findData(A tabParkir, n int, nomor_kendaraan string) int {
	/* Mencari mahasiswa berdasarkan NIM dengan sequential search
	   Mengembalikan indeks jika ditemukan, -1 jika tidak */
	for i := 0; i < n; i++ {
		if A[i].nomorKendaraan == nomor_kendaraan {
			return i
		}
	}
	return -1
}

func editData(A *tabParkir, n int, jenisKendaraan string, nomorKendaraanBaru string) {
	/*
		IS: Array A sebanyak n elemen terdefinisi.
			nim mahasiswa yang akan diedit (nim) terdefinisi
			jurusan baru (jurusanBaru) terdefinisi
		Proses: 1. Lakukan pencarian indeks mahasiswa berdasarkan NIM
				2. Jika ditemukan, ubah jurusannya dan cetak "Data has been edited"
				3. Jika tidak ditemukan, cetak "Data has not been edited"
		FS: Jurusan mahasiswa berubah jika ditemukan
	*/
	idx := findData(*A, n, jenisKendaraan)
	if idx != -1 {
		A[idx].nomorKendaraan = nomorKendaraanBaru
		fmt.Println("Data has been edited")
	} else {
		fmt.Println("Data has not been edited")
	}
}

func Mengurutkan(A tabParkir, n int) {
	var i, j int
	var temp parkir
	for i = 1; i < n; i++ {
		temp = A[i]
		j = i - 1
		for j >= 0 && A[j].nomorKendaraan > temp.nomorKendaraan {
			A[j+1] = A[j]
			j = j - 1
		}
		A[j+1] = temp
	}
	// Menampilkan hasil yang telah diurutkan
	for i = 0; i < n; i++ {
		fmt.Println(i+1, A[i].jenisKendaraan, A[i].nomorKendaraan, "Rp", A[i].totalUang)
	}
}

func deleteData(A *tabParkir, n *int, plat string) {
	var idx int
	idx = findData(*A, *n, plat)
	if idx != -1 {
		for i := idx + 1; i < *n; i++ {
			A[i-1] = A[i]
		}
		*n--
	} else {
		fmt.Println("Data not found")
	}
}
