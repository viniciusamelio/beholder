dev:
	air
up:
	goose up
down:
	goose down
schema:
	jet --source=sqlite -dsn=./beholder.db -schema=dvds -path=./internal/jet
tailwind:
	tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch
templ:
	templ generate --watch