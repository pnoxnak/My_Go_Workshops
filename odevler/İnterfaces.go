package odevler

import "fmt"

type krediodeme interface {
	krediode() float64
}

type ihtiyackredisi struct {
	kredi_miktari float64
	faiz_orani    float64
}

type konutkredisi struct {
	kredi_miktari float64
	faiz_orani    float64
}

type evlilikkredisi struct {
	kredi_miktari float64
	faiz_orani    float64
}

func (i ihtiyackredisi) krediode() float64 {
	return i.kredi_miktari + (i.kredi_miktari * i.faiz_orani)

}
func (k konutkredisi) krediode() float64 {
	return k.kredi_miktari + (k.kredi_miktari * k.faiz_orani)

}
func (e evlilikkredisi) krediode() float64 {
	return e.kredi_miktari + (e.kredi_miktari * e.faiz_orani)

}

func kredi_Hesapla(k krediodeme) {
	fmt.Println("ödenecek krediniz: ", k.krediode())

}

func For_main() {
	kredim := ihtiyackredisi{kredi_miktari: 150000, faiz_orani: 0.08}
	kredi_Hesapla(kredim)

}
