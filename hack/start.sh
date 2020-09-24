#!/bin/bash

make build
./build/pkg/cmd/sample-container-runtime/sample-container-runtime bosybox /bin/bash