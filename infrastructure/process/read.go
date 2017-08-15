package process

import (
	"errors"
	"syscall"
)

// Read ...
func Read(wpid int, addr int64, len int) ([]byte, error) {
	out := make([]byte, len)
	count, err := syscall.PtracePeekData(wpid, uintptr(addr), out)
	if err != nil {
		return nil, err
	}
	if count < len {
		return nil, errors.New("count and len are not equal")
	}
	return out, nil
}
