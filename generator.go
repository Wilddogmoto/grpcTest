//go:generate protoc -I ./ sender/sender.proto --go_out=plugins=grpc:./

package main
