run-development:
	go run ./cmd/api/main.go

run-production:
	go run ./cmd/api/main.go -env=production