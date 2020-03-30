package client

import (
	"context"
	"fmt"
	"io"
	"time"

	pb "github.com/eloyekunle/world-bank-grpc/pkg/worldbank"
	"k8s.io/klog/v2"
)

func PrintRegions(client pb.WorldBankClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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

func PrintIncomeLevels(client pb.WorldBankClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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

func PrintLendingTypes(client pb.WorldBankClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.ListLendingTypes(ctx, &pb.Void{})
	if err != nil {
		klog.Fatalf("%v.PrintLendingTypes(_) = _, %v: ", client, err)
	}

	for {
		lendingType, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			klog.Fatalf("%v.PrintLendingTypes(_) = _, %v: ", client, err)
		}

		fmt.Println(lendingType)
	}
}

func PrintCountries(client pb.WorldBankClient, filter *pb.CountryFilter) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.ListCountries(ctx, filter)
	if err != nil {
		klog.Fatalf("%v.PrintCountries(_) = _, %v: ", client, err)
	}

	for {
		country, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			klog.Fatalf("%v.PrintCountries(_) = _, %v: ", client, err)
		}

		fmt.Println(country)
	}
}
