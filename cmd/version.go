package main

import (
	"fmt"

	"github.com/eloyekunle/world-bank-grpc/pkg/version"
	"github.com/spf13/cobra"
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Display the app version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version.VERSION)
		},
	}
}
