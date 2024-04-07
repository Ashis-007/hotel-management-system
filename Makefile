install:
	go mod tidy

dev:
	GO_ENV="development" go run ./cmd/web/

prod:
	GO_ENV="production" go run ./cmd/web/
