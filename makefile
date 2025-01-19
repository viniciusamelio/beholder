PHONY: .dev .schema


dev:
	go run cmd/api/main.go

schema:
	go run github.com/go-surreal/som/cmd/som@latest gen internal/application/models internal/gen/som