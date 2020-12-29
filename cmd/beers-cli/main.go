package main

import (
	"flag"

	"github.com/jorgechavezrnd/test_project/internal/cli"
	"github.com/jorgechavezrnd/test_project/internal/fetching"
	"github.com/jorgechavezrnd/test_project/internal/storage/csv"
	"github.com/jorgechavezrnd/test_project/internal/storage/ontario"
	"github.com/spf13/cobra"

	beerscli "github.com/jorgechavezrnd/test_project/internal"
)

func main() {
	csvData := flag.Bool("csv", false, "load data from csv")
	flag.Parse()

	var repo beerscli.BeerRepo
	repo = csv.NewRepository()

	if !*csvData {
		repo = ontario.NewOntarioRepository()
	}

	fetchingService := fetching.NewService(repo)

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(fetchingService))
	rootCmd.Execute()
}
