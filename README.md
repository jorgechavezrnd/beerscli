# Test project for Go course in CodelyTV

# Commands for create project
- In $GOPATH/src/github.com/jorgechavezrnd: `mkdir test_project`
- In test_project directory: `go mod init`

# Example commands for run the app (for use csv repository, modifie main.go)
- `go run cmd/beers-cli/main.go`
- `go run cmd/beers-cli/main.go beers --help`
- `go run cmd/beers-cli/main.go beers`
- `go run cmd/beers-cli/main.go beers -i 131` or `go run cmd/beers-cli/main.go beers --id 131`
- If we run `go run cmd/beers-cli/main.go beer`, it will show a pretty error message, and will show suggestions

# For add cobra
- Add "github.com/spf13/cobra" in imports where it will be used
- Run `go mod tidy` (This is something like `composer install` or `composer update`, It get all the dependencies that we have in "imports" and add the dependencies in 'go.mod' file and create 'go.sum' file)

# Commands for install moq and generate mock repository with moq (URL: https://github.com/matryer/moq)
- Install moq: `go get github.com/matryer/moq`
- `cd internal/storage/mock`
- `moq -pkg mock -out repository.go ../../. BeerRepo`

# Command for run all tests
- `go test -v ./...`

# Comand example for run one specific test
- `go test ./... -run=TestFetchByID/"valid beer"`
