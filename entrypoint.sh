# wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"
wait-for "db:3306" -t 0 -- "$@" 

# Watch your .go files and invoke go build if the files changed.
# CompileDaemon --build="go build -o main main.go"  --command=./main
go run main.go