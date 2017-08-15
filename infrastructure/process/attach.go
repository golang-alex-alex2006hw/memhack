package process

import (
	"os"
	"syscall"
)

// Attach target process and return its working PID.
func Attach(pid int) (int, error) {
	proc, err := os.FindProcess(pid)
	if err != nil {
		return 0, err
	}
	wpid := proc.Pid
	err = syscall.PtraceAttach(wpid)
	if err != nil {
		return 0, err
	}
	return wpid, err
}
