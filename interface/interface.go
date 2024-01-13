package main

import "fmt"

type BangunDatar interface {
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Lebar * persegiPanjang.Panjang
}

func main() {
	// bangunDatar := Persegi{Sisi: 4}
	persegiPanjang := PersegiPanjang{Panjang: 10, Lebar: 11}
	luas := SeberapaLuas(persegiPanjang)
	fmt.Println("Luas persegi panjang: ", luas)

	persegi := Persegi{Sisi: 5}
	luas2 := SeberapaLuas(persegi)
	fmt.Println("Luas persegi: ", luas2)
}

func SeberapaLuas(bangunDatar BangunDatar)int{
	return bangunDatar.HitungLuas()
}