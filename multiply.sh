#!/usr/bin/env bash

dir=`pwd`
echo ${dir}
for i in ${dir}/srv/auth/protoc ${dir}/srv/report/protoc; do
    for f in ${i}/*.proto; do
        protoc --proto_path=${GOPATH}/src --go_out=plugins=grpc:${GOPATH}/src $f;
        echo compiled:  ${f};
    done
done
