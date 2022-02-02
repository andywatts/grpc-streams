#!/bin/bash

#export GOPATH=$HOME/Documents/goworkspace
export PATH=$PATH:$HOME/go/bin
export PATH=$PATH:/usr/local/go/bin

SRC_DIR=.
DST_DIR=.
PROTO_FILE=$1
protoc  -I=$SRC_DIR \
        --go_out=$DST_DIR \
        --go_opt=paths=source_relative \
        --go-grpc_out=$DST_DIR \
        --go-grpc_opt=paths=source_relative \
        $SRC_DIR/$PROTO_FILE

