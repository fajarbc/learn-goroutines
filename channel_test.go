package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestBuatChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // ada defer agar meskipun ada crash tetap dieksekusi

	// membuat anonymous function goroutines
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Fajar" // mengisi channel
		fmt.Println("selesai mengirim ke penerima")
	}()

	fmt.Println("menunggu data dari channel")
	penerima := <-channel // menerima isi channel
	fmt.Println(penerima)
	time.Sleep(5 * time.Second)

}
