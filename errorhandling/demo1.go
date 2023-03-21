package errorhandling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("deneme.txt")

	if err == nil { //eğer hata değeri nil ise yani yoksa dosyanın adını yazdır
		fmt.Println(f.Name())
	} else {
		if err, ok := err.(*os.PathError); ok { // noktalı virgülden öncesi atama işlemi yapar: ok değeri true ise bloğa girer. burda yapılan işlemin amacı hata türüne göre tepki vermektir
			fmt.Printf("Dosya yyoğğ : %v \n", err.Path)
			return
		} else {
			return
		}

	}

}
