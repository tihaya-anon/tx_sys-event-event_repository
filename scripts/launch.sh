#!/bin/bash
SCRIPT_DIR=$(cd $(dirname $0); pwd)
cd $SCRIPT_DIR/../src/cmd/
go mod tidy
echo "Building app..."
go build -o ../../build/app
cd $SCRIPT_DIR
../build/app