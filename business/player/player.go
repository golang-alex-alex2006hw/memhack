package player

// Player stores the value object.
type Player struct {
	Name  string
	Score int
}

// NewPlayer creates a new player and returns its address.
func NewPlayer(name string) *Player {
	// each players starts with a score of 0.
	return &Player{name, 0}
}

// GetScore ...
func (p *Player) GetScore() int {
	return p.Score
}

// SetScore ...
func (p *Player) SetScore(val int) {
	p.Score = val
}
