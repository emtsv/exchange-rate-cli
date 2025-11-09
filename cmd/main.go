package main

import (
	"os"

	"github.com/emtsv/exchange-rate-cli/internal/commands"
	"github.com/spf13/cobra"
)

var version = "0.1.0"

func main() {
	rootCmd := &cobra.Command{
		Use:     "fx",
		Short:   "fx — CLI для получения курсов валют ЦБ РФ",
		Version: version,
	}

	rootCmd.AddCommand(commands.NewRateCMD())
	rootCmd.AddCommand(commands.NewListCMD())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
