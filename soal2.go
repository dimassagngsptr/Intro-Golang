package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type Nilai struct {
// 	pelajaran string
// 	nilai     int
// }

// func main() {
// 	nilai := []Nilai{
// 		{pelajaran: "IPA", nilai: 80},
// 		{pelajaran: "IPS", nilai: 90},
// 		{pelajaran: "B.Indonesia", nilai: 89},
// 		{pelajaran: "B.Inggris", nilai: 69},
// 	}
// 	totalNilai := 0
// 	if len(nilai) < 4 {
// 		fmt.Println("Semua pelajaran wajib diisi")
// 		return
// 	}
// 	if len(nilai) > 4 {
// 		fmt.Println("Pelajaran tidak boleh lebih dari 4")
// 		return
// 	}

// 	for _, nilai := range nilai {
// 		fmt.Printf("Pelajaran: %s nilai: %d\n", nilai.pelajaran, nilai.nilai)
// 		if reflect.TypeOf(nilai.nilai) != reflect.TypeOf(0) {
// 			fmt.Println("tipe data harus integer")
// 			return
// 		}
// 		if nilai.nilai < 0 || nilai.nilai > 100 {
// 			fmt.Println("Masukan nilai diantara 0-100")
// 			return
// 		}
// 		totalNilai += nilai.nilai
// 	}

// 	rataRata := totalNilai / len(nilai)
// 	grade := ""
// 	if rataRata >= 90 && rataRata <= 100 {
// 		grade = "A"
// 	} else if rataRata >= 80 && rataRata <= 89 {
// 		grade = "B"
// 	} else if rataRata >= 70 && rataRata <= 79 {
// 		grade = "C"
// 	} else if rataRata >= 60 && rataRata <= 69 {
// 		grade = "D"
// 	} else if rataRata >= 0 && rataRata <= 59 {
// 		grade = "E"
// 	} else {
// 		grade = "Nilai tidak valid"
// 	}
// 	fmt.Println("Rata-rata nilai: ", rataRata)
// 	fmt.Println("Grade nilai: ", grade)

// }
