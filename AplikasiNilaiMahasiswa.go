package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NMAX int = 1000

type dataMahasiswa struct {
	nama, nim string
	matkul    [1024]MK
	ipk       float64
	totsks    int
	nMK       int
}

type MK struct {
	namaMK string
	sks    int
	nilai  dataNilai
	ip     float64
}

type dataNilai struct {
	nQuiz, nUTS, nUAS, rata float64
	indexNilai              string
}

type tabM [NMAX]dataMahasiswa

func main() {
	// Kamus
	var mahasiswa tabM
	var input string
	var no, nM, idx int
	var hapus bool
	var out bool = false
	// Algoritma
	for !out {
		menu()
		fmt.Scan(&no)
		switch no {
		case 1:
			addData(&mahasiswa, &nM)
			sortData(&mahasiswa, nM, 1, "ascd")
		case 2:
			fmt.Println()
			fmt.Println("================= Edit Data =======================")
			fmt.Print("Ketik Nama/NIM mahasiswa yang ingin diedit: ")
			input = baca()
			input = baca()
			editData(&mahasiswa, nM, input)
		case 3:
			hapus = true
			for hapus {
				fmt.Println()
				fmt.Println("================= Hapus Data ======================")
				fmt.Print("Ketik Nama/NIM mahasiswa yang ingin dihapus: ")
				input = baca()
				input = baca()
				if input[0] >= '1' && input[0] <= '9' {
					idx = findNIM(mahasiswa, nM, input)
				} else {
					idx = findNama(mahasiswa, nM, input)
				}
				if idx == -1 {
					fmt.Println("Maaf, data tidak ditemukan.")
				} else {
					deleteData(&mahasiswa, &nM, idx, &hapus)
				}
			}
		case 4:
			cetakData(mahasiswa, nM)
		case 5:
			out = exit()
		}
	}
}

// Fungsi membaca string inputan
func baca() string {
	reader := bufio.NewReader(os.Stdin)
	nama, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	nama = strings.TrimSpace(nama)
	return nama
}

// Fungsi menambahkan data mahasiswa
func addData(ms *tabM, n *int) {
	// Kamus
	var cek string = "N"
	// Algoritma
	fmt.Println("=============== Tambah Data Mahasiswa ===============")
	fmt.Print("Nama mahasiswa: ")
	ms[*n].nama = baca()
	ms[*n].nama = baca()
	fmt.Print("NIM mahasiswa: ")
	fmt.Scan(&ms[*n].nim)
	addMK(&ms[*n], ms[*n].nMK)
	cek = cekData(ms[*n])
	if cek == "Y" {
		*n++
		fmt.Println("Data tersimpan")
	} else {
		editData(&*ms, *n, ms[*n].nim)
	}
}

// Procedure menambahkan data Mata kuliah yang diambil mahasiswa
func addMK(ms *dataMahasiswa, i int) {
	// Kamus
	var no, idx int
	var edit string
	// Algoritma
	daftarMK()
	fmt.Print("Pilih: ")
	fmt.Scan(&no)
	for no != 5 {
		edit = "Y"
		switch no {
		case 1:
			idx = findMatkul(*ms, "Algoritma Pemrograman")
			if idx != -1 {
				fmt.Println("Data sudah terisi, apakah anda ingin mengubah data? Y/N")
				fmt.Scan(&edit)
				if edit == "Y" {
					ms.totsks -= 4
					ms.nMK--
				}
			}
			if edit == "Y" {
				ms.matkul[i].namaMK = "Algoritma Pemrograman"
				ms.matkul[i].sks = 4
				editNilai(&*ms, i)
				ms.nMK++
			}
		case 2:
			idx = findMatkul(*ms, "Kalkulus Lanjut")
			if idx != -1 {
				fmt.Println("Data sudah terisi, apakah anda ingin mengubah data? Y/N")
				fmt.Scan(&edit)
				if edit == "Y" {
					i = idx
					ms.totsks -= 3
					ms.nMK--
				}
			}
			if edit == "Y" {
				ms.matkul[i].namaMK = "Kalkulus Lanjut"
				ms.matkul[i].sks = 3
				editNilai(&*ms, i)
				ms.nMK++
			}
		case 3:
			idx = findMatkul(*ms, "Matematika Diskrit")
			if idx != -1 {
				fmt.Println("Data sudah terisi, apakah anda ingin mengubah data? Y/N")
				fmt.Scan(&edit)
				if edit == "Y" {
					ms.totsks -= 3
					ms.nMK--
				}
			}
			if edit == "Y" {
				ms.matkul[i].namaMK = "Matematika Diskrit"
				ms.matkul[i].sks = 3
				editNilai(&*ms, i)
				ms.nMK++
			}
		case 4:
			idx = findMatkul(*ms, "Sistem Digital")
			if idx != -1 {
				fmt.Println("Data sudah terisi, apakah anda ingin mengubah data? Y/N")
				fmt.Scan(&edit)
				if edit == "Y" {
					i = idx
					ms.totsks -= 3
					ms.nMK--
				}
			}
			if edit == "Y" {
				ms.matkul[i].namaMK = "Sistem Digital"
				ms.matkul[i].sks = 3
				editNilai(&*ms, i)
				ms.nMK++
			}
		}
		ms.totsks += ms.matkul[i].sks
		i = ms.nMK
		daftarMK()
		fmt.Println("Pilih: ")
		fmt.Scan(&no)
	}
	calculateIPK(&*ms)
}

// Procedure mengedit data mahasiswa yang dicari
func editData(ms *tabM, n int, x string) {
	// Kamus
	var idx, no int
	// Algoritma
	daftarEdit()
	fmt.Scan(&no)
	for no != 3 {
		if x[0] >= '1' && x[0] <= '9' {
			idx = findNIM(*ms, n, x)
		} else {
			idx = findNama(*ms, n, x)
		}
		if idx == -1 {
			fmt.Println("Maaf data tidak ditemukan")
		} else {
			switch no {
			case 1:
				addMK(&ms[idx], ms[idx].nMK)
			case 2:
				editMK(&ms[idx])
			}
			fmt.Println(ms[idx].nMK)
			cekData(ms[idx])
		}
		daftarEdit()
		fmt.Scan(&no)
	}
}

// Procedure mengedit data mata kuliah yang diambil mahasiswa
func editMK(ms *dataMahasiswa) {
	// Kamus
	var edit, x string
	var i int
	// Algoritma
	edit = "Y"
	for edit != "N" {
		for i = 0; i < ms.nMK; i++ {
			fmt.Printf("%d. %s \n", i+1, ms.matkul[i].namaMK)
		}
		fmt.Println()
		fmt.Println()
		fmt.Println("========================================================")
		fmt.Println("Nama: ", ms.nama)
		fmt.Println("MIN: ", ms.nim)
		for i = 0; i < ms.nMK; i++ {
			fmt.Printf("%d. %s\n", i+1, ms.matkul[i].namaMK)
		}
		fmt.Println("========================================================")
		fmt.Println("Ketik nama matakuliah yang ingin di edit")
		fmt.Println("Ketikan (keluar) jika selesai")
		fmt.Println("========================================================")
		x = baca()
		x = baca()
		if x == "keluar" {
			edit = "N"
		} else {
			i = findMatkul(*ms, x)
			if i == -1 {
				fmt.Println("Maaf mata kuliah tidak ditemukan")
			} else {
				editNilai(&*ms, i)
			}
		}
	}
	calculateIPK(&*ms)
}

// Procedur mengedit nilai dari mata kuliah yang diambil mahasiswa
func editNilai(ms *dataMahasiswa, i int) {
	//Kamus
	var aksi string
	var nilai float64
	//Algoritma
	fmt.Println()
	fmt.Println()
	fmt.Println("========================================================")
	fmt.Printf("Silahkan pilih (Quiz/UTS/UAS) yang ingin dimasukan nilai.\nKetik (kembali) untuk kembali ke pemilihan mata kuliah.\n")
	fmt.Println("========================================================")
	fmt.Scan(&aksi)
	for aksi != "kembali" {
		if aksi == "Quiz" {
			fmt.Print("Nilai Quiz: ")
			fmt.Scan(&nilai)
			if nilai < 0 && nilai > 100 {
				fmt.Println("Maaf input tidak valid")
			} else {
				ms.matkul[i].nilai.nQuiz = nilai
			}
		} else if aksi == "UTS" {
			fmt.Print("Nilai UTS: ")
			fmt.Scan(&nilai)
			if nilai < 0 && nilai > 100 {
				fmt.Println("Maaf input tidak valid")
			} else {
				ms.matkul[i].nilai.nUTS = nilai
			}
		} else if aksi == "UAS" {
			fmt.Print("Nilai UAS: ")
			fmt.Scan(&nilai)
			if nilai < 0 && nilai > 100 {
				fmt.Println("Maaf input tidak valid")
			} else {
				ms.matkul[i].nilai.nUAS = nilai
			}
		} else {
			fmt.Println("Maaf inputan tidak sesuai, mohon pilih ulang")
		}
		fmt.Println()
		fmt.Println()
		fmt.Println("========================================================")
		fmt.Println("Mata Kuliah: ", ms.matkul[i].namaMK)
		fmt.Println("     Nilai Quiz: ", ms.matkul[i].nilai.nQuiz)
		fmt.Println("     Nilai UTS:  ", ms.matkul[i].nilai.nUTS)
		fmt.Println("     Nilai UAS:  ", ms.matkul[i].nilai.nUAS)
		fmt.Println("========================================================")
		fmt.Printf("Silahkan pilih (Quiz/UTS/UAS) yang ingin dimasukan nilai.\nKetik (kembali) untuk kembali ke pemilihan mata kuliah.\n")
		fmt.Println("========================================================")
		fmt.Scan(&aksi)
	}
	calculateAverageGrade(&ms.matkul[i])
	calculateIPK(&*ms)
}

// Menampilkan data yang dimasukan dan mengexek apakah sudah sesuai dengan apa yang diinginkan pengguna
func cekData(ms dataMahasiswa) string {
	//Kamus
	var cek string
	//Algoritma
	fmt.Println()
	fmt.Println("==================================================")
	fmt.Printf("Nama: %s\n", ms.nama)
	fmt.Printf("NIM: %s\n", ms.nim)
	for i := 0; i < ms.nMK; i++ {
		fmt.Printf("%d. Mata Kuliah: %s\n", i+1, ms.matkul[i].namaMK)
		fmt.Printf("    Nilai Quiz: %v\n", ms.matkul[i].nilai.nQuiz)
		fmt.Printf("    Nilai UTS: %v\n", ms.matkul[i].nilai.nUTS)
		fmt.Printf("    Nilai UAS: %v\n", ms.matkul[i].nilai.nUAS)

	}
	fmt.Printf("Total sks: %v \n", ms.totsks)
	fmt.Println("==================================================")
	fmt.Println("Apakah data sudah benar?(Y/N) ")
	fmt.Scan(&cek)
	return cek
}

// Prosedur untuk mencari data mahasiswa berdasarkan masukan NIM
func findNIM(ms tabM, n int, x string) int {
	// Mengembalikan index dari array data sesuai NIM yang dimasukan dengan BINARY SEARCH
	// Kamus
	var l, r, m, found int
	// Algoritma
	found = -1
	l = 0
	r = n
	for l <= r && found == -1 {
		m = (l + r) / 2
		if ms[m].nim == x {
			found = m
		} else if ms[m].nim < x {
			l = m + 1
		} else if ms[m].nim > x {
			r = m - 1
		}
	}
	return found
}

// Fungsi mencari data mahasiswa berdasarkan masukan Nama
func findNama(ms tabM, n int, x string) int {
	// Mengembalikan index dari Nama Mahasiswa yang diinginkan menggunakan SEQUENTIAL SEARCH
	//Kamus
	var i int
	//Algoritma
	found := -1
	for i < n && found == -1 {
		if ms[i].nama == x {
			found = i
		}
		i++
	}
	return found
}

// Fungsi mencari data mahasiswa berdasarkan nama Mata Kuliah
func findMatkul(ms dataMahasiswa, x string) int {
	// Mengembalikan index dari mata kuliah yang ingin diinginkan
	//Kamus
	var i int
	var m int = ms.nMK
	//Algoritma
	found := -1
	for i < m && found == -1 {
		if ms.matkul[i].namaMK == x {
			found = i
		}
		i++
	}
	return found
}

// Prosedur untuk menghapus data mahasiswa
func deleteData(ms *tabM, n *int, idx int, hapus *bool) {
	//Kamus
	type tempMK [1024]MK
	type tempMS [1024]dataMahasiswa
	var A tempMS
	var B tempMK
	var no, idMk int
	var aksi, x string
	//Algoritma
	daftarHapus()
	fmt.Scan(&no)
	switch no {
	case 1:
		fmt.Println()
		fmt.Println("========================================================")
		fmt.Println("Nama: ", ms[idx].nama)
		fmt.Println("NIM:  ", ms[idx].nim)
		fmt.Println("========================================================")
		for i := 0; i < ms[idx].nMK; i++ {
			fmt.Printf("%d. %s \n", i+1, ms[idx].matkul[i].namaMK)
		}
		fmt.Println("========================================================")
		fmt.Println("Apakah anda yakin menghapus data ini?(Y/N)")
		fmt.Scan(&aksi)
		if aksi == "Y" {
			if idx == *n-1 {
				ms[idx], ms[idx].matkul = A[0], A[0].matkul
			} else {
				for i := idx; i < *n-1; i++ {
					ms[i], ms[i].matkul = ms[i+1], ms[i+1].matkul
				}
			}
			*n--
			fmt.Println("Data berhasil dihapus.")
			*hapus = false
		}
	case 2:
		for *hapus {
			fmt.Println()
			fmt.Println()
			fmt.Println("========================================================")
			fmt.Println("Nama: ", ms[idx].nama)
			fmt.Println("NIM:  ", ms[idx].nim)
			fmt.Println("========================================================")
			for i := 0; i < ms[idx].nMK; i++ {
				fmt.Printf("%d. %s \n", i+1, ms[idx].matkul[i].namaMK)
			}
			fmt.Println("========================================================")
			fmt.Println("Masukan Nama Mata kuliah: ")
			x = baca()
			x = baca()
			idMk = findMatkul(ms[idx], x)
			if idMk == -1 {
				fmt.Println("Maaf data inputan salah")
			} else {
				fmt.Println()
				fmt.Println("========================================================")
				fmt.Println("Mata Kuliah: ", ms[idx].matkul[idMk].namaMK)
				fmt.Println("     Nilai Quiz: ", ms[idx].matkul[idMk].nilai.nQuiz)
				fmt.Println("     Nilai UTS:  ", ms[idx].matkul[idMk].nilai.nUTS)
				fmt.Println("     Nilai UAS:  ", ms[idx].matkul[idMk].nilai.nUAS)
				fmt.Println("========================================================")
				fmt.Print("Apakah anda yakin menghapus data ini?(Y/N)")
				fmt.Scan(&aksi)
				if aksi == "Y" {
					ms[idx].totsks -= ms[idx].matkul[idMk].sks
					if idMk == ms[idx].nMK-1 {
						ms[idx].matkul[idMk] = B[0]
					} else {
						for i := idMk; i < ms[idx].nMK-1; i++ {
							ms[idx].matkul[i] = ms[idx].matkul[i+1]
						}
					}

					ms[idx].nMK--
					fmt.Println("Data berhasil dihapus.")
					*hapus = false
				}
			}
		}
	}
}

// Prosedur mencetak data mahasiswa
func cetakData(ms tabM, n int) {
	//Kamus
	var t, input string
	var nPilih, no, idx, nMS int
	var msMK tabM
	//Algoritma
	menuCetak()
	fmt.Scan(&no)
	for no != 6 {
		fmt.Println()
		fmt.Println()
		switch no {
		// cetak mahasiswa, NIM terurut
		case 1:
			fmt.Scan(&t)
			sortData(&ms, n, no, t)
		case 2:
			// cetak mahasiswa, Total Sks terurut
			fmt.Scan(&t)
			sortData(&ms, n, no, t)
		case 3:
			// cetak mahasiswa yang mengikuti mata kuliah yang di input
			fmt.Scan(&t)
			daftarMK()
			fmt.Println("Masukan nama Mata Kuliah: ")
			input = baca()
			input = baca()
			idx = -1
			for i := 0; i < n && idx == -1; i++ {
				idx = findMatkul(ms[i], input)
			}

			if idx == -1 {
				fmt.Println()
				fmt.Println("Maaf data tidak ditemukan")
				fmt.Println()
			} else {
				fmt.Println("==================================================")
				fmt.Println("1. Nilai Quiz")
				fmt.Println("2. Nilai UTS")
				fmt.Println("3. Nilai UAS")
				fmt.Println("4. Nilai Rata-rata")
				fmt.Println("==================================================")
				fmt.Print("Pilih: ")
				fmt.Scan(&nPilih)
				for i := 0; i < n; i++ {
					idx = findMatkul(ms[i], input)
					if idx != -1 {
						msMK[nMS] = ms[i]
						nMS++
					}
				}
				sortDataMatkul(&msMK, nMS, input, t, nPilih)
				fmt.Printf("Data Nilai Matakuliah %s: \n", input)
				fmt.Printf("%15v %15v %10v %6v %6v\n", "NIM", "Nama", "Quiz", "UTS", "UAS")
				for i := 0; i < n; i++ {
					idx = findMatkul(msMK[i], input)
					if idx != -1 {
						fmt.Printf("%15v %15v %10v %6v %6v\n", msMK[i].nim, msMK[i].nama, msMK[i].matkul[idx].nilai.nQuiz, msMK[i].matkul[idx].nilai.nUTS, msMK[i].matkul[idx].nilai.nUAS)
					}
				}
			}
		case 4:
			// cetak mahasiswa, IPK terurut
			fmt.Scan(&t)
			sortData(&ms, n, no, t)
		case 5:
			fmt.Print("Ketik nama/NIM mahasiswa yang ingn dicari: ")
			input = baca()
			input = baca()
			if input[0] >= '1' && input[0] <= '9' {
				idx = findNIM(ms, n, input)
			} else {
				idx = findNama(ms, n, input)
			}
			if idx == -1 {
				fmt.Println("Maaf data tidak ditemukan")
			} else {
				fmt.Println("==================================================")
				fmt.Printf("Nama: %s\n", ms[idx].nama)
				fmt.Printf("NIM: %s\n", ms[idx].nim)
				fmt.Printf("IPK: %.2v\n", ms[idx].ipk)
				fmt.Println("=============== Daftar Mata Kuliah ===============")
				for i := 0; i < ms[idx].nMK; i++ {
					fmt.Printf("%d. Mata Kuliah: %s\n", i+1, ms[idx].matkul[i].namaMK)
					fmt.Printf("    Nilai Quiz: %v\n", ms[idx].matkul[i].nilai.nQuiz)
					fmt.Printf("    Nilai UTS: %v\n", ms[idx].matkul[i].nilai.nUTS)
					fmt.Printf("    Nilai UAS: %v\n", ms[idx].matkul[i].nilai.nUAS)
					fmt.Printf("Index Nilai Mata Kuliah: %s\n", ms[idx].matkul[i].nilai.indexNilai)
					fmt.Println("==================================================")
				}
				fmt.Printf("Total sks: %v \n", ms[idx].totsks)
			}
		}
		fmt.Println()
		fmt.Println()
		menuCetak()
		fmt.Scan(&no)

	}
}

// Procedure mengurutkan data mahasiswa utama
func sortData(ms *tabM, n int, urut int, t string) {
	// Mengurutkan data mahasiswa (ascending/descending) berdasarkan NIM/total SKS/IPK menggunakan INSERTION SORT
	var i, j int
	var temp dataMahasiswa
	fmt.Println()
	fmt.Println()
	if t == "ascd" {
		i = 1
		for i < n {
			temp = ms[i]
			j = i
			for (j > 0 && ms[j-1].nim > temp.nim && urut == 1) || (j > 0 && ms[j-1].totsks > temp.totsks && urut == 2) || (j > 0 && ms[j-1].ipk > temp.ipk && urut == 4) {
				ms[j] = ms[j-1]
				j--
			}
			ms[j] = temp
			i++
		}
	} else if t == "desc" {
		i = 1
		for i < n {
			temp = ms[i]
			j = i
			for (j > 0 && ms[j-1].nim < temp.nim && urut == 1) || (j > 0 && ms[j-1].totsks < temp.totsks && urut == 2) || (j > 0 && ms[j-1].ipk < temp.ipk && urut == 4) {
				ms[j] = ms[j-1]
				j--
			}
			ms[j] = temp
			i++
		}
	}
	fmt.Println("==================================================")
	fmt.Printf("%15v %15v %10v %6v\n", "NIM", "Nama", "Total SKS", "IPK")
	for i := 0; i < n; i++ {
		fmt.Printf("%15v %15v %10v %6.2v\n", ms[i].nim, ms[i].nama, ms[i].totsks, ms[i].ipk)
	}
	fmt.Println("==================================================")
	fmt.Println()
	fmt.Println()
}

// Procedure mengurutkan data mahasiswa yang mengikuti mata kuliah tertentu
func sortDataMatkul(msMK *tabM, n int, x string, t string, urut int) {
	// Mengurutkan data mahasiswa (ascending/descending) yang mengikuti mata kuliah tertentu menggunakan SELECTION SORT
	//Kamus
	var j, min, max, pass, iMin, iMax, ij int
	var temp dataMahasiswa
	//Algoritma
	if t == "ascd" {
		pass = 1
		for pass < n {
			min = pass - 1
			j = pass
			for j < n {
				iMin = findMatkul(msMK[min], x)
				ij = findMatkul(msMK[j], x)
				if iMax != -1 && ij != -1 {
					if urut == 1 {
						if msMK[min].matkul[iMin].nilai.nQuiz > msMK[j].matkul[ij].nilai.nQuiz {
							min = j
						}
					} else if urut == 2 {
						if msMK[min].matkul[iMin].nilai.nUTS > msMK[j].matkul[ij].nilai.nUTS {
							min = j
						}
					} else if urut == 3 {
						if msMK[min].matkul[iMin].nilai.nUAS > msMK[j].matkul[ij].nilai.nUAS {
							min = j
						}
					}
					j++
				}
				temp, temp.matkul = msMK[pass-1], msMK[pass-1].matkul
				msMK[pass-1], msMK[pass-1].matkul = msMK[min], msMK[min].matkul
				msMK[min], msMK[min].matkul = temp, temp.matkul
				pass++
			}
		}
	} else if t == "decs" {
		pass = 1
		for pass < n {
			max = pass - 1
			j = pass
			for j < n {
				iMax = findMatkul(msMK[max], x)
				ij = findMatkul(msMK[j], x)
				if iMax != -1 && ij != -1 {
					if urut == 1 {
						if msMK[max].matkul[iMax].nilai.nQuiz < msMK[j].matkul[ij].nilai.nQuiz {
							max = j
						}
					} else if urut == 2 {
						if msMK[min].matkul[iMax].nilai.nUTS < msMK[j].matkul[ij].nilai.nUTS {
							max = j
						}
					} else if urut == 3 {
						if msMK[min].matkul[iMax].nilai.nUAS < msMK[j].matkul[ij].nilai.nUAS {
							max = j
						}
					}
					j++
				}
				temp, temp.matkul = msMK[pass-1], msMK[pass-1].matkul
				msMK[pass-1], msMK[pass-1].matkul = msMK[max], msMK[max].matkul
				msMK[max], msMK[max].matkul = temp, temp.matkul
				pass++
			}
		}
	}
}

// function konfirmasi keluar aplikasi
func exit() bool {
	// Program akan selesai jika diberi masukan string "Y"
	// Kamus
	var x string
	var out bool = false
	// Algoritma
	for x != "Y" && x != "N" {
		fmt.Println()
		fmt.Println("========================================================")
		fmt.Println("Apakah anda yakin untuk keluar?(Y/N) ")
		fmt.Scan(&x)
		if x == "Y" {
			out = true
		} else if x == "N" {
			out = false
		} else {
			fmt.Println("Input tidak valid")
		}
	}
	return out
}

// Prosedur untuk menghitung rata-rata nilai mata kuliah
func calculateAverageGrade(matkul *MK) {
	total := matkul.nilai.nQuiz + matkul.nilai.nUTS + matkul.nilai.nUAS
	matkul.nilai.rata = total / 3.0
	// Menentukan indeks nilai berdasarkan rata-rata
	switch {
	case matkul.nilai.rata >= 85:
		matkul.nilai.indexNilai = "A"
	case matkul.nilai.rata >= 70:
		matkul.nilai.indexNilai = "B"
	case matkul.nilai.rata >= 55:
		matkul.nilai.indexNilai = "C"
	case matkul.nilai.rata >= 40:
		matkul.nilai.indexNilai = "D"
	default:
		matkul.nilai.indexNilai = "E"
	}
}

// Prosedur untuk menghitung IPK mahasiswa
func calculateIPK(ms *dataMahasiswa) {
	// Kamus
	var totalNilai float64
	var totalSKS int
	// Algoritma
	for i := 0; i < ms.nMK; i++ {
		calculateAverageGrade(&ms.matkul[i])
		nilaiAngka := 0.0
		switch ms.matkul[i].nilai.indexNilai {
		case "A":
			nilaiAngka = 4.0
		case "B":
			nilaiAngka = 3.0
		case "C":
			nilaiAngka = 2.0
		case "D":
			nilaiAngka = 1.0
		case "E":
			nilaiAngka = 0.0
		}
		ms.matkul[i].ip = nilaiAngka * float64(ms.matkul[i].sks)
		totalNilai += ms.matkul[i].ip
		totalSKS += ms.matkul[i].sks
	}

	if totalSKS > 0 {
		ms.ipk = totalNilai / float64(totalSKS)
	} else {
		ms.ipk = 0
	}
}

// Menampilkan daftar pilihan hapus
func daftarHapus() {
	fmt.Println("================= Hapus Data =====================")
	fmt.Println()
	fmt.Println("1. Hapus data Mahasiswa")
	fmt.Println("2. Hapus data Mata Kuliah")
	fmt.Println()
	fmt.Println("==================================================")
}

// Menampilkan daftar pilihan mengedit data
func daftarEdit() {
	fmt.Println("================== Edit Data =====================")
	fmt.Println()
	fmt.Println("1. Tambah Mata Kuliah")
	fmt.Println("2. Edit Nilai Mata Kuliah")
	fmt.Println("3. Selesai")
	fmt.Println()
	fmt.Println("==================================================")
}

// Menampilkan daftar mata kuliah
func daftarMK() {
	fmt.Println("=============== Daftar Mata Kuliah ===============")
	fmt.Println()
	fmt.Println("1. Algoritma Pemrograman")
	fmt.Println("2. Kalkulus Lanjut")
	fmt.Println("3. Matematika Diskrit")
	fmt.Println("4. Sistem Digital")
	fmt.Println("5. Selesai")
	fmt.Println()
	fmt.Println("==================================================")
}

// Menampilkan pilihan mencetak
func menuCetak() {
	fmt.Println("=============== Daftar Cetak ===============")
	fmt.Println()
	fmt.Println("1. List Mahasiswa berdasarkan NIM (ascd/desc)")
	fmt.Println("2. List Mahasiswa berdasarkan Total SKS (ascd/desc)")
	fmt.Println("3. List Mahasiswa berdasarkan Nilai Mata Kuliah (ascd/desc)")
	fmt.Println("4. List Mahasiswa berdasarkan IPK (ascd/desc)")
	fmt.Println("5. Mencari data Mahasiswa")
	fmt.Println("6. Selesai")
	fmt.Println()
	fmt.Println("Tuliskan Nomer Beserta (ascd/desc)")
	fmt.Println()
	fmt.Println("==================================================")
	fmt.Print("Pilih: ")
}

// Menampilkan menu utama dari aplikasi
func menu() {
	fmt.Println("=============== Nilai Mahasiswa ===============")
	fmt.Println()
	fmt.Printf("Selamat datang di aplikasi nilai Mahasiswa\nSilahkan memilih nomor terkait dengan aksi yang anda ingin lakukan \n")
	fmt.Println()
	fmt.Println("1. Menambahkan Data Mahasiswa")
	fmt.Println("2. Edit Data Mahasiswa")
	fmt.Println("3. Hapus Data Mahasiswa")
	fmt.Println("4. Cetak Data")
	fmt.Println("5. Exit")
	fmt.Println()
	fmt.Println("===============================================")
	fmt.Print("Pilih no: ")
}
