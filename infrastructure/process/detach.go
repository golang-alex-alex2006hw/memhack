package process

import (
	"syscall"
)

// Detach from working process
func Detach(wpid int) error {
	err := syscall.PtraceDetach(wpid)
	if err != nil {
		return err
	}
	return nil
}
