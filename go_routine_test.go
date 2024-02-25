package go_routine_test

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Halo Dunia")
}

func TestCreateGoRoutines(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Oops")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Nomor ", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}
