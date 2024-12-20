package internal

import (
	"github.com/RobbingDaHood/my_little_snake/configs"
	"github.com/RobbingDaHood/my_little_snake/internal/datastructure"
	"time"
)

type GameState struct {
	Snake     []datastructure.Point
	Direction datastructure.Point
	Food      datastructure.Point
	Obstacles []datastructure.Point
}

type changeDirectionRequest struct {
	newDirection datastructure.Point
}

var gameStateReads = make(chan GameState)
var changeDirectionRequestWrites = make(chan changeDirectionRequest)
var GameOverChannel = make(chan bool)

func UpdateGameStateLoop() {
	var state = GameState{
		Snake:     []datastructure.Point{{X: 5, Y: 5}},
		Direction: datastructure.Point{Y: 1}, // Start moving right
	}

	triggerMoveSnake := time.NewTicker(configs.MoveSnakeDelay)
	defer triggerMoveSnake.Stop()

	for {
		select {
		case gameStateReads <- state:
		case newDirectionRequest := <-changeDirectionRequestWrites:
			state.Direction = newDirectionRequest.newDirection
		case <-triggerMoveSnake.C:
			handleMoveSnake(&state)
		default:
			// Add a small sleep to prevent busy-waiting
			time.Sleep(10 * time.Millisecond)
		}
	}
}
