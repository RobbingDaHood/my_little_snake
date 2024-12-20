package internal

import (
	"github.com/RobbingDaHood/my_little_snake/configs"
	"github.com/RobbingDaHood/my_little_snake/internal/datastructure"
)

func handleMoveSnake(state *GameState) {
	newHead := generateHeadInDirection(state)
	checkCollision(state, newHead)
	addNewHeadToSnake(state, newHead)
	eatFoodOrShrinkSnake(state, newHead)
	maybeAddNewObstacle(state)
}

func generateHeadInDirection(state *GameState) datastructure.Point {
	return datastructure.Point{X: state.Snake[0].X + state.Direction.X, Y: state.Snake[0].Y + state.Direction.Y}
}

func eatFoodOrShrinkSnake(state *GameState, newHead datastructure.Point) {
	if newHead == state.Food {
		spawnFood(state)
	} else {
		state.Snake = state.Snake[:len(state.Snake)-1]
	}
}

func addNewHeadToSnake(state *GameState, newHead datastructure.Point) {
	state.Snake = append([]datastructure.Point{newHead}, state.Snake...)
}

func checkCollision(state *GameState, newHead datastructure.Point) {
	if newHead.X < 0 || newHead.Y < 0 || newHead.X >= configs.GridHeight || newHead.Y >= configs.GridWidth {
		GameOverChannel <- true
	}
	for _, part := range state.Snake {
		if part == newHead {
			GameOverChannel <- true
		}
	}
	for _, obs := range state.Obstacles {
		if obs == newHead {
			GameOverChannel <- true
		}
	}
}
