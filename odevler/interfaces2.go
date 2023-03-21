package odevler

import (
	"fmt"
	"os"
)

//Bir Writer adında bir interface oluşturun ve Write adında bir metod ekleyin. Daha sonra, bir ConsoleWriter ve FileWriter struct'ları oluşturun ve her ikisinde de Write metodunu uygulayın.

type Writer interface {
	write()
}

type ConsoleWriter struct {
	data string
}

type FileWriter struct {
	data string
}

func (c ConsoleWriter) write() {
	fmt.Println(c.data)
}

func (f FileWriter) write() {
	file, err := os.OpenFile("myfile1.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// dosyaya veri yazma işlemi
	_, err = file.WriteString(f.data)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func Yazdir(w Writer) {
	w.write()
}

func Mainicin() {

	konsol := ConsoleWriter{data: "çikolatalı Pasta"}
	dosya := FileWriter{data: "çikolatasız Pasta"}

	Yazdir(konsol)
	Yazdir(dosya)

}
