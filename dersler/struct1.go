package dersler

import "fmt"

type urun struct {
	kategori string
	uruntipi string
	marka    string
	fiyat    int
}

func Structyapisi() {
	z := urun{kategori: "Hazırgıda", uruntipi: "gofret", marka: "eti", fiyat: 10}
	x := urun{"unlumamül", "tambuğdayekmeği", "uno", 15}
	fmt.Println(z)
	fmt.Println(x)
	fmt.Println(urun{kategori: "kişiselBakım", uruntipi: "şampuan", marka: "elidor", fiyat: 41})

}
