package workshop2

import "fmt"

// sayı tahmin oyunu

func Tahminoyunu1() {
	var akildaki_sayi int
	var tahmin int
	fmt.Print("Oyuncu 1 aklından 1-100 arasında bir sayı tutsun:")
	fmt.Scanln(&akildaki_sayi)
	var sayac int = 0
	for true {
		sayac = sayac + 1
		fmt.Print("oyuncu 2 tahmin etmeye çalış:")
		fmt.Scanln(&tahmin)
		if tahmin == akildaki_sayi {
			fmt.Printf("KAZANDINIZ, oyunu %v tahminde kazandın \n", sayac)
			break
		}
	}
}

func Tahminoyunu2() {
	var akildaki_sayi int
	var tahmin int
	fmt.Print("Oyuncu 1 aklından 1-100 arasında bir sayı tutsun:")
	fmt.Scanln(&akildaki_sayi)
	var sayac int = 0
	for true {
		sayac = sayac + 1
		fmt.Print("oyuncu 2 tahmin etmeye çalış:")
		fmt.Scanln(&tahmin)
		if tahmin < akildaki_sayi {
			fmt.Println("yukarı")
		} else if tahmin > akildaki_sayi {
			fmt.Println("aşağı")
		} else {
			fmt.Printf("KAZANDINIZ, oyunu %v tahminde kazandın \n", sayac)
			break
		}
	}
}
