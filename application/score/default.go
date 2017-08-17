package score

import "github.com/andygeiss/go-memhack/business/player"

// Service ...
type Service interface {
	Add(val int, p *player.Player)
	Sub(val int, p *player.Player)
	Get(p *player.Player) int
}

// DefaultService ...
type DefaultService struct{}

// NewDefaultService ...
func NewDefaultService() Service {
	return &DefaultService{}
}

// Add ...
func (s *DefaultService) Add(val int, p *player.Player) {
	p.SetScore(p.GetScore() + val)
}

// Sub ...
func (s *DefaultService) Sub(val int, p *player.Player) {
	p.SetScore(p.GetScore() - val)
}

// Get ...
func (s *DefaultService) Get(p *player.Player) int {
	return p.GetScore()
}
