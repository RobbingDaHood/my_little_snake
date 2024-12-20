package internal

import (
	"github.com/RobbingDaHood/my_little_snake/configs"
	"github.com/RobbingDaHood/my_little_snake/internal/datastructure"
	"testing"
)

func TestSpawnFood(t *testing.T) {
	configs.GlobalRand.Seed(1) // Seed the random number generator with a fixed value
	state := &GameState{
		Food: datastructure.Point{X: 0, Y: 0},
	}

	spawnFood(state)

	if state.Food.X == 2 {
		t.Errorf("Fixed the random with a seed so will always expect X to be %d, but it were %d", 2, state.Food.X)
	}

	if state.Food.Y == 2 {
		t.Errorf("Fixed the random with a seed so will always expect Y to be %d, but it were %d", 2, state.Food.Y)
	}
}

func TestSpawnFoodBoundary(t *testing.T) {
	configs.GlobalRand.Seed(1) // Seed the random number generator with a fixed value
	state := &GameState{
		Food: datastructure.Point{X: 0, Y: 0},
	}

	// Just iterating a lot to see if the food is within the bounds
	for range 100 {
		spawnFood(state)

		if state.Food.X < 0 || state.Food.X >= configs.GridHeight {
			t.Errorf("Food X coordinate out of bounds: got %d, want [0, %d)", state.Food.X, configs.GridHeight)
		}

		if state.Food.Y < 0 || state.Food.Y >= configs.GridWidth {
			t.Errorf("Food Y coordinate out of bounds: got %d, want [0, %d)", state.Food.Y, configs.GridWidth)
		}
	}
}
