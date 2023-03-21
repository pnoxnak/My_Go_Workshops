package dersler

func Ciftsayilar(ciftsayicn chan int) {
	var toplam int = 0
	for i := 0; i < 10; i += 2 {
		toplam += i
	}
	ciftsayicn <- toplam
}

func Teksayilar(teksayicn chan int) {
	var toplam int = 0
	for i := 1; i < 10; i += 2 {
		toplam += i
	}
	teksayicn <- toplam
}
