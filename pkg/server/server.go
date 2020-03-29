//go:generate protoc -I ../worldbank --go_out=plugins=grpc:../worldbank ../worldbank/world_bank.proto

package server

import (
	pb "github.com/eloyekunle/world-bank-grpc/pkg/worldbank"
)

type WorldBankServer struct {
	pb.UnimplementedWorldBankServer
}

func NewServer() *WorldBankServer {
	s := &WorldBankServer{}
	return s
}
