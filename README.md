# reportpipe

### Development
```zsh
make db-up
go run ./cmd/main.go
```

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
