package main

import (
	"fmt"
	"time"
)

// Secret should be hacked! ;-)
var Secret int

func main() {
	Secret = 0x12345678abcdef00
	for {
		time.Sleep(time.Second * 1)
		fmt.Printf("Hackme [Secret: 0x%x, Address: %p]\n", Secret, &Secret)
	}
}
