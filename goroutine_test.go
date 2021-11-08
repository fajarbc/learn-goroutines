package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func Tambah(a int32, b int32) {
	fmt.Println(a + b)
}

func TestHelloWorld(t *testing.T) {
	go Tambah(100, 200)
	go RunHelloWorld()
	fmt.Println("Hi")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestBanyakGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
