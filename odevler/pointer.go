package odevler

import "fmt"

//Bir sayıyı gösteren bir pointer oluşturun ve bu sayıyı kullanarak, pointer ile değişkenin değerini 2 ile çarpın.

func Odev1() {
	var sayi int = 12

	fmt.Println("bellekteki adres: ", &sayi)
	var sayiadresi *int = &sayi

	fmt.Println("sayi: ", *sayiadresi)

	*sayiadresi = *sayiadresi * 2
	fmt.Println("yeni sayi: ", *sayiadresi)

}

//Bir diziye işaret eden bir pointer oluşturun ve bu dizinin elemanlarını tek tek dolaşarak, her elemanın değerini 2 ile çarpın.
func Odev2() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ptr := &arr[0]

	for i := 0; i < len(arr); i++ {
		fmt.Println(*ptr)    //her bir dizi elemanını kontrol ettiğimi doğrulamak için yazdım
		*ptr = *ptr * 2      // burda her bir elemanı ikiye çarpıyorum
		if i == len(arr)-1 { //index hatası vermesin diye eğer son elemandaysak döngüyü sonlandrıyorum
			break
		}
		ptr = &arr[i+1] //sonraki dizi elemanına geçiyor
	}

	fmt.Println(arr)
}

//Bir işaretçi kullanarak, iki değişkenin değerlerini takas edin.

func Odev3() {
	var sayi1 int = 12
	var sayi2 int = 15

	ptr1 := &sayi1
	ptr2 := &sayi2

	fmt.Println("ptr1 adres:", ptr1, " sayi: ", *ptr1)
	fmt.Println("ptr2 adres:", ptr2, " sayi: ", *ptr2)
	var temp int = *ptr1
	*ptr1 = *ptr2
	*ptr2 = temp
	fmt.Println("takas işlemi sonrası")
	fmt.Println("ptr1 adres:", ptr1, " sayi: ", *ptr1)
	fmt.Println("ptr2 adres:", ptr2, " sayi: ", *ptr2)

}

//Bir diziyi tersine çeviren bir işlev oluşturun. İşlev, dizinin boyutunu ve bir işaretçi kullanarak dizinin başlangıç adresini almalıdır.

func Odev4(dizi []int) {
	dizptr := &dizi[0]
	//parametre olarak verilen dizinin tersini almam gerekiyor
	tersdizi := make([]int, 0) //slice oluşturdum

	for i := 0; i < len(dizi); i++ {
		tersdizi = append(tersdizi, dizi[len(dizi)-1-i]) // parametre dizinin tersini "tersdizi" ye atıyorum
	}

	for i := 0; i < len(dizi); i++ {
		*dizptr = tersdizi[i]
		if i == len(dizi)-1 {
			break
		}
		dizptr = &dizi[i+1]

	}

}
