# wait for MySQL env variables are defined before continue initialization
wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"

# sets Gin to be on release, not debug mode
export GIN_MODE=release

# builds application binary for better perfomance
go build -o ./bin/main main.go

# runs binary
./bin/main