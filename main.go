package main

import (
	"fmt"
	"strconv"
)

type city struct {
	Name string
}

type Manusia struct {
	nama string
	umur int
	kota city
}

func (s Manusia) getNamaUmur() string {
	conv := strconv.Itoa(s.umur)
	return s.nama + "-" + conv
}

func (s *Manusia) change(newlast string) {
	s.nama = newlast
}

func main() {

	manusia := Manusia{
		nama: "John",
		umur: 30,
		kota: city{Name: "New York"},
	}

	// ashel := manusia{
	// 	nama: "ashel",
	// 	umur: 12,
	// }
	fmt.Println(manusia.nama)
	manusia.change("taylor")
	fmt.Println(manusia.nama)
	// 	fmt.Println(manusia.getNamaUmur())
}
