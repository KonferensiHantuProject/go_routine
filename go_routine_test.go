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
