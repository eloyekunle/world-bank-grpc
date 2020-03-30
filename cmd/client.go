package main

import (
	"github.com/eloyekunle/world-bank-grpc/pkg/client"
	"github.com/eloyekunle/world-bank-grpc/pkg/util"
	pb "github.com/eloyekunle/world-bank-grpc/pkg/worldbank"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"
)

func newClientCommand() *cobra.Command {
	clientCmd := &cobra.Command{
		Use:   "client",
		Short: "Run the gRPC client",
	}

	regionsCmd := &cobra.Command{
		Use:   "regions",
		Short: "Fetch all regions",
		Run:   runClient,
	}

	incomeLevelsCmd := &cobra.Command{
		Use:   "income-levels",
		Short: "Fetch all income levels",
		Run:   runClient,
	}

	lendingTypesCmd := &cobra.Command{
		Use:   "lending-types",
		Short: "Fetch all lending types",
		Run:   runClient,
	}

	countriesCmd := &cobra.Command{
		Use:   "countries",
		Short: "Fetch all countries",
		Run:   runClient,
	}

	countriesCmd.Flags().StringP("region", "r", util.All, "Filter countries by region")
	countriesCmd.Flags().StringP("income-level", "i", util.All, "Filter countries by income level")
	countriesCmd.Flags().StringP("lending-type", "l", util.All, "Filter countries by lending type")

	clientCmd.AddCommand(regionsCmd, incomeLevelsCmd, lendingTypesCmd, countriesCmd)

	return clientCmd
}

func runClient(cmd *cobra.Command, args []string) {
	conn, err := grpc.Dial(":50001", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		klog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	c := pb.NewWorldBankClient(conn)

	switch cmd.Name() {
	case "countries":
		countryFilter := &pb.CountryFilter{
			Region:      mustStringFlag(cmd, "region"),
			IncomeLevel: mustStringFlag(cmd, "income-level"),
			LendingType: mustStringFlag(cmd, "lending-type"),
		}
		client.PrintCountries(c, countryFilter)
	case "regions":
		client.PrintRegions(c)
	case "income-levels":
		client.PrintIncomeLevels(c)
	case "lending-types":
		client.PrintLendingTypes(c)
	}
}

func mustStringFlag(cmd *cobra.Command, key string) string {
	out, err := cmd.Flags().GetString(key)
	if err != nil {
		klog.Fatal(err)
	}

	return out
}
