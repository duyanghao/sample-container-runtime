package nsisolation

import (
	"os"
	"syscall"
)

// ProcPrepare creates a new /proc directory in the container's file system, which has just been mounted.
// Additionally, it mounts the proc mount of a parent process to the /proc of the container.
// The proc mount exists on the list returned by mount command invoked from a host system:
// proc on /proc type proc (rw,nosuid,nodev,noexec,relatime)
func PrepareProc() error {
	source := "proc"
	target := "/proc"
	fsType := "proc"
	mntFlags := 0
	data := ""

	os.MkdirAll(target, 0755)
	if err := syscall.Mount(source, target, fsType, uintptr(mntFlags), data); err != nil {
		return err
	}

	return nil
}
