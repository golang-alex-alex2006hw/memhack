package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

func main() {
	//
	addr := os.Args[1]
	this := os.Args[2]
	args := os.Args[3:]
	//
	cmd, err := runCommand(this, args)
	//
	ptr := getPointer(addr)
	//
	var regs syscall.PtraceRegs
	pid := cmd.Process.Pid
	exit := true
	for {
		if exit {
			err = syscall.PtraceGetRegs(pid, &regs)
			// no such process
			if err != nil {
				break
			}
		}
		hackAdmin(pid, ptr)
		nextSyscall(pid)
		wait4Syscall(pid)
		exit = !exit
	}
	fmt.Print("Exit.") // possibly never reached
}

// transform string address into pointer
func getPointer(addr string) int64 {
	dst, err := strconv.ParseInt(addr, 0, 64)
	if err != nil {
		panic(err)
	}
	return dst
}

// set bool value of IsAdmin to true
func hackAdmin(pid int, addr int64) {
	//buf := make([]byte, 1)
	//_, err = syscall.PtracePeekData(pid, 0x59fac2, buf[:1])
	_, err := syscall.PtracePokeData(pid, uintptr(addr), []byte{1})
	if err != nil {
		panic(err)
	}
}

// trace next syscall
func nextSyscall(pid int) {
	err := syscall.PtraceSyscall(pid, 0)
	if err != nil {
		panic(err)
	}
}

// prepare & run command
func runCommand(this string, args []string) (*exec.Cmd, error) {
	fmt.Printf("Running %v ...\n", this)
	cmd := exec.Command(this, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Ptrace: true,
	}
	cmd.Start()
	err := cmd.Wait()
	if err != nil {
		fmt.Printf("Wait returned: %v\n", err)
	}
	return cmd, err
}

// wait for syscall exit
func wait4Syscall(pid int) {
	_, err := syscall.Wait4(pid, nil, 0, nil)
	if err != nil {
		panic(err)
	}
}
