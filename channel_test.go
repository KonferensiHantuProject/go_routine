package go_routine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	defer close(channel)
	go func() {

		time.Sleep(2 * time.Second)
		channel <- "Anjay Mabar"
		fmt.Println("Selesai Mengirim data Ke Channel")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}
