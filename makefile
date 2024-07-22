serve:
	go run cmd/api/main.go
migrate:
	go run cmd/migration/main.go
test:
	go test -v ./tests/...