package dersler

import "fmt"

func Maplar() {
	// iki farklı tanımlama şekli vardır:
	sozluk1 := make(map[string]string)
	sozluk1["Merhaba"] = "Hello"
	sozluk1["dünya"] = "world"

	sozluk2 := map[string]string{"masa": "table", "koltuk": "chair"}

	fmt.Println(sozluk1)
	fmt.Println(sozluk2)
	fmt.Println(sozluk1["dünya"])

	deger, exist := sozluk2["masa"] // sozluk2["masa"] dendiğinde bu 2 değer döndürür. ilki masanın karşılığı ikincisi ise var olup olmadığı(true/false)
	fmt.Println(deger, exist)

	elemansayisi := len(sozluk2)
	fmt.Println("elemansayisi: ", elemansayisi)
}
