package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/andygeiss/go-ptrace/infrastructure/process"
)

func main() {
	// read input parameters
	flagVal := flag.String("value", "12345678", "64bit int value")
	flagPID := flag.String("pid", "0", "target process PID")
	flag.Parse()
	val64, err := strconv.ParseInt(*flagVal, 0, 64)
	pid64, err := strconv.ParseInt(*flagPID, 0, 64)
	wanted := uint64(val64)
	pid := int(pid64)
	//
	wpid, err := process.Attach(pid)
	if err != nil {
		log.Fatalf("Attach error: %s\n", err.Error())
	}
	fmt.Printf("Process [%d] attached.\n", pid)
	// read data from process memory
	start := 0x0000000000585000 // anon
	size := 1024 * 124          // 124K
	end := start + size
	counter := 0
	for addr := start; addr < end; addr += 8 {
		bs, err := process.Read(wpid, int64(addr), 8)
		if err != nil {
			log.Fatalf("Read error: %s\n", err.Error())
		}
		value := binary.LittleEndian.Uint64(bs)
		if value == wanted {
			// count occurence of wanted value
			fmt.Printf("Value found at [0x%x].\n", addr)
			counter++
		}
	}
	// detach process and exit
	err = process.Detach(wpid)
	if err != nil {
		log.Fatalf("Detach error: %s\n", err.Error())
	}
	fmt.Printf("Process [%d] detached.\n", wpid)
	fmt.Printf("Value found [%d] times.\n", counter)
}
