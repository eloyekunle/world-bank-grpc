package main

import "github.com/spf13/cobra"

func newServerCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Run the gRPC server",
		Run:   runServer,
	}
}

func runServer(cmd *cobra.Command, args []string) {

}
