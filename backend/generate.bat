@echo off
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest

protoc -I . ^
  --go_out=. --go_opt=paths=source_relative ^
  --connect-go_out=. --connect-go_opt=paths=source_relative ^
  --ts_out=. --ts_opt=paths=source_relative ^
  api/calculator.proto 