package score

import (
	"math/rand"
	"time"

	"github.com/andygeiss/go-ptrace/business/player"
)

// SecurityService ...
type SecurityService struct {
	key int
}

// NewSecurityService ...
func NewSecurityService() Service {
	s := rand.NewSource(time.Now().UnixNano())
	k := int(s.Int63())
	return &SecurityService{k}
}

// Add ...
func (s *SecurityService) Add(val int, p *player.Player) {
	p.SetScore((p.GetScore() + val) ^ s.key)
}

// Del ...
func (s *SecurityService) Del(val int, p *player.Player) {
	p.SetScore((p.GetScore() - val) ^ s.key)
}

// Get ...
func (s *SecurityService) Get(p *player.Player) int {
	return p.GetScore() ^ s.key
}
