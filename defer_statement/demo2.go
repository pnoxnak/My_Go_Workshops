package defer_statement

import "fmt"

type Product struct {
	productname  string
	productprice float64
}

func (p Product) save() {
	fmt.Printf("ürün kaydedildi. adı: %v, değeri :%v \n", p.productname, p.productprice)
	defer logla()

}
func logla() {
	fmt.Println("loglandı...")
}
func Demo2() {
	urun1 := Product{productname: "pc", productprice: 11233}
	defer urun1.save() //defer fonksiyon don sonra çalıştırılır
	urun2 := Product{productname: "sdf", productprice: 11233}
	urun2.save()
}
