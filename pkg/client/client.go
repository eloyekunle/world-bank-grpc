package client

import (
	"context"
	"fmt"
	"io"
	"time"

	pb "github.com/eloyekunle/world-bank-grpc/pkg/worldbank"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"
)

func newClient() (*grpc.ClientConn, pb.WorldBankClient) {
	conn, err := grpc.Dial(":50001", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		klog.Fatalf("fail to dial: %v", err)
	}

	client := pb.NewWorldBankClient(conn)
	return conn, client
}

func PrintRegions() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, client := newClient()
	defer conn.Close()

	stream, err := client.ListRegions(ctx, &pb.Void{})
	if err != nil {
		klog.Fatalf("%v.PrintRegions(_) = _, %v: ", client, err)
	}

	for {
		region, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			klog.Fatalf("%v.PrintRegions(_) = _, %v: ", client, err)
		}

		fmt.Println(region)
	}
}

func PrintIncomeLevels() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, client := newClient()
	defer conn.Close()

	stream, err := client.ListIncomeLevels(ctx, &pb.Void{})
	if err != nil {
		klog.Fatalf("%v.PrintIncomeLevels(_) = _, %v: ", client, err)
	}

	for {
		incomeLevel, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			klog.Fatalf("%v.PrintIncomeLevels(_) = _, %v: ", client, err)
		}

		fmt.Println(incomeLevel)
	}
}

func PrintLendingTypes() {

}
