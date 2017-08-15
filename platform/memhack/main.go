package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

func main() {
	// Prepare command execution
	fmt.Printf("Running %v ...\n", os.Args[2:])
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Ptrace: true,
	}
	addr, err := strconv.ParseInt(os.Args[1], 0, 64)
	if err != nil {
		panic(err)
	}
	// Start command
	cmd.Start()
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Wait returned: %v\n", err)
	}
	// Main loop
	var regs syscall.PtraceRegs
	pid := cmd.Process.Pid
	exit := true
	for {
		// If error occured (no such process) then leave main loop
		if exit {
			err = syscall.PtraceGetRegs(pid, &regs)
			if err != nil {
				break
			}
		}
		//buf := make([]byte, 1)
		//_, err = syscall.PtracePeekData(pid, 0x59fac2, buf[:1])

		// Set bool value of IsAdmin to true.
		syscall.PtracePokeData(pid, uintptr(addr), []byte{1})
		if err != nil {
			panic(err)
		}
		// Trace next syscall
		err = syscall.PtraceSyscall(pid, 0)
		if err != nil {
			panic(err)
		}
		// Wait for syscall exit
		_, err = syscall.Wait4(pid, nil, 0, nil)
		if err != nil {
			panic(err)
		}
		exit = !exit
	}
	fmt.Print("Exit.")
}
