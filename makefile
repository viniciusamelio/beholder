dev:
	air
schema:
	go run github.com/go-surreal/som/cmd/som@latest gen internal/application/models internal/gen/som
tailwind:
	tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch
templ:
	templ generate --watch