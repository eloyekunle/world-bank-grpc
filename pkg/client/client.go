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

	i := 1
	for {
		region, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			klog.Exitf("%v.PrintRegions(_) = _, %v: ", client, err)
		}

		fmt.Printf("%d - %s (ID: %s)\n", i, region.GetName(), region.GetId())
		i++
	}
}

func PrintIncomeLevels(ctx context.Context, client pb.WorldBankClient) {
	stream, err := client.ListIncomeLevels(ctx, &pb.Void{})
	if err != nil {
		klog.Exitf("%v.PrintIncomeLevels(_) = _, %v: ", client, err)
	}

	i := 1
	for {
		incomeLevel, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			klog.Exitf("%v.PrintIncomeLevels(_) = _, %v: ", client, err)
		}

		fmt.Printf("%d - %s (ID: %s)\n", i, incomeLevel.GetName(), incomeLevel.GetId())
		i++
	}
}

func PrintLendingTypes(ctx context.Context, client pb.WorldBankClient) {
	stream, err := client.ListLendingTypes(ctx, &pb.Void{})
	if err != nil {
		klog.Exitf("%v.PrintLendingTypes(_) = _, %v: ", client, err)
	}

	i := 1
	for {
		lendingType, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			klog.Exitf("%v.PrintLendingTypes(_) = _, %v: ", client, err)
		}

		fmt.Printf("%d - %s (ID: %s)\n", i, lendingType.GetName(), lendingType.GetId())
		i++
	}
}

func PrintCountries(ctx context.Context, client pb.WorldBankClient, filter *pb.CountryFilter) {
	stream, err := client.ListCountries(ctx, filter)
	if err != nil {
		klog.Exitf("%v.PrintCountries(_) = _, %v: ", client, err)
	}

	i := 1
	for {
		country, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			klog.Exitf("%v.PrintCountries(_) = _, %v: ", client, err)
		}

		fmt.Printf("%d - %s (ID: %s)\n", i, country.GetName(), country.GetId())
		i++
	}
}

func PrintCountry(ctx context.Context, client pb.WorldBankClient, countryID *pb.CountryID) {
	country, err := client.GetCountry(ctx, countryID)
	if err != nil {
		klog.Exitf("%v.PrintCountry(_) = _, %v: ", client, err)
	}

	fmt.Printf("ID: %s\nName: %s\nCapital: %s\nRegion: %s\nIncome Level: %s\nLending Type: %s\n",
		country.GetId(), country.GetName(), country.GetCapital(), country.GetRegion(), country.GetIncomeLevel(), country.GetLendingType())
}
