wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"

# watch your .go files and invoke `go build`` if the files changed.
CompileDaemon --build="go build -o main main.go"  --command=./main