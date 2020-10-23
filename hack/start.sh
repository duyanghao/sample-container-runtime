#!/bin/bash

make build
./build/pkg/cmd/sample-container-runtime/sample-container-runtime run -d --name container1 -v /root/tmp/from1:/to1 busybox top