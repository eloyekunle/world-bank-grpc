package client

import (
	"context"
	"fmt"
	"io"

	pb "github.com/eloyekunle/world-bank-grpc/pkg/worldbank"
	"k8s.io/klog/v2"
)

func PrintRegions(ctx context.Context, client pb.WorldBankClient) {
	stream, err := client.ListRegions(ctx, &pb.Void{})
	if err != nil {
		klog.Exitf("%v.PrintRegions(_) = _, %v: ", client, err)
	}

	for {
		region, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			klog.Exitf("%v.PrintRegions(_) = _, %v: ", client, err)
		}

		fmt.Println(region)
	}
}

func PrintIncomeLevels(ctx context.Context, client pb.WorldBankClient) {
	stream, err := client.ListIncomeLevels(ctx, &pb.Void{})
	if err != nil {
		klog.Exitf("%v.PrintIncomeLevels(_) = _, %v: ", client, err)
	}

	for {
		incomeLevel, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			klog.Exitf("%v.PrintIncomeLevels(_) = _, %v: ", client, err)
		}

		fmt.Println(incomeLevel)
	}
}

func PrintLendingTypes(ctx context.Context, client pb.WorldBankClient) {
	stream, err := client.ListLendingTypes(ctx, &pb.Void{})
	if err != nil {
		klog.Exitf("%v.PrintLendingTypes(_) = _, %v: ", client, err)
	}

	for {
		lendingType, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			klog.Exitf("%v.PrintLendingTypes(_) = _, %v: ", client, err)
		}

		fmt.Println(lendingType)
	}
}

func PrintCountries(ctx context.Context, client pb.WorldBankClient, filter *pb.CountryFilter) {
	stream, err := client.ListCountries(ctx, filter)
	if err != nil {
		klog.Exitf("%v.PrintCountries(_) = _, %v: ", client, err)
	}

	for {
		country, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			klog.Exitf("%v.PrintCountries(_) = _, %v: ", client, err)
		}

		fmt.Println(country)
	}
}

func PrintCountry(ctx context.Context, client pb.WorldBankClient, countryID *pb.CountryID) {
	country, err := client.GetCountry(ctx, countryID)
	if err != nil {
		klog.Exitf("%v.PrintCountry(_) = _, %v: ", client, err)
	}

	fmt.Println(country)
}
