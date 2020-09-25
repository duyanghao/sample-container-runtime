package nsisolation

import (
	"os"
	"path/filepath"
)

func ProcPrepare(newRoot string) error {
	source := "proc"
	target := filepath.Join(newRoot, "/proc")
	fsType := "proc"
	mntFlags := 0
	data := ""

	os.MkdirAll(target, 0755)
	if err := syscall.Mount(source, target, fsType, uintptr(mntFlags), data); err != nil {
		return err
	}

	return nil
}
