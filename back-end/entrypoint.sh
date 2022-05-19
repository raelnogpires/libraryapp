wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"

go build -o ./bin/main main.go
./bin/main