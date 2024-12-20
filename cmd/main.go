package main

import (
	"fmt"
	"github.com/RobbingDaHood/my_little_snake/internal"
)

func main() {
	go internal.UpdateGameStateLoop()
	go internal.RenderLoop()
	go internal.HandleInputLoop()
	<-internal.GameOverChannel
	fmt.Printf("Game Over!")
}
