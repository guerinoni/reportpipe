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
```zsh
make db-up
go test ./...
```
