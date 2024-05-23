# Go build with module replacement

```bash
go build -o bin/main main.go && ./bin/main
reversed: [5 4 3 2 1]
```

```bash
go build -o bin/main -modfile go-mods/go.mod main.go && ./bin/main
reversed: [5 4 3 2 1 5]
```