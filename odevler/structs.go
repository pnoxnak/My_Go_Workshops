package odevler

//Bir kullanıcının adını, soyadını ve e-posta adresini içeren bir User struct'u tanımlayın. Ardından, kullanıcı verilerini alın ve bu struct'a atayın.

type User struct {
	name    string
	surname string
	epost   string
}

func Yenikayitekle(adi string, soyadi string, mail string) User {
	var kullanici User
	kullanici.name = adi
	kullanici.surname = soyadi
	kullanici.epost = mail

	return kullanici

}

//Bir restoran menüsü için bir MenuItem struct'u tanımlayın. Her öğe için bir ad, açıklama ve fiyat içermesi gerektiğini unutmayın. Ardından, birkaç öğe ekleyin ve menüyü yazdırın.

type MenuItem struct {
	name        string
	description string
	price       float64
}

type Menu []MenuItem

func (m Menu) AddItem(name, description string, price float64) Menu {
	return append(m, MenuItem{name, description, price})
}

//Bir Rectangle struct'u tanımlayın ve width ve height özelliklerine sahip olsun. Ardından, struct'a, dikdörtgenin alanını hesaplamak için bir area() fonksiyonu ekleyin.

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Area() int {
	return r.Height * r.Width

}

//Bir Student struct'u tanımlayın ve name, age, gender ve grades özelliklerine sahip olsun. grades özelliği, bir map olmalı ve ders adlarına göre notları içermelidir. Ardından, her öğrenci için ortalama bir not hesaplamak için bir calculateAverage() fonksiyonu ekleyin.

type Student struct {
	Name   string
	Age    int
	Gender string
	Grades map[string]float64
}

func (s Student) CalculateAverage() float64 {
	total := 0.0
	for _, grade := range s.Grades {
		total += grade
	}
	return total / float64(len(s.Grades))

}
