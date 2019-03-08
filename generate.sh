#!/bin/bash

protoc -I=./search --go_out=plugins=grpc:./search search/protoc/search.proto
