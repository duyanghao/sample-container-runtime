package nsisolation

import (
	"os"
	"path/filepath"
	"syscall"
)

// PivotRoot creates MNT namespace for container. It unmounts global filesystem, mounts local plus external
// mount given by user and prepares the container filesystem.
func PivotRoot(newRoot string) error {
	// Bind mount newRoot to itself - this is a slight hack needed to satisfy the
	// pivot_root requirement that newRoot and putold must not be on the same
	// filesystem as the current root
	oldRoot := "/.pivot_oldroot"
	putOldRoot := filepath.Join(newRoot, oldRoot)
	if err := syscall.Mount(newRoot, newRoot, "", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		return err
	}

	// Create putOldRoot directory where the old root will be stored
	if err := os.MkdirAll(putOldRoot, 0700); err != nil {
		return err
	}

	// Pivots root
	if err := syscall.PivotRoot(newRoot, putOldRoot); err != nil {
		return err
	}

	// Ensure current working directory is set to new root
	if err := os.Chdir("/"); err != nil {
		return err
	}

	// Umount old root, which now lives at /.pivot_oldroot
	if err := syscall.Unmount(oldRoot, syscall.MNT_DETACH); err != nil {
		return err
	}

	// Remove old root
	if err := os.RemoveAll(oldRoot); err != nil {
		return err
	}

	return nil
}
