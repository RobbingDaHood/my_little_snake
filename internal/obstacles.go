package internal

import (
	"github.com/RobbingDaHood/my_little_snake/configs"
	"github.com/RobbingDaHood/my_little_snake/internal/datastructure"
)

func maybeAddNewObstacle(state *GameState) {
	if len(state.Obstacles) < configs.MaxObstacles && configs.GlobalRand.Intn(100) > configs.PercentageRiskOfNewObstacle {
		spawnObstacle(state)
	}
}

func spawnObstacle(state *GameState) {
	state.Obstacles = append(state.Obstacles, datastructure.Point{X: configs.GlobalRand.Intn(configs.GridHeight), Y: configs.GlobalRand.Intn(configs.GridWidth)})
}
