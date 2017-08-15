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
	flagAddr := flag.String("addr", "0x12345678", "target process address")
	flagPID := flag.String("pid", "0", "target process PID")
	flagVal := flag.String("value", "12345678", "64bit int value")
	flag.Parse()
	addr64, err := strconv.ParseInt(*flagAddr, 0, 64)
	pid64, err := strconv.ParseInt(*flagPID, 0, 64)
	pid := int(pid64)
	val64, err := strconv.ParseInt(*flagVal, 0, 64)
	wanted := uint64(val64)
	//
	wpid, err := process.Attach(pid)
	if err != nil {
		log.Fatalf("Attach error: %s\n", err.Error())
	}
	fmt.Printf("Process [%d] attached.\n", pid)
	// write data to process memory
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, wanted)
	err = process.Write(wpid, addr64, bs)
	if err != nil {
		log.Fatalf("Write error: %s\n", err.Error())
	}
	// read data from process memory
	data, err := process.Read(wpid, addr64, 8)
	if err != nil {
		log.Fatalf("Read error: %s\n", err.Error())
	}
	value := binary.LittleEndian.Uint64(data)
	fmt.Printf("Value at [0x%x] is now [%d].\n", addr64, value)
	// detach process and exit
	err = process.Detach(wpid)
	if err != nil {
		log.Fatalf("Detach error: %s\n", err.Error())
	}
	fmt.Printf("Process [%d] detached.\n", wpid)
}
