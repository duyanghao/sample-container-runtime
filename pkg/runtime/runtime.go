package runtime

import (
	"context"
	log "github.com/sirupsen/logrus"
)

// ContainerRuntime generates a sample container
type ContainerRuntime struct {
	crOpts containerRuntimeOptions
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
	return &ContainerRuntime{crOpts: options}, nil
}

// Run begins creating a sample container
func (cr *ContainerRuntime) Run(ctx context.Context) {
	log.Println(cr.crOpts)
}
