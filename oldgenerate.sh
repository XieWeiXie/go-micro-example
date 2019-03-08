#!/bin/bash


protoc -I=. --go_out=. search/protoc/search.proto
