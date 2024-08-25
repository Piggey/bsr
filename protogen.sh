#!/bin/bash

PROTO_DIR="./packet/proto"

cd $PROTO_DIR &&
protoc --go_out=../ --go_opt=paths=source_relative *.proto
