package ladder

type Ladder struct {
	players            map[int]Player
	results            []Result
	currentPlayerIndex int
}
