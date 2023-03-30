package main

import (
	"fmt"
)

func main() {

	fmt.Println("merhaba dünya")

	//deneme()
	//Benim_fonksiyonum()
	//pointer_kullanimi()
	//fmt.Scanf("")
	//sabit_ve_iota()
	//diziler()
	map_kullanimi()

}

// eğer bir fonksiyon oluşturuyorsan ve fonksiyon adının ilk harfini büyük yaparsan
// bu fonksiyonu diğer packagelerden kullanabilirsin
func deneme() {
	var i float32 = 3.1212
	fmt.Print("benim değerim= ", i, "\n")
}

func Benim_fonksiyonum() {

	var f string
	f = "tabi efenim"
	fmt.Println(f)

	otomatik_atama_işareti := 42

	fmt.Println(otomatik_atama_işareti)

}

func pointer_kullanimi() {
	fmt.Println("-----------------pointer kullanimi-------------------")
	var isim *string = new(string) //yeni bir bellek alanı kullanılacak ve string tutacak
	*isim = "merhaba"

	fmt.Println("isim değişkeninin bulunduğu adres: ", isim)
	fmt.Println("isim değişkeninde bulunan değer: ", *isim)

	var benimsayim int = 42

	var yeni *int = &benimsayim
	*yeni = 43

	fmt.Println("yeni değişkeninin bulunduğu adres: ", yeni)
	fmt.Println("yeni değişkeninde bulunan değer: ", *yeni)
	fmt.Println("benimsayim değişkeninin bulunduğu adres: ", &benimsayim)
	fmt.Println("benimsayim değişkeninin bulunduğu adresteki değer: ", benimsayim)
	fmt.Println("-----------------pointer kullanimi-------------------")

}
func sabit_ve_iota() {
	const sabit_degisken int = 41
	fmt.Println("sabit değişken: ", sabit_degisken)

	const (
		t = 1
		u
		v
	)

	fmt.Println(t, u, v)

	const (
		f = 0
		h
		i = 55
		j
	)

	fmt.Println(f, h, i, j)
	// iota kullanımı bize devamı olarak ertan değerler atamamızı sağlar
	const (
		pazartesi = iota
		sali
		carsamba
		persembe
		cuma
		cumartesi
		pazar
	)
	fmt.Println("pazartesi:", pazartesi, ", salı:", sali, carsamba, persembe, cuma, cumartesi, pazar)
	fmt.Println("şeklinde devam ediyor")

	const (
		first = iota + 1
		second
		third
		fourth
	)

	fmt.Println(first, second, third, fourth)

	const (
		x = iota + 1
		y
		z = iota + 650
		a
		b
		c
	)
	fmt.Println(x, y, z, a, b, c)

	const (
		k = iota + 1
		l = 1 << iota
		m
		n
		o
		p
	)
	fmt.Println(k, l, m, n, o, p)

	const (
		aa = iota
		bb
		cc
		df = 0
		gh
		jk = iota
		li
	)
	fmt.Println(aa, bb, cc, df, gh, jk, li)

}

func diziler() {
	var dizim [3]int
	dizim[0] = 1
	fmt.Println(dizim)
	dizim[1] = 345
	dizim[2] = 567876678
	fmt.Println(dizim)

	//farklı bir array tanımlama
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	slices := arr[1:3]
	fmt.Println("slices", slices)
}

func map_kullanimi() {
	//map aslında dictionary yapısıdır
	// sözdizimi şu şekildedir: map[keyin veri tipi]anahtarın veri tipi
	ornek := map[string]string{}
	ornek["deneme"] = "12343212"
	ornek["merhaba"] = "dünya"

	fmt.Println(ornek)
	fmt.Println(ornek["merhaba"])
}

//kaldığım süre : 1.36.27 struct yapısı/go for hackers bölüm 1
//iota switch case, select case,
