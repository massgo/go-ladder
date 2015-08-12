package ladder

import (
	"time"
)

type blackOrWhite int

const (
	black = iota
	white
)

type Result struct {
	white  Player
	black  Player
	winner blackOrWhite
	time   time.Time
}
