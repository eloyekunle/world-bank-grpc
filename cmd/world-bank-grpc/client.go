package main

import (
	"context"
	"time"

	"github.com/eloyekunle/world-bank-grpc/pkg/env"

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

	countriesCmd.Flags().StringP("region", "r", util.All, "Filter countries by region ID")
	countriesCmd.Flags().StringP("income-level", "i", util.All, "Filter countries by income level ID")
	countriesCmd.Flags().StringP("lending-type", "l", util.All, "Filter countries by lending type ID")

	countryCmd := &cobra.Command{
		Use:   "country",
		Short: "Get a single country",
		Run:   runClient,
	}

	countryCmd.Flags().String("id", util.All, "Country ID to fetch")
	must(countryCmd.MarkFlagRequired("id"))

	clientCmd.AddCommand(regionsCmd, incomeLevelsCmd, lendingTypesCmd, countriesCmd, countryCmd)

	return clientCmd
}

func runClient(cmd *cobra.Command, args []string) {
	port := env.GetEnvFallback(env.EnvPort, env.DefaultPort)
	target := ":" + port

	conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		klog.Exitf("fail to dial: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(cmd.Context(), 10*time.Second)
	defer cancel()

	c := pb.NewWorldBankClient(conn)

	switch cmd.Name() {
	case "countries":
		countryFilter := &pb.CountryFilter{
			Region:      mustStringFlag(cmd, "region"),
			IncomeLevel: mustStringFlag(cmd, "income-level"),
			LendingType: mustStringFlag(cmd, "lending-type"),
		}
		client.PrintCountries(ctx, c, countryFilter)
	case "country":
		countryID := &pb.CountryID{Id: mustStringFlag(cmd, "id")}
		client.PrintCountry(ctx, c, countryID)
	case "regions":
		client.PrintRegions(ctx, c)
	case "income-levels":
		client.PrintIncomeLevels(ctx, c)
	case "lending-types":
		client.PrintLendingTypes(ctx, c)
	}
}

func mustStringFlag(cmd *cobra.Command, key string) string {
	out, err := cmd.Flags().GetString(key)
	if err != nil {
		klog.Exit(err)
	}

	return out
}
