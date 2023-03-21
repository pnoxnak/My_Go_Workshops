package dersler

import (
	"fmt"
	"time"
)

func Go_routines1() {
	//sadece çift sayıları yazdır
	for i := 0; i < 10; i += 2 {
		fmt.Println("sayi: ", i)
		time.Sleep(10 * time.Millisecond)
	}
}

func Go_routines2() {
	//sadece tek sayıları yazdır
	for i := 1; i < 10; i += 2 {
		fmt.Println("sayi: ", i)
		time.Sleep(5 * time.Millisecond)
	}

}
