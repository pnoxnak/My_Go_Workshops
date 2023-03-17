package dersler

import "fmt"

func For_range_kullanimi() {
	//dizilerde/slice for range kullanımı:
	dizi := []string{"a", "b", "c", "d"}

	for i, harf := range dizi { // dizideki elemanları dolaş ve her seferinde elemanı harf değişkenine ata. i indexi belirtir, eğer kullanmayacaksan _ koyabilirsin
		fmt.Printf("%v indexindeki harf : %v \n", i, harf)

	}

}

func Maplerde_forrange() {

	urunler := make(map[string]int)
	urunler["elma"] = 11
	urunler["armut"] = 25
	urunler["kavun"] = 10
	urunler["ayva"] = 44
	urunler["lahana"] = 52
	fmt.Println(urunler)

	for urun, miktar := range urunler {
		fmt.Println("ürün: ", urun, "miktar: ", miktar)

	}

}
