package nsisolation

import (
	"fmt"
	"github.com/duyanghao/sample-container-runtime/pkg/runtime/constant"
	"io/ioutil"
	"os"
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

// PrepareUid2GidMap configures relevant uid_map and gid_map contents for child process
func PrepareUid2GidMap(pid, uid, gid int, pipeW *os.File) error {
	// Configure uid_map
	uidMapPath := fmt.Sprintf("/proc/%d/uid_map", pid)
	if err := ioutil.WriteFile(uidMapPath, []byte(fmt.Sprintf("0 %d 1", uid)), 0644); err != nil {
		return fmt.Errorf("write uid_map: %s failure: %v", uidMapPath, err)
	}
	// Configure gid_map
	gidMapPath := fmt.Sprintf("/proc/%d/gid_map", pid)
	if err := ioutil.WriteFile(gidMapPath, []byte(fmt.Sprintf("0 %d 1", gid)), 0644); err != nil {
		return fmt.Errorf("write gid_map: %s failure: %v", gidMapPath, err)
	}
	// Send signal to child process
	_, err := pipeW.Write([]byte(constant.UID2GIDMAPDONE))
	if err != nil {
		return fmt.Errorf("send signal to child process failure: %v", err)
	}
	return nil
}

// WaitPrepareUid2GidMap waits for the parent to configure user/group ID mappings. If data received, move on.
func WaitPrepareUid2GidMap(signalPath string) error {
	pipeR, err := os.Open(signalPath)
	if err != nil {
		return fmt.Errorf("open signal pipe: %s failure: %v", signalPath, err)
	}
	defer pipeR.Close()
	data := make([]byte, len(constant.UID2GIDMAPDONE))
	if _, err := pipeR.Read(data); err != nil {
		return fmt.Errorf("read pipe failure: %v", err)
	}
	if string(data) != constant.UID2GIDMAPDONE {
		return fmt.Errorf("received invalid data: %s", string(data))
	}
	return nil
}
