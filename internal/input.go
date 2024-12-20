package internal

import (
	"fmt"
	"github.com/RobbingDaHood/my_little_snake/internal/datastructure"
)

var inputCharToDirectionMap = map[string]datastructure.Point{
	"w": {X: -1}, // Up
	"s": {X: 1},  // Down
	"a": {Y: -1}, // Left
	"d": {Y: 1},  // Right
}

func HandleInputLoop() {
	for {
		input := readInputRequired()
		newDirection, success := inputCharToDirectionMap[input]
		if success {
			changeDirectionRequestWrites <- changeDirectionRequest{newDirection}
		}
	}
}

func readInputRequired() string {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		println("Error reading input: " + err.Error())
	}
	return input
}
