package workshop5

import (
	"fmt"
	"time"
)

// bu workshop ile iki farklı işlem aynı anda başlatılır ve hangi işlem ilk biterse o bilgi verir

func LongRunningTask(ch chan string) {
	fmt.Println("Uzun işlem başladı")
	time.Sleep(5 * time.Second)
	ch <- "Uzun işlem bitti"
}

func ShortRunningTask(ch chan string) {
	fmt.Println("Kısa işlem başladı")
	time.Sleep(2 * time.Second)
	ch <- "Kısa işlem bitti"
}
