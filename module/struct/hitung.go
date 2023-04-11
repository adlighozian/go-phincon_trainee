package main

import "fmt"

// interface
type Hitung interface {
	luas() float32
	keliling() float32
}

type HitungBalok interface {
	volume() float32
}

// struct
type Segitiga struct {
	alas   float32
	tinggi float32
}

type Balok struct {
	panjang float32
	luas    float32
	tinggi  float32
}

// method
func (s Segitiga) luas() float32 {
	hasil := s.alas * s.tinggi / 2
	return hasil
}

func (s Segitiga) keliling() float32 {
	hasil := s.alas * 3
	return hasil
}

func (b Balok) volume() float32 {
	hasil := b.panjang * b.luas * b.tinggi
	return hasil
}

// function
func hitungLuas(hitung Hitung) {
	luas := hitung.luas()
	fmt.Println(luas)
}

func hitungKeliling(hitung Hitung) {
	keliling := hitung.keliling()
	fmt.Println(keliling)
}

func hitungVolume(hitung HitungBalok) {
	volume := hitung.volume()
	fmt.Println("ini balok", volume)
}

// main
func main() {
	segitiga := Segitiga{alas: 12, tinggi: 2}
	hitungKeliling(segitiga)
	balok := Balok{tinggi: 12, luas: 12, panjang: 12}
	hitungVolume(balok)

}
