#!/bin/bash

make build
./build/pkg/cmd/sample-container-runtime/sample-container-runtime run -ti --name container1 busybox sh
