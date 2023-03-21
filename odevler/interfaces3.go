package odevler

import "fmt"

//Bir Animal adında bir interface oluşturun ve Speak adında bir metod ekleyin. Daha sonra, Dog, Cat ve Bird adında üç farklı struct oluşturun
//ve her bir struct'a Speak metodunu uygulayın. Son olarak, bir Zoo adında bir struct oluşturun ve bu struct içindeki bir animals slice'ına
//her bir hayvanı ekleyin. Daha sonra, animals slice'ındaki her hayvanın Speak metodunu çağırarak hayvanların konuşmasını sağlayın.

type Animal interface {
	Speak()
}

//hayvanlar----------------------
type Dog struct {
	voice string
}

type Cat struct {
	voice string
}

type Bird struct {
	voice string
}

//fonksiyonlar--------------
func (d Dog) Speak() {
	fmt.Println("HAV")
}
func (c Cat) Speak() {
	fmt.Println("MİYAV")
}
func (b Bird) Speak() {
	fmt.Println("CİKCİK")
}

type Zoo struct {
	animal []Animal
}

// henüz tamamlanmadı...

func Mainicin1() {

}
