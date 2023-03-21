package errorhandling

import "fmt"

func dogrula(i interface{}) {

	sayi, ok := i.(int)
	fmt.Println(sayi, ok)

}

func Demo2() {

	dogrula("merhaba") //sayi için 0 ve ok için false değeri atanır
	dogrula(23)        //sayi için 23 ve ok için true değeri atanır

	//herşey bir interface dir
	var x interface{} = 15
	var y interface{} = "baybay"
	dogrula(x)
	dogrula(y)

}
