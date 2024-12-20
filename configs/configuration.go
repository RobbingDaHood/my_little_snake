package configs

import (
	"math/rand"
	"time"
)

const (
	GridWidth                   = 10
	GridHeight                  = 10
	MaxObstacles                = 5
	PercentageRiskOfNewObstacle = 30
	OsClearCommand              = "clear"
	UpdateUiDelay               = 20 * time.Millisecond
	MoveSnakeDelay              = 500 * time.Millisecond
)

// GlobalRand Both good to only initialize once and it makes it possible to test the code using random by seeding this.
var GlobalRand = rand.New(rand.NewSource(time.Now().UnixNano()))
