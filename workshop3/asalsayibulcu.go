package workshop3

import "fmt"

func isprime(arg int64) bool {
	var sonuc bool = true
	if arg == 1 || arg == 2 || arg == 3 {
		sonuc = true
	} else {
		var i int64 = 2
		for i < arg {
			if arg%i == 0 {
				sonuc = false
			}
			i = i + 1
		}
	}
	return sonuc
}

func Asalsayiyibul() {

	var sayi int64
	fmt.Println("Lütfen bir sayı giriniz: ")
	fmt.Scanln(&sayi)
	if isprime(sayi) {
		fmt.Println(sayi, "asal sayıdır")
	} else {
		fmt.Println(sayi, "asal sayı değildir")
	}

}
