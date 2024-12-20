package internal

import (
	"bytes"
	"github.com/RobbingDaHood/my_little_snake/internal/datastructure"
	"os"
	"testing"
)

func TestRender(t *testing.T) {
	state := GameState{
		Food:      datastructure.Point{X: 1, Y: 1},
		Snake:     []datastructure.Point{{X: 2, Y: 2}},
		Obstacles: []datastructure.Point{{X: 3, Y: 3}},
	}
	expectedOutput := []string{
		"..........",
		".F........",
		"..S.......",
		"...X......",
		"..........",
		"..........",
		"..........",
		"..........",
		"..........",
		"..........",
	}

	// Mock the gameStateReads channel
	gameStateReads = make(chan GameState, 1)
	gameStateReads <- state

	// Create a pipe to capture the output
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}

	// Redirect stdout to the write end of the pipe
	old := os.Stdout
	os.Stdout = w

	// Call the render function
	render()

	// Close the write end of the pipe and restore stdout
	err = w.Close()
	if err != nil {
		t.Fatalf("Failed to close pipe: %v", err)
	}
	os.Stdout = old

	// Read the captured output from the read end of the pipe
	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("Failed to read from pipe: %v", err)
	}

	// Check the captured output
	outputLines := bytes.Split(buf.Bytes(), []byte("\n"))
	for i, line := range expectedOutput {
		if string(outputLines[i]) != line {
			t.Errorf("Expected output:\n%s\nGot:\n%s", line, string(outputLines[i]))
		}
	}
}
