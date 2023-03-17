package dersler

import "fmt"

func Demo2(arg1 int, arg2 int) (int, int, int, float32) {
	toplama := arg1 + arg2
	cikarma := arg1 - arg2
	carpma := arg1 * arg2
	bolme := float32(arg1) / float32(arg2)

	return toplama, cikarma, carpma, bolme

}

func Variadicfonksiyon(args ...int) int { // variadic fonksiyonlara parametre olarak dizi verebilirsin
	fmt.Println("\nBu fonksiyon kendisine verilen bütün parametrelerin toplamını döndürür")

	toplam := 0
	for i := 0; i < len(args); i++ {
		toplam = toplam + args[i]

	}
	return toplam
}
