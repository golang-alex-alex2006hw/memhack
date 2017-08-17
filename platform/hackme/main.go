package main

import (
	"fmt"
	"time"

	"github.com/andygeiss/go-ptrace/application/score"
	"github.com/andygeiss/go-ptrace/business/player"
)

func main() {
	var s score.Service
	//s = score.NewDefaultService()
	s = score.NewSecurityService()
	var p *player.Player
	p = player.NewPlayer("you")
	s.Add(12345678, p)
	for {
		time.Sleep(time.Second * 1)
		fmt.Printf("Hackme [Player: 0x%x, Score: %d (%p)]\n", &p, s.Get(p), &p.Score)
	}
}
