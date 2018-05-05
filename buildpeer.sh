#!/bin/bash

BUILD_BIN=build/bin

if [ ! -d ${BUILD_BIN} ]; then
    mkdir -p ${BUILD_BIN}
fi

if [ -f ${BUILD_BIN}/peer ]; then
    rm ${BUILD_BIN}/peer
fi

CGO_CFLAGS=" " CGO_LDFLAGS="-lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy" \
    GOBIN=${GOPATH}/src/github.com/abchain/fabric/${BUILD_BIN} go install github.com/abchain/fabric/peer


if [ ! -f ${BUILD_BIN}/peer ]; then
    echo 'Failed to build peer!'
    exit -1
fi
