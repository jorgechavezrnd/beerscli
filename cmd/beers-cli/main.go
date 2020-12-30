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
	// CPU profiling code starts here
	// f, _ := os.Create("beers.cpu.prof")
	// defer f.Close()
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()
	// CPU profiling code ends here

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

	// Memory profiling code starts here
	// f2, _ := os.Create("beers.mem.prof")
	// defer f2.Close()
	// pprof.WriteHeapProfile(f2)
	// Memory profiling code ends here
}
