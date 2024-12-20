package internal

import (
	"github.com/RobbingDaHood/my_little_snake/configs"
	"github.com/RobbingDaHood/my_little_snake/internal/datastructure"
	"testing"
)

func TestHandleMoveSnake(t *testing.T) {
	configs.GlobalRand.Seed(1) // Seed the random number generator with a fixed value
	tests := []struct {
		name           string
		initialState   GameState
		expectedState  GameState
		expectGameOver bool
	}{
		{
			name: "Move snake without collision",
			initialState: GameState{
				Snake:     []datastructure.Point{{X: 1, Y: 1}},
				Direction: datastructure.Point{X: 1, Y: 0},
			},
			expectedState: GameState{
				Snake:     []datastructure.Point{{X: 2, Y: 1}},
				Direction: datastructure.Point{X: 1, Y: 0},
				Food:      datastructure.Point{X: 0, Y: 0},
				Obstacles: []datastructure.Point{{X: 7, Y: 7}},
			},
			expectGameOver: false,
		},
		{
			name: "Move snake eating food",
			initialState: GameState{
				Snake:     []datastructure.Point{{X: 1, Y: 1}},
				Direction: datastructure.Point{X: 1, Y: 0},
				Food:      datastructure.Point{X: 2, Y: 1},
			},
			expectedState: GameState{
				Snake:     []datastructure.Point{{X: 2, Y: 1}, {X: 1, Y: 1}},
				Direction: datastructure.Point{X: 1, Y: 0},
				Food:      datastructure.Point{X: 9, Y: 1},
				Obstacles: []datastructure.Point{},
			},
			expectGameOver: false,
		},
		{
			name: "Move snake collision with obstacle",
			initialState: GameState{
				Snake:     []datastructure.Point{{X: 1, Y: 1}},
				Direction: datastructure.Point{X: 1, Y: 0},
				Obstacles: []datastructure.Point{{X: 2, Y: 1}},
			},
			expectedState: GameState{
				Snake:     []datastructure.Point{{X: 2, Y: 1}},
				Direction: datastructure.Point{X: 1, Y: 0},
				Food:      datastructure.Point{X: 0, Y: 0},
				Obstacles: []datastructure.Point{{X: 7, Y: 7}},
			},
			expectGameOver: true,
		},
		{
			name: "Move snake collision with edge",
			initialState: GameState{
				Snake:     []datastructure.Point{{X: 0, Y: 0}},
				Direction: datastructure.Point{X: -1, Y: 0},
			},
			expectedState: GameState{
				Snake:     []datastructure.Point{{X: -1, Y: 0}},
				Direction: datastructure.Point{X: -1, Y: 0},
				Food:      datastructure.Point{X: 0, Y: 0},
				Obstacles: []datastructure.Point{{X: 6, Y: 0}},
			},
			expectGameOver: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GameOverChannel = make(chan bool, 1)
			handleMoveSnake(&tt.initialState)

			if tt.expectGameOver {
				select {
				case <-GameOverChannel:
				default:
					t.Errorf("Expected game over but it did not occur")
				}
			} else {
				if len(GameOverChannel) > 0 {
					t.Errorf("Did not expect game over but it occurred")
				}
			}

			if !equalGameState(tt.initialState, tt.expectedState) {
				t.Errorf("Expected state: %+v, got: %+v", tt.expectedState, tt.initialState)
			}
		})
	}
}

func equalGameState(a, b GameState) bool {
	if len(a.Snake) != len(b.Snake) {
		return false
	}
	for i := range a.Snake {
		if a.Snake[i] != b.Snake[i] {
			return false
		}
	}
	if a.Direction != b.Direction || a.Food != b.Food {
		return false
	}
	return true
}

func TestGenerateHeadInDirection(t *testing.T) {
	tests := []struct {
		name     string
		state    GameState
		expected datastructure.Point
	}{
		{
			name: "Generate head in positive direction",
			state: GameState{
				Snake:     []datastructure.Point{{X: 1, Y: 1}},
				Direction: datastructure.Point{X: 1, Y: 0},
			},
			expected: datastructure.Point{X: 2, Y: 1},
		},
		{
			name: "Generate head in negative direction",
			state: GameState{
				Snake:     []datastructure.Point{{X: 1, Y: 1}},
				Direction: datastructure.Point{X: -1, Y: 0},
			},
			expected: datastructure.Point{X: 0, Y: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newHead := generateHeadInDirection(&tt.state)
			if newHead != tt.expected {
				t.Errorf("Expected new head at (%d, %d), got (%d, %d)", tt.expected.X, tt.expected.Y, newHead.X, newHead.Y)
			}
		})
	}
}

func TestEatFoodOrShrink(t *testing.T) {
	configs.GlobalRand.Seed(1) // Seed the random number generator with a fixed value
	tests := []struct {
		name          string
		state         GameState
		newHead       datastructure.Point
		expectedSnake []datastructure.Point
		expectedFood  datastructure.Point
	}{
		{
			name: "Eat food so does not shrink and change the food",
			state: GameState{
				Snake: []datastructure.Point{{X: 1, Y: 1}},
				Food:  datastructure.Point{X: 2, Y: 1},
			},
			newHead:       datastructure.Point{X: 2, Y: 1},
			expectedSnake: []datastructure.Point{{X: 1, Y: 1}},
			expectedFood:  datastructure.Point{X: 1, Y: 7},
		},
		{
			name: "No food, so shrink the snake and keep the food where it is",
			state: GameState{
				Snake: []datastructure.Point{{X: 1, Y: 1}},
				Food:  datastructure.Point{X: 3, Y: 1},
			},
			newHead:       datastructure.Point{X: 2, Y: 1},
			expectedSnake: []datastructure.Point{},
			expectedFood:  datastructure.Point{X: 3, Y: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eatFoodOrShrinkSnake(&tt.state, tt.newHead)
			if !equalPoints(tt.state.Snake, tt.expectedSnake) {
				t.Errorf("Expected snake: %+v, got: %+v", tt.expectedSnake, tt.state.Snake)
			}
			if tt.state.Food != tt.expectedFood {
				t.Errorf("Expected food at (%d, %d), got (%d, %d)", tt.expectedFood.X, tt.expectedFood.Y, tt.state.Food.X, tt.state.Food.Y)
			}
		})
	}
}

func equalPoints(a, b []datastructure.Point) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestMoveSnake(t *testing.T) {
	tests := []struct {
		name          string
		state         GameState
		newHead       datastructure.Point
		expectedSnake []datastructure.Point
	}{
		{
			name: "Move snake",
			state: GameState{
				Snake: []datastructure.Point{{X: 1, Y: 1}},
			},
			newHead:       datastructure.Point{X: 2, Y: 1},
			expectedSnake: []datastructure.Point{{X: 2, Y: 1}, {X: 1, Y: 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addNewHeadToSnake(&tt.state, tt.newHead)
			if !equalPoints(tt.state.Snake, tt.expectedSnake) {
				t.Errorf("Expected snake: %+v, got: %+v", tt.expectedSnake, tt.state.Snake)
			}
		})
	}
}

func TestCheckCollision(t *testing.T) {
	tests := []struct {
		name           string
		state          GameState
		newHead        datastructure.Point
		expectGameOver bool
	}{
		{
			name: "Collision with wall",
			state: GameState{
				Snake: []datastructure.Point{{X: 1, Y: 1}},
			},
			newHead:        datastructure.Point{X: -1, Y: 1},
			expectGameOver: true,
		},
		{
			name: "Collision with itself",
			state: GameState{
				Snake: []datastructure.Point{{X: 1, Y: 1}, {X: 2, Y: 1}},
			},
			newHead:        datastructure.Point{X: 1, Y: 1},
			expectGameOver: true,
		},
		{
			name: "Collision with obstacle",
			state: GameState{
				Snake:     []datastructure.Point{{X: 1, Y: 1}},
				Obstacles: []datastructure.Point{{X: 2, Y: 2}},
			},
			newHead:        datastructure.Point{X: 2, Y: 2},
			expectGameOver: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GameOverChannel = make(chan bool, 1)
			checkCollision(&tt.state, tt.newHead)

			if tt.expectGameOver {
				select {
				case <-GameOverChannel:
				default:
					t.Errorf("Expected game over but it did not occur")
				}
			} else {
				if len(GameOverChannel) > 0 {
					t.Errorf("Did not expect game over but it occurred")
				}
			}
		})
	}
}
