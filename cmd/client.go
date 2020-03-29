package main

import "github.com/spf13/cobra"

func newClientCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Run the gRPC client",
		Run:   runClient,
	}
}

func runClient(cmd *cobra.Command, args []string) {

}
