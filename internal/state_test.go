package internal

import (
	"github.com/RobbingDaHood/my_little_snake/configs"
	"github.com/RobbingDaHood/my_little_snake/internal/datastructure"
	"testing"
	"time"
)

func TestUpdateGameStateLoop(t *testing.T) {
	configs.GlobalRand.Seed(1) // Seed the random number generator with a fixed value
	tests := []struct {
		name           string
		initialState   GameState
		direction      datastructure.Point
		expectedState  GameState
		expectGameOver bool
	}{
		{
			name: "Initial state with no direction change",
			initialState: GameState{
				Snake:     []datastructure.Point{{X: 5, Y: 5}},
				Direction: datastructure.Point{Y: 1},
			},
			direction: datastructure.Point{Y: 1},
			expectedState: GameState{
				Snake:     []datastructure.Point{{X: 5, Y: 6}},
				Direction: datastructure.Point{Y: 1},
			},
			expectGameOver: false,
		},
		{
			name: "Change direction and move",
			initialState: GameState{
				Snake:     []datastructure.Point{{X: 5, Y: 5}},
				Direction: datastructure.Point{Y: 1},
			},
			direction: datastructure.Point{X: 1},
			expectedState: GameState{
				Snake:     []datastructure.Point{{X: 6, Y: 5}},
				Direction: datastructure.Point{X: 1},
			},
			expectGameOver: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameStateReads = make(chan GameState, 1)
			changeDirectionRequestWrites = make(chan changeDirectionRequest, 1)
			GameOverChannel = make(chan bool, 1)

			changeDirectionRequestWrites <- changeDirectionRequest{newDirection: tt.direction}
			go UpdateGameStateLoop()

			var state GameState
			var gameOver bool
			fiveSecondTimer := time.Now().Add(5 * time.Second)
			for {
				select {
				case state = <-gameStateReads:
				case <-time.After(100 * time.Millisecond):
					t.Fatal("Timed out waiting for game state")
				}

				select {
				case gameOver = <-GameOverChannel:
				case <-time.After(100 * time.Millisecond): // Don't wait forever
				}

				if state.Direction == tt.direction && gameOver == tt.expectGameOver && equalGameState(state, tt.expectedState) {
					break
				} else if time.Now().After(fiveSecondTimer) {
					t.Fatal("Timed out waiting for game state")
				}
			}
		})
	}
}
