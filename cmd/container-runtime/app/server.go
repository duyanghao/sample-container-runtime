/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package app implements a Server object for running the scheduler.
package app

import (
	"context"
	"github.com/duyanghao/sample-container-runtime/pkg/runtime"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

// NewSchedulerCommand creates a *cobra.Command object with default parameters and registryOptions
func NewContainerRuntimeCommand() *cobra.Command {
	// TODO: init command options
	cmd := &cobra.Command{
		Use:  "sample-runtime-container",
		Long: `This repository implements a simple container runtime for learning purposes.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := runContainerRuntime(cmd, args); err != nil {
				log.Errorf("runContainerRuntime failed: %v", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

// runCommand runs the container runtime.
func runContainerRuntime(cmd *cobra.Command, args []string) error {
	containerRuntime, err := runtime.New(runtime.WithCommand(args[0]), runtime.WithRootfsDir(args[1]))
	if err != nil {
		log.Errorf("Create ContainerRuntime failed: %v", err)
		return err
	}
	containerRuntime.Run(context.Background())
	return nil
}
