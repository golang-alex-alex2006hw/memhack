package main

import (
	"fmt"
	"time"
)

// IsAdmin should be hacked! ;-)
var IsAdmin bool

func main() {
	IsAdmin = false
	times := 0
	// We have 10 seconds to hack IsAdmin ;-)
	for {
		time.Sleep(time.Second * 1)
		fmt.Printf("Hackme [Admin: %v, Address: %p]\n", IsAdmin, &IsAdmin)
		times++
		if times > 10 {
			break
		}
	}
}
