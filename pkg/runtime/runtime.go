package runtime

import (
	"context"
	"github.com/duyanghao/sample-container-runtime/pkg/runtime/nsisolation"
	"github.com/duyanghao/sample-container-runtime/pkg/runtime/util"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"syscall"

	"github.com/docker/docker/pkg/reexec"
)

// ContainerRuntime generates a sample container
type ContainerRuntime struct {
	// Container hostname
	Hostname string
	// Container startup command
	Command string
	// Container rootfs directory
	RootfsDir string
}

type containerRuntimeOptions struct {
	// Container startup command
	command string
	// Container rootfs directory
	rootfsDir string
}

// Option configures a ContainerRuntime
type Option func(*containerRuntimeOptions)

// WithCommand sets command for ContainerRuntime
func WithCommand(command string) Option {
	return func(o *containerRuntimeOptions) {
		o.command = command
	}
}

// WithRootfsDir sets rootfs directory for ContainerRuntime
func WithRootfsDir(rootfsDir string) Option {
	return func(o *containerRuntimeOptions) {
		o.rootfsDir = rootfsDir
	}
}

// New returns a ContainerRuntime
func New(opts ...Option) (*ContainerRuntime, error) {
	var options containerRuntimeOptions
	for _, opt := range opts {
		opt(&options)
	}
	return &ContainerRuntime{
		Hostname:  util.RandomSeq(10),
		Command:   options.command,
		RootfsDir: options.rootfsDir,
	}, nil
}

func init() {
	reexec.Register("nsInit", nsInit)
	if reexec.Init() {
		os.Exit(0)
	}
}

// nsInit prepares child process namespace isolation initialization work and exec container command
func nsInit() {
	// Set container hostname
	hostname := util.RandomSeq(10)
	if err := syscall.Sethostname([]byte(hostname)); err != nil {
		log.Errorf("setting hostname failure: %v", err)
		os.Exit(1)
	}
	// Set container new rootfs
	newRoot := os.Args[1]
	if err := nsisolation.PivotRoot(newRoot); err != nil {
		log.Errorf("pivoting container rootfs failure: %v", err)
		os.Exit(1)
	}
	// Execute container command
	command := os.Args[2]
	containerRun(command)
}

// containerRun executes command normally
func containerRun(command string) {
	cmd := exec.Command(command)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Errorf("running container command: %s failure: %v", command, err)
		os.Exit(1)
	}
}

// createChildProcess creates a child process and waits it out
func (cr *ContainerRuntime) createChildProcess(ctx context.Context) error {
	cmd := reexec.Command("nsInit", cr.RootfsDir, cr.Command)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Errorf("starting the reexec.Command failure: %v", err)
		return err
	}
	if err := cmd.Wait(); err != nil {
		log.Errorf("waiting for the reexec.Command failure: %v", err)
		return err
	}
	log.Info("container exit normally")
	return nil
}

// Run begins creating a sample container
func (cr *ContainerRuntime) Run(ctx context.Context) {
	if err := cr.createChildProcess(ctx); err != nil {
		log.Errorf("createChildProcess failed: %v", err)
	}
}
