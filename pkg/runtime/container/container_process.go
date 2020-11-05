package container

import (
	"fmt"
	"github.com/duyanghao/sample-container-runtime/pkg/runtime/cgroups/subsystems"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"syscall"
)

var (
	RUNNING             string = "running"
	STOP                string = "stopped"
	Exit                string = "exited"
	DefaultInfoLocation string = "/var/run/sample-container-runtime/%s/"
	ConfigName          string = "config.json"
	ContainerLogFile    string = "container.log"
	RootUrl             string = "/var/lib/sample-container-runtime"
	MntUrl              string = "/var/lib/sample-container-runtime/mnt/%s"
	WriteLayerUrl       string = "/var/lib/sample-container-runtime/writeLayer/%s"
)

type ContainerInfo struct {
	Pid         string                     `json:"pid"`         // 容器的init进程在宿主机上的 PID
	Id          string                     `json:"id"`          // 容器Id
	Name        string                     `json:"name"`        // 容器名
	Command     []string                   `json:"command"`     // 容器内init运行命令
	CreatedTime string                     `json:"createTime"`  // 创建时间
	Status      string                     `json:"status"`      // 容器的状态
	Volume      string                     `json:"volume"`      // 容器的数据卷
	PortMapping []string                   `json:"portmapping"` // 端口映射
	ImageName   string                     `json:"imageName"`   // 镜像名
	Detached    bool                       `json:"detached"`    // 是否后端执行
	ResConf     *subsystems.ResourceConfig `json:"resConf"`     // cgroup限制
	Env         []string                   `json:"env"`         // 容器环境变量
	Network     string                     `json:"network"`     // 容器网络
}

func NewParentProcess(tty bool, containerName, volume, imageName string, envSlice []string) (*exec.Cmd, *os.File) {
	readPipe, writePipe, err := NewPipe()
	if err != nil {
		log.Errorf("New pipe error %v", err)
		return nil, nil
	}
	initCmd, err := os.Readlink("/proc/self/exe")
	if err != nil {
		log.Errorf("get init process error %v", err)
		return nil, nil
	}

	cmd := exec.Command(initCmd, "init")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
		/*Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC | syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},*/
	}

	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		dirURL := fmt.Sprintf(DefaultInfoLocation, containerName)
		if err := os.MkdirAll(dirURL, 0622); err != nil {
			log.Errorf("NewParentProcess mkdir %s error %v", dirURL, err)
			return nil, nil
		}
		stdLogFilePath := dirURL + ContainerLogFile
		stdLogFile, err := os.Create(stdLogFilePath)
		if err != nil {
			log.Errorf("NewParentProcess create file %s error %v", stdLogFilePath, err)
			return nil, nil
		}
		cmd.Stdout = stdLogFile
	}

	cmd.ExtraFiles = []*os.File{readPipe}
	cmd.Env = append(os.Environ(), envSlice...)
	NewWorkSpace(volume, imageName, containerName)
	cmd.Dir = fmt.Sprintf(MntUrl, containerName)
	return cmd, writePipe
}

func NewPipe() (*os.File, *os.File, error) {
	read, write, err := os.Pipe()
	if err != nil {
		return nil, nil, err
	}
	return read, write, nil
}
