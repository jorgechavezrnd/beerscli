# Test project for Go course in CodelyTV

# Commands for create project
- In $GOPATH/src/github.com/jorgechavezrnd: `mkdir test_project`
- In test_project directory: `go mod init`

# Example commands for run the app
- `go run cmd/beers-cli/main.go`
- `go run cmd/beers-cli/main.go beers --help`
- `go run cmd/beers-cli/main.go beers`
- `go run cmd/beers-cli/main.go beers -i 01D9X5BQ5X48XMMVZ2F2G3R5MS` or `go run cmd/beers-cli/main.go beers --id 01D9X5BQ5X48XMMVZ2F2G3R5MS`
- If we run `go run cmd/beers-cli/main.go beer`, it will show a pretty error message, and will show suggestions

# For add cobra
- Add "github.com/spf13/cobra" in imports where it will be used
- Run `go mod tidy` (This is something like `composer install` or `composer update`, It get all the dependencies that we have in "imports" and add the dependencies in 'go.mod' file and create 'go.sum' file)
