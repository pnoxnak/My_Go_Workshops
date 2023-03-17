package dersler

import "fmt"

func Demo1() {
	// slices ile array başlatma, ilki dizinin tipini ikincisi boyutunu belirtir
	// normalde bir dizi tanımlandığında diziye eleman sayısı tanımı da yapılır ve dizideki eleman sayısı arttırılamaz
	// ancak slices kullanarak append fonksiyonuyla diziye yeni bir eleman eklenebilir ve dizi boyutu yeni eklenen eleman ile birlikte güncellenir
	isimler := make([]string, 3) //slice oluşturma
	isimler[0] = "Murat"
	isimler[1] = "Beyza"
	isimler[2] = "Esma"

	//sayilar := []int{1,2,3,4,5,6}
	// Bu şekilde dizi tanımlaması da yapabilirsin

	//isimler[3] = "Burak" // yukarda 3 olarak tanımlanan eleman sayısını aştığı için program hata verecektir
	fmt.Println(isimler)
	fmt.Println("dizinin uzunluğu: ", len(isimler))
	//ancak append fonksiyonuyla bunu atlatabiliriz
	isimler = append(isimler, "Burak") // burada isimler dizisini tamamen değiştiriyoruz ve yeni eklenen eleman ile yeni bir isimler dizisi oluşturuyoruz.
	fmt.Println(isimler)
	fmt.Println("append sonrası dizinin uzunluğu: ", len(isimler))

	// kopyalama işlemi
	isimlerkopya := make([]string, len(isimler))
	copy(isimlerkopya, isimler)
	fmt.Println("isimlerkopya: ", isimlerkopya)
	fmt.Println("isimler: ", isimler)
	isimlerkopya = append(isimlerkopya, "Gürhan")
	isimler = append(isimlerkopya, "Mehmet")
	fmt.Println("isimlerkopya: ", isimlerkopya)
	fmt.Println("isimler: ", isimler)
	fmt.Println("isimler[2:5]: ", isimler[2:5])
	fmt.Println("isimler[3:]: ", isimler[3:])
	fmt.Println("isimler[:4]: ", isimler[:4])

}
