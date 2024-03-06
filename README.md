# reportpipe

### Run the BE side


```zsh
make db-up
go run ./cmd/main.go
```

For the FE part look at [readme](./app/README.md)

### Create new table in psql

This is an example for creating a new table in db.

```zsh
migrate create -ext sql -dir migrations -seq users
```

### Testing

The testing suite is based only on integration tests and require docker running.


```zsh
make db-up
go test ./...
```
