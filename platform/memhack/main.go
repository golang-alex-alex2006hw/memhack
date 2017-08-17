package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/andygeiss/go-memhack/infrastructure/process"
)

func main() {
	// read input parameters
	flagAddr := flag.Int64("addr", 0, "64 bit target address")
	flagPID := flag.Int64("pid", 0, "target process PID")
	flagVal := flag.Int64("value", 0, "64 bit int value")
	flag.Parse()
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage: %s\n", args[0])
		flag.PrintDefaults()
		return
	}
	// convert args to types needed
	addr := int64(*flagAddr)
	pid := int(*flagPID)
	value := uint64(*flagVal)
	// attach target process
	wpid, err := process.Attach(pid)
	if err != nil {
		log.Fatalf("Attach error: %s\n", err.Error())
	}
	fmt.Printf("Process [%d] attached.\n", pid)
	// write data to process memory
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, value)
	err = process.Write(wpid, addr, bs)
	if err != nil {
		log.Fatalf("Write error: %s\n", err.Error())
	}
	// detach process and exit
	err = process.Detach(wpid)
	if err != nil {
		log.Fatalf("Detach error: %s\n", err.Error())
	}
	fmt.Printf("Process [%d] detached.\n", wpid)
}
