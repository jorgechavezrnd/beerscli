package main

import (
	"github.com/jorgechavezrnd/test_project/internal/cli"
	"github.com/jorgechavezrnd/test_project/internal/storage/csv"
	"github.com/spf13/cobra"
)

func main() {
	csvRepo := csv.NewRepository()

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(csvRepo))
	rootCmd.Execute()
}
