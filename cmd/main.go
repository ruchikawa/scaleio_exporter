package main

import (
	"log"
	"os"

	"github.com/ruchikawa/scaleio_exporter/cmd/scaleio_exporter/cmd"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	rootCmd := cmd.GetRootCmd(os.Args[1:])

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
