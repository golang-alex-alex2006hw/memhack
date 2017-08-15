package main

import (
	"fmt"
	"time"
)

// IsAdmin should be hacked! ;-)
var IsAdmin bool

func main() {
	IsAdmin = false
	for {
		time.Sleep(time.Second * 1)
		fmt.Printf("Hackme [Admin: %v, Address: %p]\n", IsAdmin, &IsAdmin)
	}
}
