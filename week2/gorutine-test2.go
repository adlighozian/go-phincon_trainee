package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Filter func(string) bool

func main() {
	log.Println("start")
	start := time.Now()
	duration := time.Since(start)

	funclat := latihan

	v := funclat("Aditya Ananta Putra",
		"ADNAN NUR JAILANI",
		"Afrila Zahra Prasetyo",
		"Aida Nirmala",
		"Amalia Rahma",
		"ANDINI DWI RIZKY PUTRI",
		"Angga Putra",
		"CAHAYA WIJAYA IJAM",
		"Dea Christina",
		"DEVITA NANDA OKTAVIA",
		"DEWI AYU LESTARI",
		"Dhevina Ananda Fitri",
		"DHEVY YANI RIZKI",
		"Dwi Mahani",
		"Eri Santosos",
		"Faiza Amalia Mahfudz",
		"FAUZIYATUNNISA",
		"Febriyanti Syabina",
		"Hafifah Nur Azmi Pratiwi",
		"Juliansyah Husien",
		"Kharisa Nur Aziza",
		"Lun Wina",
		"MAHREVA ROESTHA RAMADANIATY",
		"MUHAMMAD FADHILAH",
		"Muhammad Fatah Firdaus",
		"Muhammad Rafly Ahya",
		"Muhammad Rahmin Noor",
		"Muhammad Rivaldi",
		"Muhammad Rizky Rachmadhani",
		"Nadiah Nahdhah Islamiyah",
		"NOLA FEBRIANA SAPUTRI",
		"RAFADYAN REYHAN MARITZA WERAT",
		"Robby Satria",
		"Salshabilla Wahyu Anafi",
		"STEVI VIONI PAKULLA",
		"Winny Rahmah Nia")

	fmt.Println(v)

	log.Println("done in", duration.Seconds(), "seconds")
}

func latihan(param ...string) []string {

	var jumlahData int
	var jumlahAk int

	slice := make([]string, jumlahData)
	jumlaha := make(map[string]int)

	for _, nama := range param {
		jumlahData++
		if strings.Contains(strings.ToLower(nama), "vioni") {
			slice = append(slice, nama)
		}
		jumlahAk = strings.Count(nama, "a")
		jumlaha[nama] = jumlahAk
	}

	return slice
}
