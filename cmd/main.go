package main

import "k8s.io/klog/v2"

func main() {
	rootCmd := newServerCommand()
	rootCmd.AddCommand(newVersionCmd(), newClientCommand())

	klog.Exit(rootCmd.Execute())
}
