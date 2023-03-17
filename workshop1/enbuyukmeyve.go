package workshop1

import (
	"fmt"
)

//senaryo:
//Ali, evdeki 3 farklı meyveyi karıştırdı ve hangi meyvenin en büyük olduğunu merak etti.
//Ali, meyveleri masanın üzerine dizerken, meyveleri sıraya koyup en büyükten en küçüğe doğru sıralamayı denedi.
//Sonunda en büyük meyveyi ve en küçük meyveyi belirleyebildi. Ali, sayıları sıralamayı öğrenmişti
//ve artık en büyük değeri bulmak için bu yöntemi kullanabilecekti.

func Enbuyukbulucu() {
	var ilk_meyve int = 12
	var ikinci_meyve int = 33
	var ucuncu_meyve int = 55

	if ilk_meyve > ikinci_meyve && ilk_meyve > ucuncu_meyve {
		fmt.Println("en buyuk ilk meyve")
	} else if ikinci_meyve > ilk_meyve && ikinci_meyve > ucuncu_meyve {
		fmt.Println("en buyuk ikinci meyve")
	} else {
		fmt.Println("en buyuk ucuncu meyve")
	}

}
