package internal

import (
	"fmt"
	"github.com/RobbingDaHood/my_little_snake/configs"
	"github.com/RobbingDaHood/my_little_snake/internal/datastructure"
	"os"
	"os/exec"
	"time"
)

func RenderLoop() {
	ticker := time.NewTicker(configs.UpdateUiDelay)
	for range ticker.C {
		clearScreen()
		render()
	}
}

func render() {
	var state = <-gameStateReads
	for i := 0; i < configs.GridHeight; i++ {
		for j := 0; j < configs.GridWidth; j++ {
			pointToRender := datastructure.Point{X: i, Y: j}
			printPointToChar(pointToRender, state)
		}
		fmt.Println()
	}
}

func printPointToChar(point datastructure.Point, state GameState) {
	if point == state.Food {
		fmt.Print("F")
	} else if contains(state.Snake, point) {
		fmt.Print("S")
	} else if contains(state.Obstacles, point) {
		fmt.Print("X")
	} else {
		fmt.Print(".")
	}
}

func clearScreen() {
	cmd := exec.Command(configs.OsClearCommand)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic("Error clearing screen")
	}
}

func contains(points []datastructure.Point, p datastructure.Point) bool {
	for _, point := range points {
		if point == p {
			return true
		}
	}
	return false
}
