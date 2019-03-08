help="make command ..."



default:
	echo ${help}



protoc:
	bash protoc.sh

multiply:
	bash multiply.sh

tidy:
	go mod tidy

vendor:
	go mod vendor

init:
	go mod init

.PHONY: default exprot protoc
