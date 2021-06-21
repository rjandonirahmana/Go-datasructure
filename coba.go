package main

import "fmt"

type persegi struct {
	sisi int32
}

type hitung interface {
	keliling(angka int32) int32
	luas(angka int32) int32
}

func (p persegi) keliling(angka int32) int32 {
	return p.sisi * 4 * angka
}

func (p persegi) luas() int32 {
	return p.sisi * 5
}

func (p persegi) tinggibangunan() {
	fmt.Println(p.sisi * 10 / 3)
}

func main() {
	persegi1 := persegi{5}
	fmt.Println(persegi1.keliling(4))
	fmt.Println(persegi1.luas())
	persegi1.tinggibangunan()

}
