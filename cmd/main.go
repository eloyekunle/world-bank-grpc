package main

import (
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

func main() {
	rootCmd := &cobra.Command{Use: "world-bank-grpc"}
	rootCmd.AddCommand(newServerCommand(), newVersionCmd(), newClientCommand())

	must(rootCmd.Execute())
}

func must(err error) {
	if err != nil {
		klog.Exit(err)
	}
}
