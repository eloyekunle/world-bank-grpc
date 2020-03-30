package main

import (
	"github.com/eloyekunle/world-bank-grpc/pkg/client"
	"github.com/spf13/cobra"
)

func newClientCommand() *cobra.Command {
	clientCmd := &cobra.Command{
		Use:   "client",
		Short: "Run the gRPC client",
		Run:   runClient,
	}

	regionsCmd := &cobra.Command{
		Use:   "regions",
		Short: "Fetch all regions",
		Run:   runRegions,
	}

	incomeLevelsCmd := &cobra.Command{
		Use:   "income-levels",
		Short: "Fetch all income levels",
		Run:   runIncomeLevels,
	}

	lendingTypesCmd := &cobra.Command{
		Use:   "lending-types",
		Short: "Fetch all lending types",
		Run:   runLendingTypes,
	}

	clientCmd.AddCommand(regionsCmd, incomeLevelsCmd, lendingTypesCmd)

	return clientCmd
}

func runClient(cmd *cobra.Command, args []string) {
	//@todo
}

func runRegions(cmd *cobra.Command, args []string) {
	client.PrintRegions()
}

func runIncomeLevels(cmd *cobra.Command, args []string) {
	client.PrintIncomeLevels()
}

func runLendingTypes(cmd *cobra.Command, args []string) {
	client.PrintLendingTypes()
}
