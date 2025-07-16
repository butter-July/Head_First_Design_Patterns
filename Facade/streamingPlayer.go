package main

import "fmt"

type StreamingPlayer struct {
}

func (player *StreamingPlayer) on() {
	fmt.Println("Streaming player on !")
}

func (player *StreamingPlayer) play(movie string) {
	fmt.Printf("Streaming Player playing%s", movie)
}
func (player *StreamingPlayer) stop(movie string) {
	fmt.Printf("Steaming Player stopped %s", movie)
}
func (player *StreamingPlayer) off() {
	fmt.Println("Steaming Player off")
}
