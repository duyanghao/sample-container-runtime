package nsisolation

import (
	"fmt"
	"syscall"
)

// PrepareUser switches to UID=0, GID=0 and become root within container
func PrepareUser() error {
	if err := syscall.Setgroups([]int{0}); err != nil {
		return fmt.Errorf("syscall Setgroups failure: %v", err)
	}
	if err := syscall.Setresgid(0, 0, 0); err != nil {
		return fmt.Errorf("syscall Setresgid failure: %v", err)
	}
	if err := syscall.Setresuid(0, 0, 0); err != nil {
		return fmt.Errorf("syscall Setresuid failure: %v", err)
	}
	return nil
}
