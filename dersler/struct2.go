package dersler

import "fmt"

// dışardan erişilmesini istediğin şeylerin baş harfini daima büyük yap yoksa kullanamazsın!!!!!!!

type Product struct {
	Name      string
	Unitprice int
	Brand     string
}

func (p Product) Uruntanit(isim string) {
	fmt.Printf("Merhaba %v, Ürünün adi : %v, ürünün değeri: %v, ürünün markası: %v", isim, p.Name, p.Unitprice, p.Brand)

}

// yapıcı fonksiyonlar, inheritance, pointer kullanımı gibi konular da gelecek
