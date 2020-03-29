//go:generate protoc -I ../world_bank --go_out=plugins=grpc:../world_bank ../world_bank/world_bank.proto

package server
