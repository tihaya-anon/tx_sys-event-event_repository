#!/bin/bash
mkdir -p ../src/pb
protoc --proto_path=../resources/proto ../resources/proto/*.proto --go_out=../src/pb --go_opt=paths=source_relative
