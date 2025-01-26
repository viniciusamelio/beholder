PHONY: .dev .schema


dev:
	air

schema:
	go run github.com/go-surreal/som/cmd/som@latest gen internal/application/models internal/gen/som