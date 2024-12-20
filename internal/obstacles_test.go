package internal

import (
	"github.com/RobbingDaHood/my_little_snake/configs"
	"github.com/RobbingDaHood/my_little_snake/internal/datastructure"
	"math/rand"
	"testing"
)

func TestMaybeAddNewObstacle(t *testing.T) {
	// Set up the global random generator with a fixed seed for reproducibility
	configs.GlobalRand = rand.New(rand.NewSource(1))

	state := &GameState{
		Obstacles: []datastructure.Point{},
	}

	maybeAddNewObstacle(state)

	if len(state.Obstacles) != 1 {
		t.Errorf("Expected 1 obstacle, got %d", len(state.Obstacles))
	}

}

func TestMaybeAddNewObstacleMax(t *testing.T) {
	// Set up the global random generator with a fixed seed for reproducibility
	configs.GlobalRand = rand.New(rand.NewSource(1))

	state := &GameState{
		Obstacles: make([]datastructure.Point, configs.MaxObstacles),
	}

	maybeAddNewObstacle(state)

	if len(state.Obstacles) != configs.MaxObstacles {
		t.Errorf("Expected %d obstacles, got %d", configs.MaxObstacles, len(state.Obstacles))
	}
}

func TestSpawnObstacle(t *testing.T) {
	// Set up the global random generator with a fixed seed for reproducibility
	configs.GlobalRand = rand.New(rand.NewSource(1))

	state := &GameState{
		Obstacles: []datastructure.Point{},
	}

	// Just iterating a lot to see if the obstacles are within the bounds
	for i := 1; i <= 100; i++ {
		spawnObstacle(state)

		if len(state.Obstacles) != i {
			t.Errorf("Expected %v obstacle, got %d", i, len(state.Obstacles))
		}

		if state.Obstacles[0].X < 0 || state.Obstacles[0].X >= configs.GridHeight {
			t.Errorf("Obstacle X coordinate out of bounds: got %d, want [0, %d)", state.Obstacles[0].X, configs.GridHeight)
		}

		if state.Obstacles[0].Y < 0 || state.Obstacles[0].Y >= configs.GridWidth {
			t.Errorf("Obstacle Y coordinate out of bounds: got %d, want [0, %d)", state.Obstacles[0].Y, configs.GridWidth)
		}
	}
}
