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
	port := 50001

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		klog.Exitf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	worldbank.RegisterWorldBankServer(grpcServer, server.NewServer())

	klog.Infof("gRPC server listening on port: %d", port)
	klog.Exit(grpcServer.Serve(lis))
}
