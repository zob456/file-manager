# File Manager
### Checks for duplicate code in csv files & compresses csv files

#### Commands to execute different functions

Pass 1 of 2 different command lines from inside the root directory of this project
- `go run ./cmd/main.go checker` executes the `file checker` function that checks multiple csv files for duplicate codes
- `go run ./cmd/main.go compressor` executes the `file compressor` function that compresses a csv file & logs the before & after sizes of the files
