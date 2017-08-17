package score

import "github.com/andygeiss/go-ptrace/business/player"

// Service ...
type Service interface {
	Add(val int, p *player.Player)
	Del(val int, p *player.Player)
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

// Del ...
func (s *DefaultService) Del(val int, p *player.Player) {
	p.SetScore(p.GetScore() - val)
}

// Get ...
func (s *DefaultService) Get(p *player.Player) int {
	return p.GetScore()
}
