#!/usr/bin/env bash
# working with protoc version libprotoc 3.17.3
cd proto && protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative store.proto


