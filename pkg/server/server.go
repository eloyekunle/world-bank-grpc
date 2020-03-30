//go:generate protoc -I ../worldbank --go_out=plugins=grpc:../worldbank ../worldbank/world_bank.proto

package server

import (
	pb "github.com/eloyekunle/world-bank-grpc/pkg/worldbank"
	"github.com/jkkitakita/wbdata-go"
)

type WorldBankServer struct {
	pb.UnimplementedWorldBankServer

	countries    map[string]*wbdata.Country
	regions      map[string]*wbdata.Region
	incomeLevels map[string]*wbdata.IncomeLevel
	lendingTypes map[string]*wbdata.LendingType
}

func NewServer() *WorldBankServer {
	s := &WorldBankServer{}
	return s
}
