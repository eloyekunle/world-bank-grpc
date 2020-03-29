package main

import "k8s.io/klog/v2"

func main() {
	rootCmd := newServerCommand()
	rootCmd.AddCommand(newVersionCmd(), newClientCommand())

	must(rootCmd.Execute())
}

func must(err error) {
	if err != nil {
		klog.Exit(err)
	}
}
