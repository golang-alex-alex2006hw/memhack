package process

import "syscall"

// Write ...
func Write(wpid int, addr int64, data []byte) error {
	_, err := syscall.PtracePokeData(wpid, uintptr(addr), data[:])
	if err != nil {
		return err
	}
	return nil
}
