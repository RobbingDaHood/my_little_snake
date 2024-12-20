package internal

import (
	"github.com/RobbingDaHood/my_little_snake/configs"
	"github.com/RobbingDaHood/my_little_snake/internal/datastructure"
)

func spawnFood(state *GameState) {
	state.Food = datastructure.Point{X: configs.GlobalRand.Intn(configs.GridHeight), Y: configs.GlobalRand.Intn(configs.GridWidth)}
}
