package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/andygeiss/go-ptrace/infrastructure/process"
)

func main() {
	// read input parameters
	flagAddr := flag.Int64("addr", 0, "64 bit start address")
	flagLen := flag.Int64("len", 0, "search length")
	flagPID := flag.Int64("pid", 0, "target process PID")
	flagVal := flag.Int64("value", 0, "64 bit int value")
	flag.Parse()
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage: %s\n", args[0])
		flag.PrintDefaults()
		return
	}
	//
	// convert args to types needed
	addr := int64(*flagAddr)
	len := int64(*flagLen)
	pid := int(*flagPID)
	value := uint64(*flagVal)
	//
	wpid, err := process.Attach(pid)
	if err != nil {
		log.Fatalf("Attach error: %s\n", err.Error())
	}
	fmt.Printf("Process [%d] attached.\n", pid)
	// read data from process memory
	fmt.Printf("Searching for value [dec: %d][hex: 0x%x] ...\n", value, value)
	start := addr // anon
	size := len   // size in KByte
	end := start + size
	counter := 0
	for addr := start; addr < end; addr += 8 {
		bs, err := process.Read(wpid, int64(addr), 8)
		if err != nil {
			log.Fatalf("Read error: %s\n", err.Error())
		}
		cur := binary.LittleEndian.Uint64(bs)
		if cur == value {
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
