package cli

import (
	"fmt"

	beerscli "github.com/jorgechavezrnd/test_project/internal"
	"github.com/spf13/cobra"
)

// InitPokemonCmd initialize pokemon command
func InitPokemonCmd(repository beerscli.PokemonRepo) *cobra.Command {
	pokemonCmd := &cobra.Command{
		Use:   "pokemon",
		Short: "Print data about a pokemon",
		Run:   runPokemonFn(repository),
	}

	pokemonCmd.Flags().StringP(idFlag, "i", "", "id or name of the pokemon")

	return pokemonCmd
}

func runPokemonFn(repository beerscli.PokemonRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idFlag)
		pokemon, _ := repository.GetPokemon(id)

		fmt.Println(pokemon)
	}
}
