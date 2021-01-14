package fetching

import (
	"math"

	beerscli "github.com/jorgechavezrnd/test_project/internal"
	"github.com/pkg/errors"
)

// Service provides beer fetching operations
type Service interface {
	// FetchBeers fetch all beers from repository
	FetchBeers() ([]beerscli.Beer, error)
	// FetchByID filter all beers and get only the beer that match with given id
	FetchByID(id int) (beerscli.Beer, error)
}

type service struct {
	bR beerscli.BeerRepo
}

// NewService creates an adding service with the necessary dependencies
func NewService(r beerscli.BeerRepo) Service {
	return &service{r}
}

func (s *service) FetchBeers() ([]beerscli.Beer, error) {
	return s.bR.GetBeers()
}

func (s *service) FetchByID(id int) (beerscli.Beer, error) {
	beers, err := s.FetchBeers()
	if err != nil {
		return beerscli.Beer{}, err
	}

	beersPerRoutine := 10
	numRoutines := numberOfRoutines(len(beers), beersPerRoutine)

	b := make(chan beerscli.Beer)
	done := make(chan bool, numRoutines)

	for i := 0; i < numRoutines; i++ {
		toSearch := make([]beerscli.Beer, beersPerRoutine)
		init := i * beersPerRoutine
		end := init + beersPerRoutine - 1
		if len(beers) < end {
			end = len(beers)
		}
		copy(toSearch[:], beers[init:end])

		go func(beers []beerscli.Beer, b chan beerscli.Beer, done chan bool) {
			for _, beer := range beers {
				if beer.ProductID == id {
					b <- beer
				}
			}

			done <- true
		}(toSearch, b, done)
	}

	var beer beerscli.Beer
	i := 0
	for i < numRoutines {
		select {
		case beer = <-b:
			return beer, nil
		case <-done:
			i++
		}
	}

	return beerscli.Beer{}, errors.Errorf("Beer %d not found", id)
}

func numberOfRoutines(numbOfBeers, beersPerRoutine int) int {
	return int(math.Ceil(float64(numbOfBeers) / float64(beersPerRoutine)))
}
