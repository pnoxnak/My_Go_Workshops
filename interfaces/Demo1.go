package interfaces

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.height * r.width

}

func (c circle) area() float64 {
	return c.radius * math.Pi * c.radius
}

func school(s shape) float64 {
	return s.area()
}

func Mainde_kullanilacak() {
	var dikdortgen rectangle
	dikdortgen.height = 12.5
	dikdortgen.width = 23.85
	//rec:=rectangle{width: 12,height: 32}//şeklinde de tanımlama yapılabilir
	fmt.Println("dikdortgenin alanı: ", school(dikdortgen))

}
