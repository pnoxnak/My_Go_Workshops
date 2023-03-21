package defer_statement

import "fmt"

func A() {
	fmt.Println("A fonksiyonu çalıştı")
}
func B() {
	defer C()
	fmt.Println("B fonksiyonu çalıştı")
	defer A()
	defer D()
}
func C() {
	fmt.Println("C fonksiyonu çalıştı")
}
func D() {
	fmt.Println("D fonksiyonu çalıştı")
}

// defer ifadesi last in first out mantığıyla çalışır
func Demo1() {
	B()
}
