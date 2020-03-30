package client

import (
	"context"
	"io"
	"log"
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
		log.Println(region)
	}
}

func PrintIncomeLevels() {

}

func PrintLendingTypes() {

}
