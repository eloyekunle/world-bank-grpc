package main

import (
	"fmt"
	"net"

	"github.com/eloyekunle/world-bank-grpc/pkg/server"
	"github.com/eloyekunle/world-bank-grpc/pkg/worldbank"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"
)

func newServerCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Args:  cobra.NoArgs,
		Short: "Run the gRPC server",
		Run:   runServer,
	}
}

func runServer(cmd *cobra.Command, args []string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50001))
	if err != nil {
		klog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	worldbank.RegisterWorldBankServer(grpcServer, server.NewServer())
	grpcServer.Serve(lis)
}
